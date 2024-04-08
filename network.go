package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
)

func StartServer() (net.Listener, net.Conn) {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	// var connection net.Conn
	connection, err := server.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("client connected")
	return server, connection
	// processClient(connection)
}

func SendMove(connection net.Conn, grid [][]int) error {
	serverMsg, err := json.Marshal(grid)
	if err != nil {
		return errors.New("something went wrong marshalling")
	}
	_, err = connection.Write(serverMsg)
	if err != nil {
		return errors.New("something went wrong sending the move")
	}
	return nil
}

func ReceiveMove(connection net.Conn) [][]int {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Something went wrong reading.")
	}
	var currentGrid [][]int
	err = json.Unmarshal(buffer[:mLen], &currentGrid)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// message := string(buffer[:mLen])
	return currentGrid
}

// func processClient(connection net.Conn) {
// 	defer connection.Close()
// 	for {
// 		buffer := make([]byte, 1024)
// 		mLen, err := connection.Read(buffer)
// 		if err != nil {
// 			fmt.Println("Error reading:", err.Error())
// 		}
// 		message := string(buffer[:mLen])
// 		fmt.Println("Received: ", message)
// 		if message == "close" {
// 			_, err = connection.Write([]byte("Closing! msg:" + string(buffer[:mLen])))
// 			break
// 		}
// 		_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
// 		// connection.Close()
// 	}
// }

func ConnectClient() net.Conn {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	// defer connection.Close()
	return connection

	///send some data
	// var scan string
	// for {
	// 	fmt.Scanln(&scan)
	// 	_, err = connection.Write([]byte(scan))
	// 	buffer := make([]byte, 1024)
	// 	mLen, err := connection.Read(buffer)
	// 	if err != nil {
	// 		fmt.Println("Error reading:", err.Error())
	// 	}
	// 	fmt.Println("Received: ", string(buffer[:mLen]))
	// }
}
