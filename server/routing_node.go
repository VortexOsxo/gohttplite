package server

import "gohttplite/messages"

type RoutingNode struct {
	route     string
	handler   *Handler
	childrens []*RoutingNode
}

func CreateTreeNode(route string) *RoutingNode {
	return &RoutingNode{
		route:     route,
		handler:   nil,
		childrens: make([]*RoutingNode, 0),
	}
}

func CreateTreeLeaf(handler *Handler) *RoutingNode {
	return &RoutingNode{
		route:     "",
		handler:   handler,
		childrens: make([]*RoutingNode, 0),
	}
}

func (node *RoutingNode) acceptRoute(route string, request messages.Request) bool {
	if node.route == "" {
		return false
	} else if node.route == "*" {
		return true
	} else if node.route[0] == ':' {
		request.Args[node.route[1:]] = route
		return true
	}

	return node.route == route
}

func (node *RoutingNode) acceptMethod(method messages.Verb) bool {
	return node.handler != nil && node.handler.method == method
}

func (node *RoutingNode) addNode(path string, newNode *RoutingNode) {
	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		node.childrens = append(node.childrens, newNode)
		return
	}

	nextNode := node.findNodeByRoute(route, messages.Request{})

	if nextNode == nil {
		nextNode = CreateTreeNode(route)
		node.childrens = append(node.childrens, nextNode)
	}

	nextNode.addNode(remainingPath, newNode)
}

// TODO: Divide finding a node and accepting a request
func (node *RoutingNode) findNodeByRoute(route string, request messages.Request) *RoutingNode {
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

func (node *RoutingNode) findHandler(request messages.Request) *Handler {
	if node == nil {
		return nil
	}

	route, remainingPath := getRouteFromPath(request.Path)

	if route == "" {
		for _, children := range node.childrens {
			if children.acceptMethod(request.Method) {
				return children.handler
			}
		}

		return nil
	}

	nextNode := node.findNodeByRoute(route, request)

	request.Path = remainingPath
	return nextNode.findHandler(request)
}
