package server

import "gohttplite/messages"

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

	nextNode := node.findNodeByRoute(route)

	if nextNode == nil {
		nextNode = CreateTreeNode(route)
		node.childrens.addNode(nextNode)
	}

	nodes = append(nodes, nextNode)

	return nextNode.createNodePath(remainingPath, nodes)
}

func (node *RoutingNode) findNodeByRoute(route string) *RoutingNode {
	if node == nil {
		return nil
	}

	for _, children := range node.childrens.getNodes() {
		if children.isRouteEqual(route) {
			return children
		}
	}

	return nil
}

// TODO: Divide finding a node and accepting a request
func (node *RoutingNode) findNodeByRouteOld(route string, request messages.Request) *RoutingNode {
	if node == nil {
		return nil
	}

	for _, children := range node.childrens.getNodes() {
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
		for _, children := range node.childrens.getNodes() {
			if children.acceptMethod(request.Method) {
				return children.handler
			}
		}

		return nil
	}

	nextNode := node.findNodeByRouteOld(route, request)

	request.Path = remainingPath
	return nextNode.findHandler(request)
}
