package server

import (
	"gohttplite/messages"
)

type Router struct {
	root *RoutingNode
}

func (rt *Router) AddHandler(path string, handler *Handler) {
	rt.ensureRootExists()
	rt.root.addNode(path, CreateTreeLeaf(handler))
}

func (rt *Router) AddRouter(path string, router *Router) {
	rt.ensureRootExists()
	rt.root.addNode(path, router.root)
}

func (rt *Router) FindHandler(request messages.Request) *Handler {
	return rt.root.findHandler(request)
}

func (rt *Router) ensureRootExists() {
	if rt.root == nil {
		rt.root = CreateTreeNode("")
	}
}
