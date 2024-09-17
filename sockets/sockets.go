package main

import (
	"fmt"
	"golearning/sockets/types"
	"golearning/sockets/types/common"
	"golearning/sockets/utils"
)

const (
	serverIp = "188.114.96.3"
	port     = 80
)

func main() {
	sock := types.NewSocket(serverIp, port).GetSock()
	types.NewRequest(sock, "GET /users HTTP/1.1\nHost: jsonplaceholder.typicode.com\nConnection: close\n\n")
	res := common.NewResponse(sock)

	headersJson, err := res.Headers.Json()
	if err != nil {
		fmt.Println("Error while parsing headers into json:", err)
		err = nil
	}
	utils.SaveStringFile("apiHeaders.json", headersJson)

	bodyJson, err := res.Body.Json()
	if err != nil {
		fmt.Println("Error while parsing body into json:", err)
		err = nil
	}
	utils.SaveStringFile("apiBody.json", bodyJson)
}
