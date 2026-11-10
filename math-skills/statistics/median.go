package statistics

// Median calculates the middle value of a sorted dataset.
// For even datasets, it returns the average of the two middle values.
func Median(numbers []int) float64 {
	// Return 0 if the dataset is empty
	if len(numbers) == 0 {
		return 0
	}
	// Create a copy to avoid modifying the original slice
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	// Sort the numbers using bubble sort
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	// If odd, return the middle element
	if len(sorted)%2 == 1 {
		return float64(sorted[len(sorted)/2])
	}
	// If even, return the average of the two middle elements
	return float64(sorted[len(sorted)/2-1]+sorted[len(sorted)/2]) / 2
}
