package server

import (
	"gohttplite/messages"
)

type HandlerFunc func(*messages.Request, *messages.Response) *messages.Response

func CreateHandler(verb messages.Verb, handler HandlerFunc) *Handler {
	return &Handler{method: verb, handler: handler}
}

type Handler struct {
	method  messages.Verb
	handler HandlerFunc
}

func (h *Handler) Handle(request *messages.Request, response *messages.Response) *messages.Response {
	return h.handler(request, response)
}

var not_found_handler = CreateHandler(messages.Verb(""), func(request *messages.Request, response *messages.Response) *messages.Response {
	return &messages.Response{StatusCode: messages.NOT_FOUND, Body: "Not Found"}
})

var server_error_handler = CreateHandler(messages.Verb(""), func(request *messages.Request, response *messages.Response) *messages.Response {
	return &messages.Response{StatusCode: messages.INTERNAL_SERVER_ERROR, Body: "Internal Server Error"}
})
