package main

import (
	"fmt"
	"strconv"
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
	// for i := 0; i < 6; i++ {
	// 	for i := 0; i < 7; i++ {
	// 		fmt.Printf("| ")
	// 	}
	// 	fmt.Printf("|\n")
	// }

	// Initialize the grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Remove into its own file
	// test(grid)

	// Game loop
	var input string
	player1Turn := true
	currentPlayer := 1
	for {
		printGrid(grid)
		fmt.Println("Chose a column from 1-7")
		// TODO: Should check if new input was put or just enter was pressed.
		fmt.Scanln(&input)
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Cannot convert to int! Err: ", err)
		}

		if inputInt > 7 || inputInt < 1 {
			fmt.Println("Can't accept that number!")
			continue
		}

		inputInt -= 1
		if player1Turn {
			Drop(&grid, inputInt, currentPlayer)
		} else {
			Drop(&grid, inputInt, currentPlayer)
		}
		won := CheckWin(grid, currentPlayer)
		if won {
			printGrid(grid)
			fmt.Printf("Player %d won!!\n", currentPlayer)
			break
		}

		if currentPlayer == 1 {
			currentPlayer = 2
		} else {
			currentPlayer = 1
		}
	}
	fmt.Println("Finished")

	// var first string
	// fmt.Scanln(&first)
	// if first == "1" {
	// 	startServer()
	// } else if first == "2" {
	// 	connectClient()
	// }
}

func printGrid(grid [][]int) {
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

func test(grid [][]int) {
	// Some testing for the Drop.
	printGrid(grid)
	err := Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	printGrid(grid)
	err = Drop(&grid, 3, 1)
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
