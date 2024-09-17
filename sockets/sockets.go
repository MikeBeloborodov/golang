package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"unsafe"
)

const (
	serverIp = "188.114.96.3"
	port     = 80
)

func main() {
	// Version requirement
	var verreq uint32 = (2 << 24) | (0 << 16) | (0 << 8) | 1
	data := new(syscall.WSAData)
	err := syscall.WSAStartup(verreq, data)
	if err != nil {
		fmt.Println("Error WSAStartup:")
	}
	defer syscall.WSACleanup()

	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println("Error creating socket:", err)
	}
	defer syscall.Closesocket(sock)

	sockaddr := &syscall.SockaddrInet4{Port: port}
	copy(sockaddr.Addr[:], net.ParseIP(serverIp).To4())

	err = syscall.Connect(sock, sockaddr)
	if err != nil {
		fmt.Println("Error connecting to so server:", err)
	}

	msg := []byte("GET /comments HTTP/1.1\nHost: jsonplaceholder.typicode.com\nConnection: close\n\n")
	var bytesSent uint32
	var flags uint32 = 0

	err = syscall.WSASend(sock, (*syscall.WSABuf)(unsafe.Pointer(&syscall.WSABuf{Len: uint32(len(msg)), Buf: &msg[0]})), 1, &bytesSent, flags, nil, nil)
	if err != nil {
		fmt.Println("Error while sending message to server:", err)
	}

	responseData := make([]byte, 0)
	var bytesRecieved uint32 = 1
	var totalBytesRecieved uint32 = 0

	for bytesRecieved > 0 {
		buf := make([]byte, 1024)
		wsaBuf := syscall.WSABuf{
			Buf: &buf[0],
			Len: uint32(len(buf)),
		}
		err = syscall.WSARecv(sock, &wsaBuf, 1, &bytesRecieved, &flags, nil, nil)
		if err != nil {
			fmt.Println("Error while recieving data:", err)
		}
		responseData = append(responseData, buf[:bytesRecieved]...)
		totalBytesRecieved += bytesRecieved
	}

	fmt.Printf("Total bytes: %d", totalBytesRecieved)
	SaveStringFile("fakeApiData.txt", responseData)
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
