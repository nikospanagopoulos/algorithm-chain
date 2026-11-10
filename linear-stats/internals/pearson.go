package internals

import "math"

// PearsonCorrelation calculates the Pearson Correlation Coefficient (r)
// which measures the strength of the linear relationship between x and y
func PearsonCorrelation(numbers []float64) float64 {
	n := float64(len(numbers))

	// x = index of line (0, 1, 2...)
	// y = value of line
	var sumX, sumY, sumXY, sumXX, sumYY float64
	for i, y := range numbers {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
		sumYY += y * y
	}

	// Calculate Pearson Correlation Coefficient
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumXX - sumX*sumX) * (n*sumYY - sumY*sumY))

	return numerator / denominator
}
