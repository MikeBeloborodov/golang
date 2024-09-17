package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverMessage := GetFileData("test.json")
	listener, err := net.Listen("tcp", "127.0.0.1:3030")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 3030")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, serverMessage)
	}
}

func handleConnection(conn net.Conn, serverMessage string) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	_, err = conn.Write([]byte(serverMessage))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
	}

	fmt.Printf("Recieved: %s", string(buffer[:n]))
}

func GetFileData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(string(data))
	return string(data)
}
