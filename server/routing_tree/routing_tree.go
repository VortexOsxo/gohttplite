package routing_tree

import (
	"gohttplite/messages"
	"strings"
)

type RoutingTree struct {
	root *RoutingTreeNode
}

func (rt *RoutingTree) AddRoute(path string, handler *func(messages.Request) messages.Response) {
	if rt.root == nil {
		rt.root = &RoutingTreeNode{
			handler:  nil,
			children: make(map[string]*RoutingTreeNode),
		}
	}

	rt.addRoute(rt.root, path, handler)
}

func (rt *RoutingTree) addRoute(currentNode *RoutingTreeNode, path string, handler *func(messages.Request) messages.Response) {
	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		currentNode.handler = handler
		return
	} else {
		if currentNode.children[route] == nil {
			currentNode.children[route] = &RoutingTreeNode{
				handler:  nil,
				children: make(map[string]*RoutingTreeNode),
			}
		}
		rt.addRoute(currentNode.children[route], remainingPath, handler)
	}
}

func (rt *RoutingTree) FindRoute(path string) *func(messages.Request) messages.Response {
	return rt.findRoute(rt.root, path)
}

func (rt RoutingTree) findRoute(currentNode *RoutingTreeNode, path string) *func(messages.Request) messages.Response {
	if currentNode == nil {
		return nil
	}

	route, remainingPath := getRouteFromPath(path)

	if route == "" {
		return currentNode.handler
	}

	return rt.findRoute(currentNode.children[route], remainingPath)
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
