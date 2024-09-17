package common

import (
	"golearning/sockets/types"
	"syscall"
)

type Response struct {
	Headers types.Headers
	Body    types.ResponseBody
}

func NewResponse(socket syscall.Handle) Response {
	defer syscall.Closesocket(socket)
	defer syscall.WSACleanup()
	return Response{
		Headers: types.NewHeaders(socket),
		Body:    types.NewResponseBody(socket),
	}
}
