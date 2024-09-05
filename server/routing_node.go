package server

import (
	"errors"
	"gohttplite/messages"
)

func CreateTreeNode(route string) *RoutingNode {
	return &RoutingNode{route: route, middleware: make([]Middleware, 0), handler: nil, childrens: CreateNodeContainer()}
}

func CreateTreeLeaf(handler *Handler) *RoutingNode {
	return &RoutingNode{route: "", middleware: make([]Middleware, 0), handler: handler, childrens: CreateNodeContainer()}
}

type RoutingNode struct {
	route      string
	middleware []Middleware
	handler    *Handler
	childrens  NodeContainer
}

func isRouteEqual(node *RoutingNode, route string) bool {
	return node.route == route
}

func isRouteHandleable(node *RoutingNode, route string) bool {
	return node.route == "*" || node.route[0] == ':' || node.route == route
}

func (node *RoutingNode) canHandleRequest(request messages.Request) bool {
	return node.handler != nil && node.handler.method == request.Method
}

func (node *RoutingNode) handleRequest(request messages.Request, route string) {
	if len(node.route) > 0 && node.route[0] == ':' {
		request.Args[node.route[1:]] = route
	}
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
			if children.canHandleRequest(request) {
				nodes = append(nodes, children)
				return nodes, nil
			}
		}
		return nil, errors.New("path not found")
	}

	nextNode := node.findNodeToHandleRoute(route)
	if nextNode == nil {
		return nil, errors.New("path not found")
	}

	nodes = append(nodes, nextNode)
	request.Path = remainingPath
	return nextNode.findHandlingPath(request, nodes)
}

func (node *RoutingNode) findNode(route string, method func(*RoutingNode, string) bool) *RoutingNode {
	for _, children := range node.childrens.getNodes() {
		if method(children, route) {
			return children
		}
	}
	return nil
}

func (node *RoutingNode) findNodeEqualToRoute(route string) *RoutingNode {
	return node.findNode(route, isRouteEqual)
}

func (node *RoutingNode) findNodeToHandleRoute(route string) *RoutingNode {
	return node.findNode(route, isRouteHandleable)
}
