# linear-stats

A command-line tool that reads a dataset from a file and calculates the Linear Regression Line and the Pearson Correlation Coefficient.

## Usage

```bash
go run main.go <inputFile>
```

## Example

Given a file `data.txt`:
```
189
113
121
114
145
110
```

Output:

Linear Regression Line: y = -8.742857x + 153.857143
Pearson Correlation Coefficient: -0.5330331012


## How to Test

Download the reference binary from [here](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip), then run:

```bash
./bin/linear-stats data.txt
```

Then run the student program with the same file:

```bash
go run main.go data.txt
```

Compare the outputs — they should match.

## Structure
```
linear-stats/
├── internals/
│ ├── linearregression.go
│ └── pearson.go
├── main.go
├── go.mod
└── README.md
```

## Build

```bash
go build -o linear-stats main.go
```