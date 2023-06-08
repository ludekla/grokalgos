package main

import (
	"fmt"
	"math/rand"
)



func LongCommSubsequence(leftword, rightword string) int {
	nrows := len(leftword) + 1
	ncols := len(rightword) + 1
	cellgrid := NewCellGrid(nrows, ncols)
	nrows--
	ncols--
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r + 1][c + 1] = cellgrid[r][c] + 1
			} else {
				if cellgrid[r + 1][c] > cellgrid[r][c + 1] {
					cellgrid[r + 1][c + 1] = cellgrid[r + 1][c]
				} else {
					cellgrid[r + 1][c + 1] = cellgrid[r][c + 1]
				}
			}	
		} 
	}
	// fmt.Println(cellgrid)
	return cellgrid.Max()
}

func LongCommSubstring(leftword, rightword string) int {
	nrows := len(leftword) + 1
	ncols := len(rightword) + 1
	cellgrid := NewCellGrid(nrows, ncols)
	nrows--
	ncols--
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r + 1][c + 1] = cellgrid[r][c] + 1
			}	
		} 
	}
	// fmt.Println(cellgrid)
	return cellgrid.Max()
}

func main() {

	
	

	fmt.Println("10. Dynamic Programming")
	r := LongCommSubsequence("clues", "blue")
	fmt.Printf("Longest common subsequence of 'clues' and 'blue': %d\n", r)
	r = LongCommSubstring("fish", "fosh")
	fmt.Printf("Longest common substring of 'fish' and 'fosh':    %d\n", r)
	r = LongCommSubsequence("fish", "fosh")
	fmt.Printf("Longest common subsequence of 'fish' and 'fosh':  %d\n", r)

}
