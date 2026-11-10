package internals

import (
	"math"
	"math/rand"
	"testing"
)

func TestTinyInputsDoNotPanic(t *testing.T) {
	data := []float64{189, 113, 121, 114, 145, 110, 130}
	for n := 1; n <= len(data); n++ {
		low, high := PredictRange(data[:n])
		if low > high {
			t.Fatalf("n=%d: low %d > high %d", n, low, high)
		}
	}
}

func TestConstantDataGivesValidRange(t *testing.T) {
	data := make([]float64, 100)
	for i := range data {
		data[i] = 150
	}
	low, high := PredictRange(data)
	if low > 150 || high < 150 || low >= high {
		t.Fatalf("expected a range around 150, got %d %d", low, high)
	}
}

func TestPerfectTrendIsFollowed(t *testing.T) {
	data := make([]float64, 60)
	for i := range data {
		data[i] = 2*float64(i) + 10
	}
	low, high := PredictRange(data)
	// The true next number of the series (index 60) is 2*60 + 10 = 130.
	if float64(low) > 130 || float64(high) < 130 {
		t.Fatalf("trend not followed: range %d %d does not contain 130", low, high)
	}
}

func TestNeverNaNOrInverted(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	var data []float64
	for i := 0; i < 500; i++ {
		data = append(data, 100+rng.Float64()*100)
		low, high := PredictRange(data)
		if low > high {
			t.Fatalf("i=%d: low %d > high %d", i, low, high)
		}
	}
}

func TestIqrKeepsOriginalPositions(t *testing.T) {
	pts := []Point{{0, 10}, {1, 11}, {2, 1000}, {3, 12}, {4, 11}, {5, 10}}
	kept := IqrFilter(pts, 1.6)
	if len(kept) != 5 {
		t.Fatalf("expected 5 points after removing the outlier, got %d", len(kept))
	}
	// The point after the removed outlier must still have x = 3, not 2.
	if kept[2].X != 3 {
		t.Fatalf("x position was squeezed: got %v, want 3", kept[2].X)
	}
}

func TestFitLineMatchesKnownLine(t *testing.T) {
	pts := []Point{{0, 1}, {1, 3}, {2, 5}, {3, 7}}
	f := FitLine(pts)
	if math.Abs(f.A-2) > 1e-9 || math.Abs(f.B-1) > 1e-9 || math.Abs(f.R-1) > 1e-9 {
		t.Fatalf("got a=%v b=%v r=%v, want 2 1 1", f.A, f.B, f.R)
	}
}

func TestFitLineSinglePointIsSafe(t *testing.T) {
	f := FitLine([]Point{{0, 5}})
	if math.IsNaN(f.A) || math.IsNaN(f.B) || math.IsNaN(f.R) {
		t.Fatalf("NaN on a single point: %+v", f)
	}
}
