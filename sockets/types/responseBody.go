package types

import (
	"fmt"
	"syscall"
	"time"
)

type ResponseBody struct {
	rawData []byte
}

func NewResponseBody(socket syscall.Handle) ResponseBody {
	return ResponseBody{rawData: getRawBody(socket)}
}

func getRawBody(socket syscall.Handle) []byte {
	var bodyData = make([]byte, 0)
	var bytesRecieved uint32 = 1
	var flags uint32 = 0
	var isLineFinished bool = false
	var isComplete bool = false

	for !isComplete {
		var singleLine = make([]byte, 0)
		time.Sleep(1 * time.Second)
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

			if buf[len(buf)-1] == 10 {
				fmt.Println(singleLine)
				bodyData = append(bodyData, singleLine...)
				isLineFinished = true
				continue
			}

			singleLine = append(singleLine, buf...)
		}
		isLineFinished = false
	}

	return bodyData
}
