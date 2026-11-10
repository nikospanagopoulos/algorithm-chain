# Task Cards — linear-stats
## Task 1 — Setup

### Goal: Project skeleton ready to build and run.
```
- Create linear-stats/ folder
- Create main.go with package main
- Create go.mod with go mod init linear-stats
- ✅ Testable: go run main.go compiles
```
## Task 2 — File Reading

## Goal: Read and parse the data file correctly.
```
- Accept filename as os.Args[1]
- Read line by line with bufio.Scanner
- Parse each line as float64
- Store in []float64 slice
- ✅ Testable: Prints correct number of values read
```
## Task 3 — Linear Regression

### Goal: Calculate slope (a) and intercept (b).
```
- x = index of line (0, 1, 2...)
- y = value of line
- Calculate sumX, sumY, sumXY, sumXX
- Apply formula for a and b
- ✅ Testable: Output matches reference program
```
## Task 4 — Pearson Correlation Coefficient

### Goal: Calculate r with correct formula.
```
- Reuse sumX, sumY, sumXY, sumXX from Task 3
- Add sumYY = Σy²
- Apply Pearson formula
- ✅ Testable: Output matches reference program
```
## Task 5 — Output Formatting

### Goal: Print results in exact format required.
```
- Linear Regression: 6 decimal places
- Pearson: 10 decimal places
- ✅ Testable: Format matches audit requirements
```
## Task 6 — Integration & Testing

### Goal: All audit cases pass.
```
- Run reference program and student program on same data
- Compare outputs for 4 datasets
- ✅ Testable: All values match
```