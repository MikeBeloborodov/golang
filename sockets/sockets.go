package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"
)

const (
	serverIp = "188.114.96.3"
	port     = 80
)

func main() {
	socket := GetBindedSocket(serverIp, port)
	defer syscall.Closesocket(socket)
	defer syscall.WSACleanup()

	request := "GET /comments HTTP/1.1\nHost: jsonplaceholder.typicode.com\nConnection: close\n\n"
	SendHttpRequest(socket, request)

	headersArr := GetHttpHeaders(socket)
	for key, value := range headersArr {
		fmt.Println("Key:", key, "|", "Value:", value)
	}

	// responseData := GetHttpResponse(socket)
	// SaveStringFile("fakeApiData.txt", responseData)
}

func GetBindedSocket(IPaddr string, port int) syscall.Handle {
	// Version requirement 8 | 8 | 8 | 8 bytes (32)
	// first byte is main version number
	// last byte is minor version number
	// version is 2.2
	var verreq uint32 = (2 << 24) | (0 << 16) | (0 << 8) | 2

	// Need to use this startup to use windows dll for sockets
	data := new(syscall.WSAData)
	err := syscall.WSAStartup(verreq, data)
	if err != nil {
		fmt.Println("Error WSAStartup:")
	}

	// creatign a tcp socket
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println("Error creating socket:", err)
	}

	// creating address
	sockaddr := &syscall.SockaddrInet4{Port: port}
	copy(sockaddr.Addr[:], net.ParseIP(IPaddr).To4())

	// connecting socket to the address
	err = syscall.Connect(sock, sockaddr)
	if err != nil {
		fmt.Println("Error connecting to so server:", err)
	}

	return sock
}

func SendHttpRequest(socket syscall.Handle, request string) {
	// constructing an HTTP get request to send
	msg := []byte(request)
	var bytesSent uint32
	var flags uint32 = 0

	// constructing WSA buffer
	wsaBuf := syscall.WSABuf{
		Buf: &msg[0],
		Len: uint32(len(msg)),
	}
	// sending a message
	err := syscall.WSASend(socket, &wsaBuf, 1, &bytesSent, flags, nil, nil)
	if err != nil {
		fmt.Println("Error while sending message to server:", err)
	}
}

func GetHttpResponse(socket syscall.Handle) []byte {
	var flags uint32 = 0
	responseData := make([]byte, 0)
	var bytesRecieved uint32 = 1
	var totalBytesRecieved uint32 = 0

	for bytesRecieved > 0 {
		buf := make([]byte, 1024)
		wsaBuf := syscall.WSABuf{
			Buf: &buf[0],
			Len: uint32(len(buf)),
		}
		err := syscall.WSARecv(socket, &wsaBuf, 1, &bytesRecieved, &flags, nil, nil)
		if err != nil {
			fmt.Println("Error while recieving data:", err)
		}
		responseData = append(responseData, buf[:bytesRecieved]...)
		totalBytesRecieved += bytesRecieved
	}
	fmt.Printf("Total bytes: %d", totalBytesRecieved)

	return responseData
}

func GetHttpHeaders(socket syscall.Handle) map[string]string {
	var headers = make(map[string]string)
	var bytesRecieved uint32 = 1
	var flags uint32 = 0
	var isLineFinished bool = false
	var isComplete bool = false
	var isFirstLine = true

	for !isComplete {
		var singleLine = make([]byte, 0)
		for !isLineFinished {
			buf := make([]byte, 1)
			wsaBuf := syscall.WSABuf{
				Buf: &buf[0],
				Len: uint32(len(buf)),
			}
			err := syscall.WSARecv(socket, &wsaBuf, 1, &bytesRecieved, &flags, nil, nil)
			if err != nil {
				fmt.Println("Error while recieving data:", err)
			}

			if buf[len(buf)-1] == 10 && len(singleLine) == 1 {
				isComplete = true
				break
			}

			if buf[len(buf)-1] == 10 && isFirstLine {
				splitString := strings.Split(string(singleLine), " ")
				headers["Protocol"] = splitString[0]
				headers["Status Code"] = splitString[1]
				headers["Status"] = splitString[2]
				isFirstLine = false
				continue
			}

			if buf[len(buf)-1] == 10 {
				splitString := strings.Split(string(singleLine), ": ")
				key, value := splitString[0], splitString[1]
				headers[key] = value
				isLineFinished = true
			}

			singleLine = append(singleLine, buf...)
		}
		isLineFinished = false
	}

	return headers
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
		os.Exit(1)
	}
}
