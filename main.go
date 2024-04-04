package main

import (
	"fmt"
)

const (
	height = 6
	width  = 7

	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
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

	// Remove into its own file
	test(grid)

	// Game loop

	// var first string
	// fmt.Scanln(&first)
	// if first == "1" {
	// 	startServer()
	// } else if first == "2" {
	// 	connectClient()
	// }
}

func printGrid(grid [][]uint8) {
	// fmt.Println(grid)
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
	fmt.Println()
}

// Grid is 6 high by 7 wide. Alternating colors.

// Flags
// -p = port
// -h = host
// -d = difficulty (for playing solo)

// Commmands:
// Host
// Play Solo
// Conenct
func parseCommands() {

}

func test(grid [][]uint8) {
	// Some testing for the Drop.
	printGrid(grid)
	err := Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)

	// Some testing for the grid checkWin
	// fmt.Println()
	// grid[0][0] = 1
	// grid[2][2] = 1
	// grid[3][3] = 1
	// grid[4][4] = 1
	// printGrid(grid)
}
