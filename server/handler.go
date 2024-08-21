package server

import (
	"gohttplite/messages"
)

type Handler struct {
	method  messages.Verb
	handler func(messages.Request) messages.Response
}

func CreateHandler(verb messages.Verb, handler func(messages.Request) messages.Response) Handler {
	return Handler{method: verb, handler: handler}
}
