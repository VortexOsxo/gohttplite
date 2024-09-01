package server

type NodePriority int

const (
	Default  NodePriority = 3
	Variable NodePriority = 2
	Wildcard NodePriority = 1
	Empty    NodePriority = 0
)
