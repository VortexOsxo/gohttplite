package server

import (
	"testing"
)

func TestRoutingNodeCreateNodesPathSimple(t *testing.T) {
	node := CreateTreeNode("")

	path := node.createNodePath("api/users/profile/like", []*RoutingNode{node})

	if len(path) != 5 {
		t.Errorf("The created path does not have the right number of node")
	}

	if firstNode := path[0]; firstNode != node {
		t.Errorf("The created path replaced the root of the tree")
	}

	if thirdNode := path[2]; thirdNode.route != "users" {
		t.Errorf("The node of the created path does not have the right route")
	}

	if lastNode := path[4]; lastNode.route != "like" {
		t.Errorf("The node of the created path does not have the right route")
	}
}

func TestRoutingNodeCreateNodesPathExistingPath(t *testing.T) {
	node := CreateTreeNode("")

	path := node.createNodePath("api/users/profile/like", []*RoutingNode{node})

	existingThirdNode := path[2]

	secondPath := node.createNodePath("api/users/subscribers", []*RoutingNode{node})

	if thirdNode := secondPath[2]; thirdNode != existingThirdNode {
		t.Errorf("The creation of the path did not use the previously created path")
	}
}
