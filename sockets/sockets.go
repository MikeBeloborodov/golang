package main

import (
	"fmt"
	"golearning/sockets/types"
	"golearning/sockets/types/common"
	"golearning/sockets/utils"
	"syscall"
)

const (
	serverIp = "188.114.96.3"
	port     = 80
)

func main() {
	sock := types.NewSocket(serverIp, port).GetSock()
	defer syscall.Closesocket(sock)
	defer syscall.WSACleanup()

	types.NewRequest(sock, "GET /comments HTTP/1.1\nHost: jsonplaceholder.typicode.com\nConnection: close\n\n")
	res := common.NewResponse(sock)
	headersJson, err := res.Headers.Json()
	if err != nil {
		fmt.Println("Error while parsing headers into json:", err)
	}

	utils.SaveStringFile("apiHeadersJson.json", headersJson)
}
