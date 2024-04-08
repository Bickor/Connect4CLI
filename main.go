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

	// Initialize the grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	//TODO: add msg value verification
	var first string
	var move string
	var msg [][]int
	var gameFinished bool
	fmt.Scanln(&first)
	if first == "1" {
		server, connection := StartServer()
		defer server.Close()
		msg = grid
		for {
			fmt.Println("Make your move, select a column from 1-7")
			ShowBoard(msg)
			validMove := false
			for !validMove {
				fmt.Scanln(&move)
				move, err := strconv.Atoi(move)
				if err != nil {
					fmt.Println("Cannot convert to int! Err: ", err)
				}
				err = Drop(&msg, move-1, 1)
				if err != nil {
					fmt.Println("Not valid column, choose another!")
				} else {
					validMove = true
				}
			}

			SendMove(connection, msg)
			gameFinished = CheckWin(msg, 1)
			if gameFinished {
				ShowBoard(msg)
				fmt.Println("You won!")
				break
			}

			msg = ReceiveMove(connection)
			gameFinished = CheckWin(msg, 2)
			if gameFinished {
				ShowBoard(msg)
				fmt.Println("You lost!")
				break
			}
		}
	} else if first == "2" {
		connection := ConnectClient()
		defer connection.Close()
		for {
			msg = ReceiveMove(connection)
			gameFinished = CheckWin(msg, 1)
			if gameFinished {
				ShowBoard(msg)
				fmt.Println("You lost!")
				break
			}
			fmt.Println("Make your move, select a column from 1-7")
			ShowBoard(msg)
			validMove := false
			for !validMove {
				fmt.Scanln(&move)
				move, err := strconv.Atoi(move)
				if err != nil {
					fmt.Println("Cannot convert to int! Err: ", err)
				}
				err = Drop(&msg, move-1, 2)
				if err != nil {
					fmt.Println("Not valid column, choose another!")
				} else {
					validMove = true
				}
			}

			SendMove(connection, msg)
			gameFinished = CheckWin(msg, 2)
			if gameFinished {
				ShowBoard(msg)
				fmt.Println("You won!")
				break
			}
		}
	}
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
	ShowBoard(grid)
	// Some testing for the grid checkWin
	// fmt.Println()
	// grid[0][0] = 1
	// grid[2][2] = 1
	// grid[3][3] = 1
	// grid[4][4] = 1
	// printGrid(grid)
}
