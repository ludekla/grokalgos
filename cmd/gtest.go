package main

import (
	"fmt"
	"math"
	"sort"
)

type Point []float64

type NeighboursRegressor struct {
	k int
	points []Point
	values []float64
	pivot Point
}

func NewNeighboursRegressor(k int) NeighboursRegressor {
	return NeighboursRegressor{k, nil, nil, nil}
}

func (nb *NeighboursRegressor) Fit(points []Point, values []float64) {
	nb.points = points
	nb.values = values
}

func (nb NeighboursRegressor) Predict(point Point) float64 {
	nb.pivot = point
	if nb.Len() == 0 {
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
	return math.Sqrt(sum)
}

func (nb NeighboursRegressor) Less(i, j int) bool {
	return nb.distance(i) < nb.distance(j)
}

func (nb NeighboursRegressor) Swap(i, j int) {
	nb.points[i], nb.points[j] = nb.points[j], nb.points[i]
	nb.values[i], nb.values[j] = nb.values[j], nb.values[i]
}

func main() {

	points := []Point{
		Point{5, 1, 0}, Point{1, 1, 0}, Point{3, 1, 1}, 
		Point{4, 0, 0}, Point{4, 0, 1}, Point{2, 0, 0},
	}
	vals := []float64{300, 75, 225, 150, 200, 50}
	nbs := NewNeighboursRegressor(4)
	nbs.Fit(points, vals)
	pt := Point{4, 1, 0}
	fmt.Printf("%v: %0.3f\n", pt, nbs.Predict(pt))

}