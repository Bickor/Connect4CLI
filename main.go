package main

import (
	"fmt"
)

const (
	height = 6
	width  = 7
)

func main() {

	// Make the grid
	for i := 0; i < 6; i++ {
		for i := 0; i < 7; i++ {
			fmt.Printf("| ")
		}
		fmt.Printf("|\n")
	}

	// Initialize the grid
	grid := make([][]uint8, height)
	for i := range grid {
		grid[i] = make([]uint8, width)
	}
	fmt.Println(grid)
}

// Grid is 6 high by 7 wide. Alternating colors.
