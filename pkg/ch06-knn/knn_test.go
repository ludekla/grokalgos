package ch06

import (
	"testing"
)

func TestNeighboursRegression(t *testing.T) {
	points := []Point{
		Point{5, 1, 0}, Point{1, 1, 0}, Point{3, 1, 1},
		Point{4, 0, 0}, Point{4, 0, 1}, Point{2, 0, 0},
	}
	vals := []float64{300, 75, 225, 150, 200, 50}
	nbs := NewNeighboursRegressor(4)
	nbs.Fit(points, vals)
	pt := []Point{{4, 1, 0}}
	expected := 218.75
	got := nbs.Predict(pt)
	if got[0] != expected {
		t.Errorf("expected %f, got %f", expected, got)
	}
}
