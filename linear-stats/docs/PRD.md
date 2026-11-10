# PRD — linear-stats
### Problem Statement

A command-line tool that reads a dataset from a file and calculates the Linear Regression Line and the Pearson Correlation Coefficient. Results are printed to the terminal with specific decimal places.

## CLI Contract
### Command
```
go run main.go data.txt
```
### Arguments
```
- <inputFile>: Path to the input file containing the dataset
```
### Output Format
```
Linear Regression Line: y = <6 decimal places>x + <6 decimal places>
Pearson Correlation Coefficient: <10 decimal places>
```
## Calculation Functions
### Linear Regression Line
```
a = (n·ΣXY - ΣX·ΣY) / (n·ΣXX - ΣX·ΣX)
b = (ΣY - a·ΣX) / n
y = ax + b
```
### Pearson Correlation Coefficient
```
r = (n·ΣXY - ΣX·ΣY) / √((n·ΣXX - ΣX²) · (n·ΣYY - ΣY²))
```
## Non-Goals
- Does not support multiple input files
- Does not export results to a file
- Does not support interactive input
### Acceptance Criteria
#### Audit Case 1

Output is in the correct format.
```
Linear Regression Line: y = 2.123456x + 1.123456
Pearson Correlation Coefficient: 0.1234567890
```
#### Audit Case 2

Values of Linear Regression Line have exactly 6 decimal places.

#### Audit Case 3

Pearson Correlation Coefficient has exactly 10 decimal places.

#### Audit Case 4

Values match the reference program for 4 different datasets.

Architecture
```
Input file
    |
    v
Read & Parse (line by line)
    |
    v
Calculate Linear Regression (a, b)
    |
    v
Calculate Pearson Correlation Coefficient (r)
    |
    v
Print results
```

## Project Structure
```
linear-stats/
├── internals/
│   ├── linearregression.go   ← calculates slope (a) and intercept (b)
│   └── pearson.go            ← calculates Pearson correlation coefficient (r)
├── main.go                   ← reads file, calls functions, prints results
├── go.mod
└── README.md
```