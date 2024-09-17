package types

import (
	"encoding/json"
	"fmt"
	"golearning/sockets/types/apitypes"
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

func (body ResponseBody) Text() string {
	return string(body.RawData)
}

func (body ResponseBody) Json() ([]byte, error) {
	var indexOfStart int
	var indexOfEnd int
	text := body.Text()

	for index, value := range text {
		if string(value) == "[" || string(value) == "{" {
			indexOfStart = index
			break
		}
	}
	for i := len(text) - 1; i >= 0; i-- {
		if string(text[i]) == "]" || string(text[i]) == "}" {
			indexOfEnd = i
			break
		}
	}

	usersArr := make([]apitypes.User, 0)
	err := json.Unmarshal(body.RawData[indexOfStart:indexOfEnd+1], &usersArr)
	if err != nil {
		fmt.Println("Error while unmarshal:", err)
	}
	return json.Marshal(usersArr)
}
