# guess-it-2 — Product Requirements Document

| | |
|---|---|
| **Project** | guess-it-2 (Zone01 Athens, algorithm chain: math-skills → guess-it-1 → linear-stats → guess-it-2) |
| **Author** | Nikos |
| **Language** | Go (standard library only) |
| **Status** | Implemented, not yet verified against the official tester |
| **Companion file** | [TASK_CARDS.md](TASK_CARDS.md) |

> Sections marked **`TODO (own words)`** are the reasoning behind a decision. They are filled in by me after I have tested and understood the decision, not before.

## Repository structure

```
guess-it-2/
├── go.mod                  Go module definition (module guess-it-2)
├── main.go                 Reads stdin, calls PredictRange, prints "low high"
├── script.sh               Command the tester runs: ./student/guess-it-2
├── internals/
│   ├── stats.go            Point, FitLine (line + Pearson r), Quantile, IqrFilter
│   ├── predict.go          PredictRange and all tunable parameters
│   └── predict_test.go     Unit tests for the edge cases in section 8
├── student/                Build output only (not committed): the compiled
│                           binary guess-it-2, copied to the tester's root folder
└── docs/
    ├── PRD.md              This document: what and why
    └── TASK_CARDS.md       The work broken into cards, with status
```

In the `algorithm-chain` repository this folder sits next to `math-skills/`, `guess-it-1/` and `linear-stats/`, each with its own `docs/PRD.md` and `docs/TASK_CARDS.md`.

---

## 1. Problem

A program receives numbers one per line on standard input. After **every** number it must print a range `low high` in which it believes the **next** number will fall.

The data is a graph: x is the line number (0, 1, 2, …), y is the value. The tester scores a hit by the **size of the range**: a correct but narrow range scores more than a correct but wide one. A wrong range scores nothing.

The core tension: **narrow range vs. probability of being right.**

## 2. Goals and non-goals

**Goals**
- Beat the reference programs `big-range`, `linear-regr`, `correlation-coef` on Data 4 and Data 5 (at least 2 wins out of 3 runs each).
- Bonus: beat `mse` and `nic`.
- Reuse the statistics from `linear-stats` (regression line, Pearson correlation) instead of writing new, unrelated maths.
- Stay fast: the output must follow each input line with no noticeable delay.

**Non-goals**
- No machine-learning libraries, no external dependencies.
- No reading of the whole file in advance. The program only knows the past, exactly like in the real test.

## 3. Input and output specification

| | |
|---|---|
| **Input** | One number per line on stdin. Blank lines and lines that are not numbers are ignored. |
| **Output** | After each valid number, one line: `<low> <high>`, two integers separated by one space, `low <= high`. |
| **Flush** | Output is flushed after every line, because the tester waits for the answer before sending the next number. |
| **Run** | `script.sh` runs `./student/guess-it-2` from the root folder of the tester. |

## 4. Functional requirements

| ID | Requirement |
|---|---|
| FR-1 | Read stdin line by line and parse each line as a float. |
| FR-2 | After each number print a range for the next number. |
| FR-3 | Use only the most recent values (sliding window), so the program adapts when the data changes. |
| FR-4 | Remove outliers before fitting, **without losing the original x position** of the remaining points. |
| FR-5 | Fit a line and compute Pearson r on the window (the `linear-stats` calculations). |
| FR-6 | Decide whether the data has a real trend or is noise around a level, and predict accordingly. |
| FR-7 | Make the range width depend on the expected error of the prediction, not on a fixed constant. |
| FR-8 | Never crash and never print `NaN` or an inverted range, for any input length (including 1 number, or all equal numbers). |

## 5. Non-functional requirements

- **Performance:** each new number costs O(W log W) with W = 40 (one sort for the quartiles). Independent of how many lines came before.
- **Robustness:** every division in the statistics code is guarded (see FR-8).
- **Portability:** a static Linux binary (`CGO_ENABLED=0`) that runs in the tester's Docker container.
- **Readability:** one place for all tunable numbers (the `const` block in `internals/predict.go`).

## 6. Success criteria (the audit)

