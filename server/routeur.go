package server

import "gohttplite/messages"

type Routeur struct {
	root *RoutingTreeNode
}

func (routeur *Routeur) AddHandler(path string, method messages.Verb, handler_func func(messages.Request, messages.Response) messages.Response) {
	handler := CreateHandler(method, handler_func)
	routeur.root.addHandler(path, &handler)
}
