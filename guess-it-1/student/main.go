package main

import (
	"bufio"
	"fmt"
	"guess-it-1/internals"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}

		numbers = append(numbers, num)
		low, high := internals.PredictRange(numbers)
		fmt.Printf("%d %d\n", low, high)
	}
}
