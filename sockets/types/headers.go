package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"syscall"
)

type Headers struct {
	headersMap map[string]string
}

func NewHeaders(socket syscall.Handle) Headers {
	return Headers{headersMap: getHttpHeaders(socket)}
}

func (headers Headers) GetValueByKey(key string) (string, error) {
	value, ok := headers.headersMap[key]
	if ok {
		return value, nil
	} else {
		return "None", errors.New("no key found")
	}
}

func (headers Headers) Json() ([]byte, error) {
	return json.Marshal(headers.headersMap)
}

func getHttpHeaders(socket syscall.Handle) map[string]string {
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
