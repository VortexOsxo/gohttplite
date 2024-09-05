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

func (rt *Router) AddMiddleware(path string, middleware Middleware) {
	rt.ensureRootExists()
}

func (rt *Router) AddHandler(path string, handler *Handler) {
	rt.ensureRootExists()
	rt.root.addNode(path, CreateTreeLeaf(handler))
}

func (rt *Router) handleRequest(request messages.Request) messages.Response {
	nodesPath, err := rt.root.findHandlingPath(request, []*RoutingNode{rt.root})

	if err != nil {
		return server_error_handler.Handle(request, messages.Response{})
	}

	decomposedPath := decomposePath(request.Path, true)

	if len(nodesPath) == 0 {
		return not_found_handler.Handle(request, messages.Response{})
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
