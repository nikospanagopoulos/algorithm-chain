# Task-Card

## Setup and Planning

##### Goal:
_The project structure is ready and all edge cases from the PRD are documented._
##### Depends on:
_None_

### Acceptance Criteria

- Project folder and files exist (main.go, go.mod, statistics/, docs/, tests/)
- PRD is finalized and stored in docs/
- CLI contract is clearly understood:
```bash
go run main.go data.txt
```
- List of edge cases is written (empty file, single value, invalid data, negative numbers)

### Tests to write first

- TestProjectSetup: proves project builds and runs

### Notes

- Focus on understanding the PRD fully before coding
- Edge cases will guide future implementation

## File Reading and Parsing

##### Goal:
_The program can read the input file line by line and store the values into a slice of integers._
##### Depends on:
- Setup and Planning
### Acceptance Criteria

- Program accepts data.txt as CLI argument
- File is read line by line correctly
- Each line is parsed into an integer and stored in a slice
- Negative numbers are parsed correctly
- If no argument is provided, program prints usage message and exits
- If file is not found, program prints error message and exits
- If file contains invalid data, program prints error message and exits

### Tests to write first

- TestFileReading: proves file is read correctly
- TestParsing: proves values are stored correctly in a slice
- TestNegativeNumbers: proves negative values are parsed correctly
- TestMissingArgument: proves usage message prints on missing argument
- TestFileNotFound: proves error message prints on missing file
- TestInvalidData: proves error message prints on invalid data

### Notes

- Use bufio.Scanner to read line by line
- Use strconv.Atoi to parse each line into an integer
- Handle all error cases before proceeding to calculations


## Implement Calculations

##### Goal:

_The program correctly calculates Average, Median, Variance, and Standard Deviation from the parsed dataset._

##### Depends on:

_File Reading and Parsing_

### Acceptance Criteria

- Average returns the correct rounded integer
- Median returns the correct rounded integer for both odd and even datasets
- Variance returns the correct rounded integer using the Average result
- Standard Deviation returns the correct rounded integer using the Variance result
- All functions handle negative numbers correctly
- All functions handle a single value correctly (Variance: 0, StdDev: 0)
- All functions handle identical values correctly (Variance: 0, StdDev: 0)


### Tests to write first

- TestAverage: proves Average is calculated correctly
- TestMedianOdd: proves Median works for odd datasets
- TestMedianEven: proves Median works for even datasets
- TestVariance: proves Variance is calculated correctly
- TestStandardDeviation: proves Standard Deviation is calculated correctly
- TestSingleValue: proves all functions handle single value correctly
- TestIdenticalValues: proves all functions handle identical values correctly
- TestNegativeValues: proves all functions handle negative numbers correctly.

### Notes

- Each function lives in its own file inside statistics/
- Variance uses the Average result — pass it as argument
- Standard Deviation uses the Variance result — pass it as argument
- Use math.Round for rounding


## Print Results

##### Goal:
_The program prints all four results in the correct format, each on a separate line._

##### Depends on:
_Implement Calculations_

### Acceptance Criteria

- Average prints correctly: Average: x
- Median prints correctly: Median: x
- Variance prints correctly: Variance: x
- Standard Deviation prints correctly: Standard Deviation: x
- Each result prints on a separate line
- All values are rounded integers

### Tests to write first

- TestPrintFormat: proves output matches expected format
- TestPrintRounding: proves all values are correctly rounded before printing

### Notes

- Use fmt.Printf for formatted output
- Order of printing must follow: Average → Median → Variance → Standard Deviation


## Integration and Final Testing

##### Goal:
_All modules work together correctly and all audit cases and edge cases pass._
##### Depends on:
_Print Results_

### Acceptance Criteria

- All audit cases pass
- All extra golden tests pass
- Negative numbers are handled correctly end-to-end
- Single value dataset works correctly end-to-end
- Identical values dataset works correctly end-to-end
- Empty file is handled correctly
- Invalid data is handled correctly
- Output format is exactly correct

### Tests to write first

- TestFullPipeline: proves full flow works correctly
- TestAuditCases: proves all audit cases pass
- TestGoldenTests: proves all extra golden tests pass
- TestEdgeCases: proves all edge cases are handled correctly

### Notes

- Run all tests before submitting
- Check output format carefully — each result on a separate line
- Bugs often come from interaction between modules
- Validate full input/output flow end-to-end