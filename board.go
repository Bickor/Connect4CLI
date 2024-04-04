package main

import (
	"errors"
	"fmt"
)

func CheckWin(grid [][]uint8) bool { // Pass in the grid
	// Check all up-down
	for y := 0; y < 3; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 1 && grid[y+1][x] == 1 && grid[y+2][x] == 1 && grid[y+3][x] == 1 {
				return true
			}
		}
	}

	// Check all left-right
	for y := 0; y < height; y++ {
		for x := 0; x < 4; x++ {
			if grid[y][x] == 1 && grid[y][x+1] == 1 && grid[y][x+2] == 1 && grid[y][x+3] == 1 {
				return true
			}
		}
	}

	// Check all diagonals
	// Check left to right diagonals
	for y := 0; y < 3; y++ {
		for x := 0; x < 4; x++ {
			if grid[y][x] == 1 && grid[y+1][x+1] == 1 && grid[y+2][x+2] == 1 && grid[y+3][x+3] == 1 {
				return true
			}
		}
	}

	// Check right to left diagonals
	for y := 0; y < 3; y++ {
		for x := width - 1; x > 2; x-- {
			if grid[y][x] == 1 && grid[y+1][x-1] == 1 && grid[y+2][x-2] == 1 && grid[y+3][x-3] == 1 {
				return true
			}
		}
	}

	return false
}

func Drop(grid *[][]uint8, column int) error { // pass in the grid and which column the thing was dropped
	if column > width {
		return errors.New("can't select that column")
	}
	for i := height - 1; i >= 0; i-- {
		if (*grid)[i][column] == 0 {
			(*grid)[i][column] = 1
			return nil
		}
	}

	fmt.Println("ERROR!")
	return errors.New("no space in that position")
}
