// socket-client project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

// https://www.developer.com/languages/intro-socket-programming-go/

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

/*
	CLIENT
*/

func example_client() {
	// establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}

	// send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))

	// get some data
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// print
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}

/*
	SERVER
*/

func example_server() {
	// start server
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	// do something with new connection
	for {
		// acceptnew connection
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// send them in another thread
		fmt.Println("client connected")
		go processClient(connection)
	}
}

// another thread
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

func main() {
	example_client()
}
