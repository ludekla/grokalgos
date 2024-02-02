package main

import (
	"fmt"

	"grokalgos/pkg/ch06-knn"
)

func main() {

	points := []ch06.Point{
		{1, 1, 0}, {1, 0.9, 0.2}, {3.2, 1, 6.2},
		{0.1, 0.2, 0}, {4.1, 1.3, 5.8}, {2, 0, 0},
	}
	vals := []float64{20, 25, 225, 10, 200, 9.3}
	nbs := ch06.NewNeighboursRegressor(3)
	nbs.Fit(points, vals)
	preds := nbs.Predict(points)
	for i, point := range points {
		fmt.Printf("%v: %.3f (gold: %.3f)\n", point, preds[i], vals[i])
	}
	fmt.Printf("Score on points: %.4f\n", nbs.Score(points, vals))

}
