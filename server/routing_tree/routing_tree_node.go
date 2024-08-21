package routing_tree

import (
	"gohttplite/messages"
)

type RoutingTreeNode struct {
	handler  *func(messages.Request) messages.Response
	children map[string]*RoutingTreeNode
}
