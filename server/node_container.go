package server

import (
	"sort"
)

func CreateNodeContainer() NodeContainer {
	return NodeContainer{make([]*RoutingNode, 0), false}
}

type NodeContainer struct {
	nodes   []*RoutingNode
	isDirty bool
}

func (container *NodeContainer) addNode(node *RoutingNode) {
	container.nodes = append(container.nodes, node)
	container.isDirty = true
}

func (container *NodeContainer) getNodes() []*RoutingNode {
	if container.isDirty {
		container.sortNodes()
		container.isDirty = false
	}
	return container.nodes
}

func (container *NodeContainer) sortNodes() {
	sort.Slice(container.nodes, func(i, j int) bool {
		return getNodePriority(container.nodes[i]) > getNodePriority(container.nodes[j])
	})
}

func getNodePriority(node *RoutingNode) NodePriority {
	if node.route == "" {
		return EmptyNodePriority
	} else if node.route == "*" {
		return WildcardNodePriority
	} else if node.route[0] == ':' {
		return VariableNodePriority
	}
	return DefaultNodePriority
}
