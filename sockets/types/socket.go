package types

import (
	"fmt"
	"net"
	"syscall"
)

type Socket struct {
	Socket syscall.Handle
}

func NewSocket(IPaddr string, port int) Socket {
	return Socket{Socket: getBindedSocket(IPaddr, port)}
}

func (sock Socket) GetSock() syscall.Handle {
	return sock.Socket
}

func getBindedSocket(IPaddr string, port int) syscall.Handle {
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

	// creating a tcp socket
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
