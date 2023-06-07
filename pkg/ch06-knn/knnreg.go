package ch06

import (
	"sort"
)

type Point []float64

type NeighboursRegressor struct {
	k      int
	points []Point
	values []float64
	pivot  Point
}

func NewNeighboursRegressor(nNeighbours int) NeighboursRegressor {
	return NeighboursRegressor{k: nNeighbours}
}

func (nb *NeighboursRegressor) Fit(points []Point, values []float64) {
	if len(points) != len(values) {
		panic("inconsistent data")
	}
	nb.points = make([]Point, len(points))
	nb.values = make([]float64, len(values))
	size := len(points[0])
	for i, point := range points {
		nb.points[i] = make(Point, size)
		copy(nb.points[i], point)
	}
	copy(nb.values, values)
}

func (nb NeighboursRegressor) Predict(points []Point) []float64 {
	if len(nb.points) == 0 {
		panic("no data to work on")
	}
	labels := make([]float64, len(points))
	for i, point := range points {
		nb.pivot = point
		sort.Sort(nb)
		labels[i] = mean(nb.values[:nb.k])
	}
	return labels
}

// Coefficient of Determination
func (nb NeighboursRegressor) Score(points []Point, values []float64) float64 {
	// total sum of squares
	tss := variance(values)
	preds := nb.Predict(points)
	rss := mse(values, preds)
	return 1.0 - rss / tss
}

// Implementation of sort.Interface
func (nb NeighboursRegressor) Len() int {
	return len(nb.points)
}

func (nb NeighboursRegressor) distance(j int) float64 {
	var sum float64
	point := nb.points[j]
	for i, elem := range nb.pivot {
		sum += (elem - point[i]) * (elem - point[i])
	}
	return sum
}

func (nb NeighboursRegressor) Less(i, j int) bool {
	return nb.distance(i) < nb.distance(j)
}

func (nb NeighboursRegressor) Swap(i, j int) {
	nb.points[i], nb.points[j] = nb.points[j], nb.points[i]
	nb.values[i], nb.values[j] = nb.values[j], nb.values[i]
}
