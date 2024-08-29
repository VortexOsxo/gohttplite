package server

import "gohttplite/messages"

type RoutingTreeNode struct {
	route     string
	handlers  []*Handler
	childrens []*RoutingTreeNode
}

func CreateRoutingTreeNode(route string) *RoutingTreeNode {
	return &RoutingTreeNode{
		route:     route,
		handlers:  make([]*Handler, 0),
		childrens: make([]*RoutingTreeNode, 0),
	}
}

func (node *RoutingTreeNode) acceptRoute(route string, request messages.Request) bool {
	if node.route == "*" {
		return true
	}

	if node.route[0] == ':' {
		request.Args[node.route[1:]] = route
		return true
	}

	return node.route == route
}

func (node *RoutingTreeNode) addHandler(path string, handler *Handler) {
	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		node.handlers = append(node.handlers, handler)
		return
	}

	nextNode := node.findRoute(route, messages.Request{})

	if nextNode == nil {
		nextNode = CreateRoutingTreeNode(route)
		node.childrens = append(node.childrens, nextNode)
	}

	nextNode.addHandler(remainingPath, handler)
}

func (node *RoutingTreeNode) findRoute(route string, request messages.Request) *RoutingTreeNode {
	if node == nil {
		return nil
	}

	for _, children := range node.childrens {
		if children.acceptRoute(route, request) {
			return children
		}
	}

	return nil
}

func (currentNode *RoutingTreeNode) findHandler(request messages.Request) *Handler {
	if currentNode == nil {
		return nil
	}

	route, remainingPath := getRouteFromPath(request.Path)

	if route == "" {
		for _, handler := range currentNode.handlers {
			if handler.method == request.Method {
				return handler
			}
		}

		return nil
	}

	nextNode := currentNode.findRoute(route, request)

	request.Path = remainingPath
	return nextNode.findHandler(request)
}
