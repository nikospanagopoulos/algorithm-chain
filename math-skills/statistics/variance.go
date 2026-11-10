package statistics

// Variance calculates how spread out the values are around the average.
// It returns the average of the squared differences from the mean.
func Variance(numbers []int) float64 {
	// Return 0 if the dataset is empty
	if len(numbers) == 0 {
		return 0
	}
	// Calculate the average using the Average function
	avg := Average(numbers)
	// Declare a float64 to store the sum of squared differences
	var sum float64
	for _, num := range numbers {
		// Calculate the difference between each value and the average
		diff := float64(num) - avg
		// Square the difference and add it to the sum
		sum += diff * diff
	}
	// Divide by the count to get the average of squared differences
	return sum / float64(len(numbers))
}
