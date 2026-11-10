# PRD — math-skills

## Problem Statement

A command-line tool that reads a dataset from a file and calculates the following statistical measures: Average, Median, Variance, and Standard Deviation. Results are printed to the terminal as rounded integers.

## CLI Contract

### Command
```bash
go run main.go data.txt
```

### Arguments
``` 
- <inputFile>: Path to the input file containing the dataset
```

### Behavior
```
The program reads the content of <inputFile>, calculates all statistical measures, and prints the results to the terminal.
```
### Error Handling

- If the number of arguments is incorrect, the program must terminate and display a usage message.
- If the input file cannot be found or opened, the program must terminate with an error message.
- If the file contains invalid or non-numeric data, the program must terminate with an error message.

## Calculation Functions

### Average

The sum of all values divided by the count.
```
Average = Σx / N
```
### Median

The middle value after sorting the dataset. If the count is even, the average of the two middle values.
```
Odd:  Median = x[N/2]
Even: Median = (x[N/2-1] + x[N/2]) / 2
```
### Variance

A measure of how spread out the values are around the average.
```
Variance = Σ(x - Average)² / N
```
### Standard Deviation

The square root of the variance.
```
StdDev = √Variance
```

## Non-Goals

- The program does not support multiple input files.
- It does not handle multilingual or Unicode characters, only basic ASCII.
- It does not export results to a file, only prints to the terminal.
- It does not support interactive input, only file-based input via CLI.
- It does not validate whether the dataset is statistically meaningful (e.g. single value).

## Acceptance Criteria

### Audit Cases

#### Audit case 1
The program correctly calculates the Average and prints it as a rounded integer.
```
Input file:
4
8
15
16
23
42

Output:
Average: 18
```
#### Audit case 2
The program correctly calculates the Median and prints it as a rounded integer.
```
Input file:
4
8
15
16
23
42

Output:
Median: 16
```
#### Audit case 3
The program correctly calculates the Variance and prints it as a rounded integer.
```
Input file:
4
8
15
16
23
42

Output:
Variance: 152
```
#### Audit case 4
The program correctly calculates the Standard Deviation and prints it as a rounded integer.
```
Input file:
4
8
15
16
23
42

Output:
Standard Deviation: 12
```

### Extra Golden Tests

#### Extra test 1
All four results print in the correct format.
```
Input file:
4
8
15
16
23
42

Output:
Average: 18
Median: 16
Variance: 152
Standard Deviation: 12
```
#### Extra test 2
Handles edge case where dataset has a single value.
```
Input file:
42

Output:
Average: 42
Median: 42
Variance: 0
Standard Deviation: 0
```
#### Extra test 3
Handles edge case where all values are identical.
```
Input file:
5
5
5
5

Output:
Average: 5
Median: 5
Variance: 0
Standard Deviation: 0
```

## Architecture

### Sequential Pipeline
Each calculation depends on the previous result — the Median requires sorted data, the Variance requires the Average, and the Standard Deviation requires the Variance. Therefore all steps execute in sequence, each passing its result to the next.
### Why this approach

Each function is independent and testable separately
The flow is linear and easy to understand
Each step depends on the output of the previous one

### Tradeoffs we accept

All data must be loaded into memory before any calculation begins
Each step must wait for the previous one to finish
```
Sketch:
Input file
    |
    v
Read & Parse (line by line)
    |
    v
Calculate Average
    |
    v
Calculate Median
    |
    v
Calculate Variance (uses Average)
    |
    v
Calculate Standard Deviation (uses Variance)
    |
    v
Print:
Average: x
Median: x
Variance: x
Standard Deviation: x
```

## Milestones
### Setup and Planning

- Create project skeleton, PRD, and task cards.
- ✅ Testable: Project builds and runs.

### File Reading and Parsing

- Read the input file line by line and parse values into a slice.
- ✅ Testable: File is read correctly and values are stored.

### Implement Calculations

- Implement Average, Median, Variance, Standard Deviation.
- ✅ Testable: Each function returns correct rounded integer.

### Print Results

- Print all four results in the correct format.
- ✅ Testable: Output matches expected format.

### Integration and Final Testing

- Combine all modules and run all audit cases and edge cases.
- ✅ Testable: All audit cases and golden tests pass.

## Risks / Open Questions

- Empty file: What happens if the input file has no data?
- Single value: How should Variance and Standard Deviation behave with only one value?
- Invalid data: What happens if a line contains non-numeric characters?
- Large dataset: Will processing a very large file cause memory or performance issues?
- Negative numbers: The program must handle negative values correctly in all calculations.