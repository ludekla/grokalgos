package ch05

import "fmt"

func LongCommSubsequence(leftword, rightword string, verbose bool) int {
	nrows := len(leftword)
	ncols := len(rightword)
	cellgrid := NewCellGrid(nrows+1, ncols+1)
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r+1][c+1] = cellgrid[r][c] + 1
			} else {
				if cellgrid[r+1][c] > cellgrid[r][c+1] {
					cellgrid[r+1][c+1] = cellgrid[r+1][c]
				} else {
					cellgrid[r+1][c+1] = cellgrid[r][c+1]
				}
			}
		}
	}
	if verbose {
		fmt.Println(cellgrid)
	}
	return cellgrid.Max()
}

func LongCommSubstring(leftword, rightword string, verbose bool) int {
	nrows := len(leftword)
	ncols := len(rightword)
	cellgrid := NewCellGrid(nrows+1, ncols+1)
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r+1][c+1] = cellgrid[r][c] + 1
			}
		}
	}
	if verbose {
		fmt.Println(cellgrid)
	}
	return cellgrid.Max()
}
