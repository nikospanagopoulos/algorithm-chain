package internals

import "math"

func PredictRange(numbers []float64) (int, int) {
	// If we have fewer than 2 numbers, return a range around the first value
	if len(numbers) < 2 {
		center := int(numbers[0])
		return center - 10, center + 10
	}

	// Sliding window - keep the last 60 numbers
	start := 0
	if len(numbers) > 60 {
		start = len(numbers) - 60
	}
	window := numbers[start:]

	// Filter outliers using IQR
	filtered := IqrFilter(window)
	if len(filtered) < 2 {
		filtered = window
	}

	// Linear regression - find the best fitting line
	n := float64(len(filtered))
	var sumX, sumY, sumXY, sumXX float64
	for i, y := range filtered {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	// Calculate slope (a) and intercept (b) of the line y = ax + b
	a := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - a*sumX) / n

	// Predict the next number using 100% linear regression
	prediction := a*n + b

	// Range ±19 around the prediction
	offset := 19.0
	low := int(math.Round(prediction - offset))
	high := int(math.Round(prediction + offset))

	return low, high
}
