package server

import (
	"gohttplite/messages"
)

type Handler struct {
	path    string
	handler func(messages.Request) messages.Response
}

func CreateHandler(path string, handler func(messages.Request) messages.Response) Handler {
	return Handler{path: path, handler: handler}
}
