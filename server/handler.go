package server

import (
	"gohttplite/messages"
)

type Handler struct {
	method  messages.Verb
	handler func(messages.Request) messages.Response
}

func (h *Handler) Handle(request messages.Request) messages.Response {
	return h.handler(request)
}

func CreateHandler(verb messages.Verb, handler func(messages.Request) messages.Response) Handler {
	return Handler{method: verb, handler: handler}
}
