# algorithm-chain

A chain of four Go projects about statistics and prediction. Each project builds on the one before it: the maths from the first one is reused in the last one.

```
math-skills  →  guess-it-1  →  linear-stats  →  guess-it-2
```

| Project | What it does | Builds on |
|---|---|---|
| [`math-skills`](./math-skills) | Reads a data file and prints basic statistics: average, median, variance and standard deviation. | none |
| [`guess-it-1`](./guess-it-1) | Reads numbers one by one and, after each one, prints a range `low high` where it expects the next number. First version, based on simple statistics. | `math-skills` |
| [`linear-stats`](./linear-stats) | Reads a data file and prints the linear regression line and the Pearson correlation coefficient. | `math-skills` |
| [`guess-it-2`](./guess-it-2) | Improved predictor. It uses the line fit and the correlation from `linear-stats`, a window of recent values, outlier filtering and a range width tuned to the scoring rule. | `guess-it-1`, `linear-stats` |

## How the projects connect

1. **`math-skills`** gives the basic tools: mean, variance and standard deviation.
2. **`guess-it-1`** uses them to guess a range for the next number.
3. **`linear-stats`** adds the least-squares line and the Pearson correlation `r`. They tell us if the data has a trend.
4. **`guess-it-2`** puts everything together. It follows the trend when `r` shows that it is real, removes outliers with the IQR rule, and chooses the width of the range from the scoring formula of the tester.

## Repository structure

```
algorithm-chain/
├── README.md              this file
├── .gitignore
├── math-skills/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── statistics/
│   ├── tests/
│   └── docs/
├── guess-it-1/
│   ├── README.md
│   └── student/           code, go.mod and script.sh
├── linear-stats/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── internals/
│   └── docs/
└── guess-it-2/
    ├── README.md
    ├── go.mod
    ├── main.go
    ├── script.sh
    ├── internals/
    └── docs/
```

Each folder is an independent Go module. To work on one project, go into its folder:

```bash
cd guess-it-2
go build -o guess-it-2 .
go test ./...
```

Note: in `guess-it-1` the Go module is inside the `student` folder, so run `cd guess-it-1/student`.

See the README of each project for its own usage.

## Requirements

- Go 1.21 or newer
- Docker, only to run the official tester of the `guess-it` projects
