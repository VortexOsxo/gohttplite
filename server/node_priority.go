package server

type NodePriority int

const (
	DefaultNodePriority  NodePriority = 3
	VariableNodePriority NodePriority = 2
	WildcardNodePriority NodePriority = 1
	EmptyNodePriority    NodePriority = 0
)
