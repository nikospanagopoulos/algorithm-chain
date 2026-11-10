package internals

import (
	"math"
	"sort"
)

// Point is one data value together with its ORIGINAL position on the x axis.
// Keeping x inside the point is what lets us drop outliers without
// "squeezing" the remaining values together (the guess-it-1 weakness).
type Point struct {
	X, Y float64
}

// Fit holds everything we need from one least-squares line fit.
type Fit struct {
	N      int     // number of points used
	A, B   float64 // the line: y = A*x + B
	R      float64 // Pearson correlation coefficient between x and y
	XMean  float64
	YMean  float64
	Sxx    float64 // sum of (x - xMean)^2
	ResStd float64 // standard deviation of the residuals (the errors of the line)
	YStd   float64 // standard deviation of y alone (the errors of the plain mean)
}

// FitLine computes the regression line, Pearson r and the two error measures.
//
// Same maths as linear-stats, written in "deviation" form
// (sum of (x-xMean)(y-yMean) ...) which is numerically safer than the
// n*sumXY - sumX*sumY form when the values are big.
// Every division is guarded, so a flat or tiny input never produces NaN.
func FitLine(pts []Point) Fit {
	n := len(pts)
	if n == 0 {
		return Fit{}
	}
	fn := float64(n)

	var sumX, sumY float64
	for _, p := range pts {
		sumX += p.X
		sumY += p.Y
	}
	xMean := sumX / fn
	yMean := sumY / fn

	var sxx, syy, sxy float64
	for _, p := range pts {
		dx := p.X - xMean
		dy := p.Y - yMean
		sxx += dx * dx
		syy += dy * dy
		sxy += dx * dy
	}

	f := Fit{N: n, XMean: xMean, YMean: yMean, Sxx: sxx}

	if sxx > 0 {
		f.A = sxy / sxx
	}
	f.B = yMean - f.A*xMean

	if sxx > 0 && syy > 0 {
		f.R = sxy / math.Sqrt(sxx*syy)
	}

	if n > 1 {
		f.YStd = math.Sqrt(syy / (fn - 1))
	}

	// Residual sum of squares = syy - a*sxy  (identity for least squares).
	rss := syy - f.A*sxy
	if rss < 0 {
		rss = 0
	}
	if n > 2 {
		f.ResStd = math.Sqrt(rss / (fn - 2))
	} else {
		f.ResStd = f.YStd
	}

	return f
}

// Quantile returns the q-th quantile (0..1) of the values, using linear
// interpolation between neighbours. It sorts a copy with the standard library.
func Quantile(values []float64, q float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)

	index := q * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	if lower == upper {
		return sorted[lower]
	}
	return sorted[lower] + (index-float64(lower))*(sorted[upper]-sorted[lower])
}

// IqrFilter removes outliers with the Interquartile Range rule:
// values outside [Q1 - k*IQR, Q3 + k*IQR] are dropped.
// The surviving points keep their original X.
func IqrFilter(pts []Point, k float64) []Point {
	if len(pts) < 4 {
		return pts
	}
	ys := make([]float64, len(pts))
	for i, p := range pts {
		ys[i] = p.Y
	}
	q1 := Quantile(ys, 0.25)
	q3 := Quantile(ys, 0.75)
	iqr := q3 - q1

	lo := q1 - k*iqr
	hi := q3 + k*iqr

	kept := make([]Point, 0, len(pts))
	for _, p := range pts {
		if p.Y >= lo && p.Y <= hi {
			kept = append(kept, p)
		}
	}
	return kept
}
