package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type ServerMessage struct {
	Message string
}

func main() {
	addr := "localhost:3030"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Error connecting to %s: %s", addr, err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "GET / HTTP/1.1\nHost: icanhazip.com\n\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error while writing a message:", err)
		return
	}
	fmt.Println("Message sent, message:", message)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error while reading a message:", err)
		return
	}
	fmt.Println("Received from server:")
	fmt.Println(string(buffer[:n]))

	var serverMess ServerMessage
	jsonErr := json.Unmarshal(buffer[:n], &serverMess)
	if jsonErr != nil {
		fmt.Println("Error parsing json:", jsonErr)
	}

	fmt.Printf("Server message: %s", serverMess.Message)
}
