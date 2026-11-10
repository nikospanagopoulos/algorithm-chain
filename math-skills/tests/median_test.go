package test

import (
	"math-skills/statistics"
	"testing"
)

func TestMed(t *testing.T) {
	result := statistics.Median([]int{1, 2, 3, 4, 5})
	expected := 3.0

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result2 := statistics.Median([]int{1, 2, 3, 4})

	expected2 := 2.5

	if result2 != expected2 {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}
}
