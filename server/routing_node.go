package server

import (
	"errors"
	"gohttplite/messages"
)

type RoutingNode struct {
	route      string
	middleware []Middleware
	handler    *Handler
	childrens  NodeContainer
}

func CreateTreeNode(route string) *RoutingNode {
	return &RoutingNode{
		route:      route,
		middleware: make([]Middleware, 0),
		handler:    nil,
		childrens:  CreateNodeContainer(),
	}
}

func CreateTreeLeaf(handler *Handler) *RoutingNode {
	return &RoutingNode{
		route:      "",
		middleware: make([]Middleware, 0),
		handler:    handler,
		childrens:  CreateNodeContainer(),
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

func (node *RoutingNode) isRouteEqual(route string) bool {
	return node.route == route
}

func (node *RoutingNode) addNode(path string, newNode *RoutingNode) {
	nodesPath := node.createNodePath(path, []*RoutingNode{node})

	nodesPath[len(nodesPath)-1].childrens.addNode(newNode)
}

func (node *RoutingNode) createNodePath(path string, nodes []*RoutingNode) []*RoutingNode {
	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		return nodes
	}

	nextNode := node.findNodeEqualToRoute(route)

	if nextNode == nil {
		nextNode = CreateTreeNode(route)
		node.childrens.addNode(nextNode)
	}

	nodes = append(nodes, nextNode)

	return nextNode.createNodePath(remainingPath, nodes)
}

func (node *RoutingNode) findHandlingPath(request messages.Request, nodes []*RoutingNode) ([]*RoutingNode, error) {
	route, remainingPath := getRouteFromPath(request.Path)

	if route == "" {
		for _, children := range node.childrens.getNodes() {
			if children.acceptMethod(request.Method) {
				nodes = append(nodes, children)
				return nodes, nil
			}
		}

		return nil, errors.New("path not found")
	}

	nextNode := node.findNodeToHandleRoute(route, request)

	if nextNode == nil {
		return nil, errors.New("path not found")
	}

	nodes = append(nodes, nextNode)

	request.Path = remainingPath
	return nextNode.findHandlingPath(request, nodes)
}

func (node *RoutingNode) findNodeEqualToRoute(route string) *RoutingNode {
	for _, children := range node.childrens.getNodes() {
		if children.isRouteEqual(route) {
			return children
		}
	}

	return nil
}

func (node *RoutingNode) findNodeToHandleRoute(route string, request messages.Request) *RoutingNode {
	for _, children := range node.childrens.getNodes() {
		if children.acceptRoute(route, request) {
			return children
		}
	}

	return nil
}

func (node *RoutingNode) findHandler(request messages.Request) *Handler {
	nodesPath, err := node.findHandlingPath(request, []*RoutingNode{node})
	if err != nil {
		return nil
	}

	return nodesPath[len(nodesPath)-1].handler
}
