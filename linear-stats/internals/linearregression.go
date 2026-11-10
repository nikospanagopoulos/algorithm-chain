package internals

// LinearRegression calculates the slope (a) and intercept (b)
// of the best fitting line y = ax + b using least squares method
func LinearRegression(numbers []float64) (float64, float64) {
	n := float64(len(numbers))

	// x = index of line (0, 1, 2...)
	// y = value of line
	var sumX, sumY, sumXY, sumXX float64
	for i, y := range numbers {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	// Calculate slope (a) and intercept (b)
	a := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - a*sumX) / n

	return a, b
}
