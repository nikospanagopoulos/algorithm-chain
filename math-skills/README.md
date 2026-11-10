# Math Skills

A command-line tool that reads a dataset from a file and calculates Average,
Median, Variance, and Standard Deviation, printed to the terminal as rounded
integers.

## Usage
```bash
go run main.go <inputFile>
```
or, after building:
```bash
./bin/math-skills <inputFile>
```

## Example
`data.txt`:
```
4
8
15
16
23
42
```
```bash
go run main.go data.txt
```
Output:

Average: 18
Median: 16
Variance: 152
Standard Deviation: 12


## Testing
```bash
go test ./tests/
```

## Build
```bash
go build -o bin/math-skills main.go
```

## What I learned
Implementing core statistical formulas from scratch, structuring a Go
project into clean packages, and writing test coverage for numerical logic.
