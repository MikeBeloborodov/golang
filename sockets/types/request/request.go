package request

import (
	"fmt"
	"syscall"
)

type Request struct {
}

func NewRequest(socket syscall.Handle, request string) Request {
	sendHttpRequest(socket, request)
	return Request{}
}

func sendHttpRequest(socket syscall.Handle, request string) {
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
