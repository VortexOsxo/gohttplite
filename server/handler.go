package server

import (
	"gohttplite/messages"
)

type Handler struct {
	method  messages.Verb
	handler func(messages.Request, messages.Response) messages.Response
}

func (h *Handler) Handle(request messages.Request, response messages.Response) messages.Response {
	return h.handler(request, response)
}

func CreateHandler(verb messages.Verb, handler func(messages.Request, messages.Response) messages.Response) Handler {
	return Handler{method: verb, handler: handler}
}

var default_handler = CreateHandler(messages.Verb(""), func(request messages.Request, response messages.Response) messages.Response {
	return messages.Response{StatusCode: messages.NOT_FOUND, Body: "Not Found"}
})
