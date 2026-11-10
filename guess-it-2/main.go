package main

import (
	"bufio"
	"fmt"
	"guess-it-2/internals"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

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

		// The tester waits for the answer after every number,
		// so we flush after every line.
		fmt.Fprintf(writer, "%d %d\n", low, high)
		writer.Flush()
	}
}
