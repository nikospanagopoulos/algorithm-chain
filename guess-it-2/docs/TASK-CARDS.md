# guess-it-2 — Task Cards

Companion to [PRD.md](PRD.md).

**Status legend:** ☐ todo · ◐ in review (written and tested by me, not yet confirmed by the official tester) · ☑ done

| Card | Title | Status |
|---|---|---|
| T-01 | Project scaffold and build | ◐ |
| T-02 | Stdin loop and output | ◐ |
| T-03 | Statistics core (line fit, Pearson, guards) | ◐ |
| T-04 | Outlier filter that keeps x | ◐ |
| T-05 | `PredictRange` (window, trend gate, error-based width) | ◐ |
| T-06 | Unit tests | ◐ |
| T-07 | Audit: `big-range` | ☐ |
| T-08 | Audit: `linear-regr` | ☐ |
| T-09 | Audit: `correlation-coef` | ☐ |
| T-10 | Tune parameters from tester results | ☐ |
| T-11 | Bonus: `mse` and `nic` | ☐ |
| T-12 | Documentation and GitHub publication | ☐ |

---

## T-01 — Project scaffold and build
**Goal:** a Go module that builds into a static Linux binary the tester can run.
**Acceptance criteria**
- `go.mod` with module `guess-it-2`.
- `go build` works; `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o student/guess-it-2 .` works.
- `script.sh` is executable and runs `./student/guess-it-2`.

**Depends on:** none  **Covers:** PRD §3, §5
**Learned (own words):** _TODO_

## T-02 — Stdin loop and output
**Goal:** read numbers line by line and answer after each one.
**Acceptance criteria**
- Blank and non-numeric lines are skipped.
- One `low high` line is printed and flushed after every valid number.

**Depends on:** T-01  **Covers:** FR-1, FR-2
**Learned (own words):** _TODO_

## T-03 — Statistics core
**Goal:** the maths of `linear-stats`, reusable inside the predictor.
**Acceptance criteria**
- `FitLine` returns slope, intercept, Pearson r, residual and plain standard deviation.
- Known line `y = 2x + 1` gives `a = 2`, `b = 1`, `r = 1`.
- A single point, or all-equal values, never produce `NaN`.

**Depends on:** T-01  **Covers:** FR-5, FR-8
**Questions to answer in my own words:** Why is the deviation form safer than the sums form with large values? What does r say about the trend and what does it *not* say about the size of the error?
**Learned (own words):** _TODO_

## T-04 — Outlier filter that keeps x
**Goal:** remove outliers without shifting the positions of the remaining values.
**Acceptance criteria**
- After removing an outlier from the middle, the next point still has its original x.
- If the filter leaves fewer than 3 points, the unfiltered window is used.

**Depends on:** T-03  **Covers:** FR-4
**Questions to answer in my own words:** What was wrong in guess-it-1 when x was re-numbered after filtering? How many positions was the prediction off when 3 values were removed?
**Learned (own words):** _TODO_

## T-05 — `PredictRange`
**Goal:** a range whose width follows the expected error, using a trend only when it is real.
**Acceptance criteria**
- Uses the last 40 values only.
- Decides trend vs. noise with the t-value of r (`tCritical = 2.5`).
- Width = `widthFactor` × expected error, never below ±1.
- Warm-up behaviour for fewer than 5 values.
- Never returns `low > high`, never `NaN`.

**Depends on:** T-03, T-04  **Covers:** FR-3, FR-6, FR-7, FR-8
**Questions to answer in my own words:** Why is a flat mean sometimes a better predictor than a regression line? What does the extrapolation term in the error formula do?
**Learned (own words):** _TODO_

## T-06 — Unit tests
**Goal:** the edge cases from PRD §8 are covered by automated tests.
**Acceptance criteria**
- `go test ./...` passes.
- Covered: 1–7 numbers, constant data, perfect trend, random data never inverted, x preserved after filtering, line fit, single point.

**Depends on:** T-03, T-04, T-05
**Learned (own words):** _TODO_

## T-07 — Audit: `big-range`
**Goal:** win at least 2 of 3 runs on Data 4 and on Data 5.
**Acceptance criteria:** results recorded below.

| Dataset | Run 1 | Run 2 | Run 3 | Pass |
|---|---|---|---|---|
| Data 4 | | | | ☐ |
| Data 5 | | | | ☐ |

**Depends on:** T-01, T-02, T-05

## T-08 — Audit: `linear-regr`
**Goal:** win at least 2 of 3 runs on Data 4 and on Data 5.

| Dataset | Run 1 | Run 2 | Run 3 | Pass |
|---|---|---|---|---|
| Data 4 | | | | ☐ |
| Data 5 | | | | ☐ |

**Depends on:** T-07

## T-09 — Audit: `correlation-coef`
**Goal:** win at least 2 of 3 runs on Data 4 and on Data 5.

| Dataset | Run 1 | Run 2 | Run 3 | Pass |
|---|---|---|---|---|
| Data 4 | | | | ☐ |
| Data 5 | | | | ☐ |

**Depends on:** T-07

## T-10 — Tune parameters from tester results
**Goal:** improve the weakest result using evidence from the tester, one parameter at a time.
**Acceptance criteria**
- Each change is recorded: parameter, old value, new value, effect on the results.
- Parameters in scope: `widthFactor`, `windowSize`, `tCritical`, `iqrFence`.

| Change | Old → New | Effect |
|---|---|---|
| | | |

**Depends on:** T-07, T-08, T-09

## T-11 — Bonus: `mse` and `nic`
**Goal:** win at least 2 of 3 runs on Data 4 and on Data 5 against each.

| Opponent | Data 4 | Data 5 | Pass |
|---|---|---|---|
| `mse` | | | ☐ |
| `nic` | | | ☐ |

**Depends on:** T-10

## T-12 — Documentation and GitHub publication
**Goal:** the repository explains itself.
**Acceptance criteria**
- `docs/PRD.md` and `docs/TASK_CARDS.md` are up to date (all `TODO (own words)` filled, statuses true).
- A README explains how to build and how to test with the official tester.
- The project is pushed to GitHub.

**Depends on:** T-10