| Audit question | Target | Status |
|---|---|---|
| Beats `big-range` on Data 4, ≥ 2 of 3 | Required | ☐ not run yet |
| Beats `big-range` on Data 5, ≥ 2 of 3 | Required | ☐ not run yet |
| Beats `linear-regr` on Data 4 and Data 5, ≥ 2 of 3 each | Required | ☐ not run yet |
| Beats `correlation-coef` on Data 4 and Data 5, ≥ 2 of 3 each | Required | ☐ not run yet |
| Beats `mse` on Data 4 and Data 5, ≥ 2 of 3 each | Bonus | ☐ not run yet |
| Beats `nic` on Data 4 and Data 5, ≥ 2 of 3 each | Bonus | ☐ not run yet |

## 7. Approach

For every new number the program does the following on the **latest 40 values**:

1. **Window.** Take the last 40 values. Each gets x = its position in the window.
2. **Outlier filter.** Drop values outside `[Q1 − 1.6·IQR, Q3 + 1.6·IQR]`. Each surviving point keeps its original x.
3. **Fit.** Compute the regression line `y = a·x + b` and the Pearson coefficient `r` (deviation form of the `linear-stats` formulas).
4. **Trend or noise?** Convert `r` to a t-value: `t = |r|·√((n−2)/(1−r²))`.
   - `t > 2.5` → the trend is real: predict with the line; expected error is the residual standard deviation, widened for extrapolation (prediction-interval formula).
   - otherwise → the data is noise around a level: predict with the mean; expected error is the standard deviation of the values.
5. **Range.** `prediction ± 1.5 · expected error`, rounded outwards to integers, never narrower than ±1.
6. **Warm-up.** With fewer than 5 values there is not enough data for statistics: return `min − 15` to `max + 15` of what has been seen.

### Parameters

| Parameter | Value | Meaning |
|---|---|---|
| `windowSize` | 40 | How many recent values are used |
| `iqrFence` | 1.6 | Outlier fence (same value as guess-it-1) |
| `minFitPoints` | 8 | Fewest points on which a trend decision is made |
| `tCritical` | 2.5 | How strong the correlation must be to trust the trend |
| `widthFactor` | 1.5 | Range half-width in units of expected error |
| `minHalfWidth` | 1.0 | Smallest allowed half-width |

**Why these values — TODO (own words):**
- Why a window of 40 and not 20 or 200?
- Why decide trend-vs-noise with Pearson r instead of always using the line?
- Why `widthFactor` = 1.5? What happens to hit rate and score when it goes up or down?

### Reuse from earlier projects

| From | What is reused | What changed |
|---|---|---|
| linear-stats | Regression line and Pearson r | Deviation form instead of the sums form (safer with large values); every division guarded; used inside a prediction |
| guess-it-1 | Sliding window, IQR filter, stdin loop | The filter now keeps x positions (before, values were re-numbered after removal, which shifted the prediction); fixed ±19 replaced by an error-based width; bubble sort replaced by `sort.Float64s` |
| math-skills | Mean and standard deviation | Used as the "no trend" model and for the error width |

## 8. Risks and edge cases

| Risk | Handling |
|---|---|
| 0 or 1 number so far | Warm-up range; `PredictRange` handles an empty slice |
| All numbers equal (zero variance) | `r` returns 0 instead of NaN; `minHalfWidth` keeps the range valid |
| Outlier filter removes almost everything (IQR = 0) | If fewer than 3 points survive, the unfiltered window is used |
| Level shifts or regime changes in the data | The window forgets old values after 40 steps; this is a known weakness for sudden jumps |
| Data 4 / Data 5 behave differently from my synthetic data | The tester is the real judge; `widthFactor` and `windowSize` are tuned from its results |
| Wrong CPU architecture of the binary in Docker | Build with `GOOS=linux GOARCH=amd64`; fall back to `go run` if the container has Go |

## 9. Testing strategy

1. **Unit tests** (`go test ./...`): tiny inputs, constant data, a perfect linear trend, outlier position handling, line-fit correctness.
2. **Simulation on synthetic series** (constant level, trend, level shifts, random walk, sine, spikes): compare hit rate and average width against a fixed ±19 line baseline. This is a sanity check only. The official score formula is unknown to me, so these numbers are not a promise about the audit.
3. **Official tester** against all five reference programs on Data 4 and Data 5, 3 runs each. Results go into the table in section 6 and drive the tuning.

## 10. Open questions

- How does the official score weigh range size against hits? (Determines the best `widthFactor`.)
- Do Data 4 / Data 5 contain trends, or mostly noise around a level?
- Is Go available inside the tester container, or must a compiled binary be shipped?