package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn := SendHttpRequest("jsonplaceholder.typicode.com", "/comments")
	defer conn.Close()

	data := make([]byte, 0)
	isCarrigeFlag := false
	for {
		chunk := make([]byte, 1024)
		n, err := conn.Read(chunk)
		chunk = chunk[:n]

		if err != nil {
			fmt.Println("Error while reading chunk:", err)
			return
		}
		if n == 0 {
			fmt.Println("reached end")
			break
		}

		data = append(data, chunk...)
		if contains(chunk, 13) && isCarrigeFlag {
			fmt.Println("End of file")
			break
		}
		if contains(chunk, 13) {
			isCarrigeFlag = true
		}
	}

	SaveStringFile("testDataApi.txt", data)
}

func contains(data []byte, target byte) bool {
	for _, value := range data {
		if value == target {
			return true
		}
	}
	return false
}

// func GetHeaders(conn net.Conn) map[string]string {
// 	headers := make(map[string]string)

// 	chunk := 1024

// 	return headers
// }

func SendHttpRequest(host, path string) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:http", host))
	if err != nil {
		fmt.Printf("Error connecting to %s: %s", host, err)
		os.Exit(1)
	}
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	request := fmt.Sprintf("GET %s HTTP/1.1\nHost: %s\n\n", path, host)
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error while writing a message:", err)
		os.Exit(1)
	}

	return conn
}

func SaveStringFile(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating file:", err)
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
