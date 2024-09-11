package server

import (
	"gohttplite/messages"
)

func CreateRouter(route string) *Router {
	return &Router{
		root: CreateTreeNode(simplifyPath(route)),
	}
}

type Router struct {
	root *RoutingNode
}

func (rt *Router) AddRouter(router *Router) {
	rt.ensureRootExists()
	rt.root.addNode("", router.root)
}

func (rt *Router) AddMiddleware(path string, middleware *Middleware) {
	rt.ensureRootExists()
	rt.root.addMiddleware(path, middleware)
}

func (rt *Router) AddHandler(path string, handler *Handler) {
	rt.ensureRootExists()
	rt.root.addNode(path, CreateTreeLeaf(handler))
}

func (rt *Router) handleRequest(request messages.Request) messages.Response {
	nodesPath, err := rt.root.findHandlingPath(request, []*RoutingNode{rt.root})
	handler := nodesPath[len(nodesPath)-1].handler

	if err != nil {
		return server_error_handler.Handle(request, messages.Response{})
	}

	decomposedPath := decomposePath(request.Path, true)

	if len(nodesPath) == 0 {
		return not_found_handler.Handle(request, messages.Response{})
	}

	emptyMiddleware := CreateMiddleWare(func(request messages.Request, response messages.Response, next *Middleware) messages.Response {
		return next.Evaluate(request, response)
	})

	currentMiddleware := emptyMiddleware

	for index, value := range nodesPath {
		currentMiddleware = value.createHandlingChain(currentMiddleware, request, decomposedPath[index])
	}

	currentMiddleware.next = CreateMiddleWare(func(request messages.Request, response messages.Response, next *Middleware) messages.Response {
		return handler.Handle(request, response)
	})

	return emptyMiddleware.Evaluate(request, messages.Response{})
}

func (rt *Router) ensureRootExists() {
	if rt.root == nil {
		rt.root = CreateTreeNode("")
	}
}
