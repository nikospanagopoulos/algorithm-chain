# Guess It

A command-line tool that reads numbers from stdin and predicts the range of
the next number using linear regression with IQR outlier filtering.

## Usage
```bash
echo "100\n150\n200" | ./student/guess-it-1
```
or via script:
```bash
echo "100\n150\n200" | ./student/script.sh
```

## Example
Input: `100`, `150`, `200`
Output:

90 110
120 200
181 219

Each line is the predicted `low high` range for the next number.

## How it works
1. **Sliding window** — keeps the last 60 numbers for prediction
2. **IQR filter** — removes outliers (multiplier: 1.6)
3. **Linear regression** — fits a line to the filtered data
4. **Range output** — prediction ± 19

## Build
```bash
cd student && go build -o guess-it-1 .
```

## Run with Docker
```bash
cd guess-it-dockerized && docker-compose up --build
```

## What I learned
Applying basic statistics (IQR, linear regression) to a streaming input
problem, and packaging the same tool for both native and Docker execution.