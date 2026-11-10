package test

import (
	"math-skills/statistics"
	"testing"
)

func TestVariance(t *testing.T) {
	result := statistics.Variance([]int{1, 2, 3, 4, 5})
	expected := 2.0

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
