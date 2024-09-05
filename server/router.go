package server

import (
	"gohttplite/messages"
)

type Router struct {
	root *RoutingNode
}

func CreateRouter(route string) *Router {
	return &Router{
		root: CreateTreeNode(simplifyPath(route)),
	}
}

func (rt *Router) AddRouter(router *Router) {
	rt.ensureRootExists()
	rt.root.addNode("", router.root)
}

func (rt *Router) AddMiddleware(path string, middleware Middleware) {
	rt.ensureRootExists()
}

func (rt *Router) AddHandler(path string, handler *Handler) {
	rt.ensureRootExists()
	rt.root.addNode(path, CreateTreeLeaf(handler))
}

func (rt *Router) findHandler(request messages.Request) *Handler {
	return rt.root.findHandler(request)
}

func (rt *Router) handleRequest(request messages.Request) messages.Response {
	nodesPath, err := rt.root.findHandlingPath(request, []*RoutingNode{rt.root})

	if err != nil {
		return default_handler.handler(request, messages.Response{})
	}

	decomposedPath := decomposePath(request.Path, true)

	if len(nodesPath) == 0 {
		return default_handler.handler(request, messages.Response{})
	}

	for index, value := range nodesPath {
		value.handleRequest(request, decomposedPath[index])
	}

	handler := nodesPath[len(nodesPath)-1].handler
	return handler.Handle(request, messages.Response{})
}

func (rt *Router) ensureRootExists() {
	if rt.root == nil {
		rt.root = CreateTreeNode("")
	}
}
