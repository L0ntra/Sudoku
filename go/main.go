package main

import (
	"fmt"
	"math"
)

func main() {
	grid := Set()

	solved := false
	for !solved {
		grid, solved = Solve(grid)
	}

	PrintGrid(grid)
}

func Set() [][]Cell {
	/*
		    s := [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},

				{9, 1, 2, 3, 4, 5, 6, 7, 8},
				{6, 7, 8, 0, 1, 2, 3, 4, 5},
				{3, 4, 5, 6, 7, 8, 9, 1, 2},

				{5, 6, 7, 8, 9, 1, 2, 3, 4},
				{8, 9, 1, 2, 3, 4, 5, 6, 7},
				{2, 3, 4, 5, 6, 7, 8, 9, 1},
			}
	*/

	s := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 3, 0, 0, 0, 0, 1, 6, 0},
		{0, 6, 7, 0, 3, 5, 0, 0, 4},

		{6, 0, 8, 1, 2, 0, 9, 0, 0},
		{0, 9, 0, 0, 8, 0, 0, 3, 0},
		{0, 0, 2, 0, 7, 9, 8, 0, 6},

		{8, 0, 0, 6, 9, 0, 3, 5, 0},
		{0, 2, 6, 0, 0, 0, 0, 9, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	/* solvable := [][]int{
	       {0, 8, 0, 2, 0, 0, 0, 0, 0},
	       {0, 3, 0, 0, 0, 0, 1, 6, 0},
	       {0, 6, 7, 0, 3, 5, 0, 0, 4},

	       {6, 0, 8, 1, 2, 0, 9, 0, 0},
	       {0, 9, 1, 0, 8, 6, 0, 3, 0},
	       {0, 0, 2, 0, 7, 9, 8, 0, 6},

	       {8, 0, 0, 6, 9, 0, 3, 5, 0},
	       {0, 2, 6, 0, 0, 0, 0, 9, 0},
	       {0, 0, 0, 0, 0, 0, 0, 0, 0},
	   }
	*/

	grid := MakeGrid()

	for x, row := range s {
		for y, cell := range row {
			if cell != 0 {
				grid[x][y].Set(cell, false)
			}
		}
	}

	return grid
}

func Solve(grid [][]Cell) ([][]Cell, bool) {
	for x, row := range grid {
		for y, cell := range row {
			if len(cell) == 1 {
				for k, v := range cell {
					if !v {
						return SetCell(grid, x, y, k), false
					}
				}
			}
		}
	}
	return grid, true
}

func SetCell(grid [][]Cell, x, y, n int) [][]Cell {

	// Remove From Box
	boxX := int(math.Floor(float64(x)/3) * 3)
	boxY := int(math.Floor(float64(y)/3) * 3)

	for i := boxX; i < boxX+3; i++ {
		for j := boxY; j < boxY+3; j++ {
			grid[i][j].Remove(n)
		}
	}

	// Remove From Row and Col
	for i := 0; i < 9; i++ {
		grid[x][i].Remove(n)
		grid[i][y].Remove(n)
	}

	// Set the Value
	grid[x][y].Set(n, true)

	return grid
}

func MakeGrid() [][]Cell {
	grid := make([][]Cell, 0, 9)
	for i := 0; i < 9; i++ {
		row := make([]Cell, 0, 9)
		for j := 0; j < 9; j++ {
			row = append(row, NewCell(nil))
		}
		grid = append(grid, row)
	}
	return grid
}

func PrintGrid(grid [][]Cell) {
	for _, row := range grid {
		for _, cell := range row {
			cell.Print()
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}
