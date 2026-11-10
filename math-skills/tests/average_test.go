package test

import (
	"math-skills/statistics"
	"testing"
)

func TestAverage(t *testing.T) {

	result := statistics.Average([]int{1, 2, 3, 4, 5})
	expected := 3.0

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
