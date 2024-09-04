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

func (rt *Router) FindHandler(request messages.Request) *Handler {
	return rt.root.findHandler(request)
}

func (rt *Router) ensureRootExists() {
	if rt.root == nil {
		rt.root = CreateTreeNode("")
	}
}
