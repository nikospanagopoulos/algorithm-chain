# guess-it-2

A Go program that predicts the next number of a data stream.

The program reads numbers from standard input, one per line. After every number it prints one line, `low high`, which is its guess for the range that contains the **next** number. A narrow range that hits scores more than a wide range that hits. A miss scores nothing.

This is the last project of my algorithm chain: `math-skills` → `guess-it-1` → `linear-stats` → `guess-it-2`.

## Requirements

- Go 1.21 or newer

## Build and run

Clone the repository and go into the project folder:

```bash
git clone <repository-url>
cd guess-it-2
```

Build the program:

```bash
go build -o guess-it-2 .
```

Run it with a few numbers (one per line):

```bash
printf '189\n113\n121\n114\n145\n110\n130\n' | ./guess-it-2
```

Run it with a whole data file:

```bash
cat data.txt | ./guess-it-2
```

The program prints one line per input number. Each line has two integers, `low high`: the range where the next number is expected. The first few lines come from a short warm-up phase, where there is too little data for statistics.

Run the unit tests:

```bash
go test ./...
```

## How it is scored

The official tester gives, for every correct guess:

```
round( 10 000 000 / (1 + high - low) / (n - 1) )
```

where `n` is the number of values in the data set. A wrong guess gives 0 points.

Two consequences shaped the design:

1. Narrower ranges earn more points per hit, but they hit less often.
2. The points per hit are rounded to an integer, so the width of the range matters in small steps.

## How it works

For every new number, `internals.PredictRange` does this:

1. **Window.** It keeps only the last 1000 values. Older values matter less and cost time.
2. **Outliers.** It removes outliers with the IQR rule, but keeps the original position of every remaining point.
3. **Line.** It fits a least-squares line and computes the Pearson correlation `r`.
4. **Trend or noise.** A t-test on `r` decides whether the trend is real.
   - Real trend: predict the next value on the line.
   - No trend: predict the mean.
5. **Range.**
   - Real trend: a fixed range of **94 integers** centred on the prediction.
   - No trend: prediction ± the expected error.

Why 94? On the test data the noise is spread evenly over about 101 integers around the line. With 12 500 values the tester pays `round(800.06 / width)` points per hit. A width of 94 pays 9 points and still hits about 93% of the time. A width of 95 or more pays only 8 points. This number is tuned to this data and this scoring rule (see `trendRangeCount` in `internals/predict.go`).

## Project structure

```
guess-it-2/
├── main.go                  reads stdin, prints one range per input number
├── script.sh                how the tester starts the program
├── go.mod
├── internals/
│   ├── predict.go           PredictRange and all tunable constants
│   ├── predict_test.go      unit tests
│   └── stats.go             line fit, Pearson r, quantiles, IQR filter
└── docs/
    ├── PRD.md               product requirements
    └── TASK_CARDS.md        task breakdown
```

## Testing with the official tester

From the root folder of the tester:

```bash
mkdir -p student
cp /path/to/guess-it-2/guess-it-2 student/guess-it-2
cp /path/to/guess-it-2/script.sh student/script.sh
docker compose up --build
```

Then open `http://localhost:3000`, choose an opponent in the URL (for example `?guesser=linear-regr`) and press **Test Data 4** or **Test Data 5**. The tester picks one of the five files of that data set at random each time, so run each test several times.

## Notes

- The program handles blank lines and non-numeric lines by skipping them.
- It never prints `NaN` and never prints a range where `low > high`.
- The program flushes its output after every line, because the tester waits for an answer after every number.