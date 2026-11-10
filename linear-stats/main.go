package main

import (
	"bufio"
	"fmt"
	"linear-stats/internals"
	"os"
	"strconv"
)

func main() {
	// Check if a file argument was provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <inputFile>")
		return
	}

	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and parse the file line by line
	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		numbers = append(numbers, num)
	}

	// Calculate and print results
	a, b := internals.LinearRegression(numbers)
	r := internals.PearsonCorrelation(numbers)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
