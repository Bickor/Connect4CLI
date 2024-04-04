package main

import (
	"fmt"
	"net"
	"os"
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
	// fmt.Println(grid)
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}

	// Some testing for the grid checkWin
	fmt.Println()
	grid[0][0] = 1
	grid[2][2] = 1
	grid[3][3] = 1
	grid[4][4] = 1
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
	fmt.Println(checkWin(grid))

	// Game loop

	// var first string
	// fmt.Scanln(&first)
	// if first == "1" {
	// 	startServer()
	// } else if first == "2" {
	// 	connectClient()
	// }
}

func startServer() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}

func connectClient() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	///send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
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

func checkWin(grid [][]uint8) bool { // Pass in the grid
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

func drop(grid [][]uint8) { // pass in the grid and which column the thing was dropped

}

// Print the current board state
func printBoard(grid [][]uint8) {

}
