package ch05

import (
	"fmt"
	"math"
)

type CellGrid [][]int

func NewCellGrid(nrows, ncols int) CellGrid {
	cell := make(CellGrid, nrows)
	for r := 0; r < nrows; r++ {
		cell[r] = make([]int, ncols)
	}
	return cell
}

func (cg CellGrid) Max() int {
	var max int = math.MinInt
	for _, row := range cg {
		for _, c := range row {
			if c > max {
				max = c
			}
		}
	}
	return max
}

func (cg CellGrid) String() string {
	var s string
	maxIdx := len(cg) - 1
	for i, row := range cg {
		if i < maxIdx {
			s += fmt.Sprintln(row)
		} else {
			s += fmt.Sprint(row)
		}
	}
	return s
}
