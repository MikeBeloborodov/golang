package main

import (
	"fmt"
	request "golearning/sockets/types/request"
	response "golearning/sockets/types/response"
	socket "golearning/sockets/types/socket"
	utils "golearning/sockets/utils"
	"syscall"
)

const (
	serverIp = "188.114.96.3"
	port     = 80
)

func main() {
	sock := socket.NewSocket(serverIp, port).GetSock()
	defer syscall.Closesocket(sock)
	defer syscall.WSACleanup()

	request.NewRequest(sock, "GET /comments HTTP/1.1\nHost: jsonplaceholder.typicode.com\nConnection: close\n\n")
	res := response.NewResponse(sock)
	headersJson, err := res.Headers.Json()
	if err != nil {
		fmt.Println("Error while parsing headers into json:", err)
	}

	utils.SaveStringFile("apiHeadersJson.json", headersJson)
}
