package server

import (
	"gohttplite/messages"
	"strings"
)

type RoutingTree struct {
	root *RoutingTreeNode
}

func (rt *RoutingTree) AddHandler(path string, handler *Handler) {
	if rt.root == nil {
		rt.root = CreateRoutingTreeNode("")
	}

	rt.addHandler(rt.root, path, handler)
}

func (rt *RoutingTree) addHandler(currentNode *RoutingTreeNode, path string, handler *Handler) {
	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		currentNode.handlers = append(currentNode.handlers, handler)
		return
	}

	nextNode := rt.findRoute(route, currentNode)

	if nextNode == nil {
		nextNode = CreateRoutingTreeNode(route)
		currentNode.childrens = append(currentNode.childrens, nextNode)
	}

	rt.addHandler(nextNode, remainingPath, handler)
}

func (rt *RoutingTree) findRoute(route string, currentNode *RoutingTreeNode) *RoutingTreeNode {
	if currentNode == nil {
		return nil
	}

	for _, children := range currentNode.childrens {
		if children.AcceptRoute(route) {
			return children
		}
	}

	return nil
}

func (rt *RoutingTree) FindHandler(method messages.Verb, path string) *Handler {
	return rt.findHandler(rt.root, method, path)
}

func (rt RoutingTree) findHandler(currentNode *RoutingTreeNode, method messages.Verb, path string) *Handler {
	if currentNode == nil {
		return nil
	}

	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		for _, handler := range currentNode.handlers {
			if handler.method == method {
				return handler
			}
		}

		return nil
	}

	nextNode := rt.findRoute(route, currentNode)

	return rt.findHandler(nextNode, method, remainingPath)
}

func getRouteFromPath(path string) (string, string) {
	if len(path) > 1 && path[0] == '/' {
		path = path[1:]
	}

	if len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	index := strings.Index(path, "/")

	if index == -1 {
		return path, ""
	}

	route := path[:index]
	remainingPath := path[index:]

	return route, remainingPath
}
