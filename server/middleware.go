package server

import "gohttplite/messages"

type MiddlewareFunc func(*messages.Request, *messages.Response, *Middleware) *messages.Response

func CreateMiddleWare(middleware MiddlewareFunc) *Middleware {
	return &Middleware{middlewareFunc: &middleware}
}

type Middleware struct {
	middlewareFunc *MiddlewareFunc
	next           *Middleware
}

func (middleware *Middleware) Evaluate(request *messages.Request, response *messages.Response) *messages.Response {
	return (*middleware.middlewareFunc)(request, response, middleware.next)
}
