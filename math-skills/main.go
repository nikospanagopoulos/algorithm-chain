package main

import (
	"bufio"
	"fmt"
	"math"
	"math-skills/statistics"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		// Check if exactly one argument is provided
		fmt.Println("Usage: go run main.go <inputFile>")
		os.Exit(1)
	} // Store the filename from the command line argument
	filename := os.Args[1]
	// Open the file and handle any errors
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: could not open file:", err)
		os.Exit(1)
	}
	// Ensure the file is closed when the function ends
	defer file.Close()
	// Declare a slice to store the parsed integers
	var numbers []int
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the current line as a string
		line := scanner.Text()
		// Convert the line to an integer and handle invalid data
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error: invalid data:", line)
			os.Exit(1)
		}
		// Append the parsed integer to the slice
		numbers = append(numbers, number)
	}
	// Calculate each statistic and round to the nearest integer
	avg := statistics.Average(numbers)
	fmt.Printf("Average: %v\n", int(math.Round(avg)))
	median := statistics.Median(numbers)
	fmt.Printf("Median: %v\n", int(math.Round(median)))
	variance := statistics.Variance(numbers)
	fmt.Printf("Variance: %v\n", int(math.Round(variance)))
	stdDev := math.Sqrt(variance)
	fmt.Printf("Standard Deviation: %v\n", int(math.Round(stdDev)))
}
