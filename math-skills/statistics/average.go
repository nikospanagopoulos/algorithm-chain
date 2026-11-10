package statistics

// Average calculates the mean of a dataset.
// It returns the sum of all values divided by the count.
func Average(numbers []int) float64 {
	// Return 0 if the dataset is empty
	if len(numbers) == 0 {
		return 0
	}
	// Sum all values in the dataset
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	// Divide by count using float64 to preserve decimal precision
	return float64(sum) / float64(len(numbers))
}
