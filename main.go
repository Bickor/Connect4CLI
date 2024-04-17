package main

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	height = 6
	width  = 7
)

var (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"

	portFlag string
	ipFlag   string
	hostFlag bool
)

func main() {

	flag.StringVar(&portFlag, "port", "8080", "Choose the port where you want to host the game")
	flag.StringVar(&ipFlag, "ip", "localhost", "Choose the IP where you want to host your game")
	flag.BoolVar(&hostFlag, "host", false, "Choose the IP where you want to host your game")

	flag.Parse()

	SERVER_PORT = portFlag
	SERVER_HOST = ipFlag

	fmt.Println("Port: ", portFlag)
	fmt.Println("ip: ", ipFlag)
	fmt.Println("host: ", hostFlag)

	// Initialize the grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	var move string
	var msg [][]int
	var gameFinished bool

	if hostFlag {
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
	} else {
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
