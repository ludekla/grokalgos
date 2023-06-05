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
	nb.points = points
	nb.values = values
}

func (nb NeighboursRegressor) Predict(point Point) float64 {
	nb.pivot = point
	if len(nb.points) == 0 {
		panic("no data to work on")
	}
	sort.Sort(nb)
	var sum float64
	for _, val := range nb.values[:nb.k] {
		sum += val
	}
	return sum / float64(nb.k)
}

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
