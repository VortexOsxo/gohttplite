package server

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

func (node *RoutingTreeNode) AcceptRoute(route string) bool {
	if node.route == "*" {
		return true
	}
	return node.route == route
}
