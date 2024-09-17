package types

import (
	"fmt"
	"syscall"
)

type ResponseBody struct {
	RawData []byte
}

func NewResponseBody(socket syscall.Handle) ResponseBody {
	return ResponseBody{RawData: getRawBody(socket)}
}

func getRawBody(socket syscall.Handle) []byte {
	var bodyData = make([]byte, 0)
	var bytesRecieved uint32 = 1
	var flags uint32 = 0
	var bufLen uint32 = 1024

	for {
		buf := make([]byte, bufLen)
		wsaBuf := syscall.WSABuf{
			Buf: &buf[0],
			Len: bufLen,
		}
		err := syscall.WSARecv(socket, &wsaBuf, 1, &bytesRecieved, &flags, nil, nil)
		if err != nil {
			fmt.Println("Error while recieving data:", err)
			break
		}
		if bytesRecieved == 0 {
			break
		}

		bodyData = append(bodyData, buf[:bytesRecieved]...)
	}

	return bodyData
}
