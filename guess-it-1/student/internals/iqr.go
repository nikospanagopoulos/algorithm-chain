package internals

import "math"

// IqrFilter removes outliers using the Interquartile Range method
// Values outside Q1 - 1.6*IQR and Q3 + 1.6*IQR are removed
func IqrFilter(numbers []float64) []float64 {
	// Need at least 4 numbers to calculate IQR
	if len(numbers) < 4 {
		return numbers
	}

	q1 := Quantile(numbers, 0.25)
	q3 := Quantile(numbers, 0.75)
	iqr := q3 - q1

	// Keep only values within the IQR bounds
	var filtered []float64
	for _, num := range numbers {
		if num >= q1-1.6*iqr && num <= q3+1.6*iqr {
			filtered = append(filtered, num)
		}
	}

	return filtered
}

// Quantile calculates the q-th quantile of a slice of numbers
// using linear interpolation between adjacent values
func Quantile(numbers []float64, q float64) float64 {
	// Sort a copy of the slice using bubble sort
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	// Calculate the index and interpolate if needed
	index := q * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))

	if lower == upper {
		return sorted[lower]
	}
	return sorted[lower] + (index-float64(lower))*(sorted[upper]-sorted[lower])
}
