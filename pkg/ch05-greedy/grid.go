package ch05

type CellGrid [][]int

func NewCellGrid(nrows, ncols int) CellGrid {
	cell := make(CellGrid, nrows)
	for r := 0; r < nrows; r++ {
		cell[r] = make([]int, ncols)
	}
	return cell
}

func (cg CellGrid) Max() int {
	var max int
	for _, row := range cg {
		for _, c := range row {
			if c > max {
				max = c
			}
		}
	}
	return max
}
