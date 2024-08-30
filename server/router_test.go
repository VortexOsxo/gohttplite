package server

import (
	"gohttplite/messages"
	"testing"
)

func CreateEmptyRequest(method messages.Verb, path string) messages.Request {
	return messages.Request{Method: method, Path: path, Args: make(map[string]string)}
}

func TestRoutingTreeDefault(t *testing.T) {
	rt := &Router{}

	handler1 := &Handler{method: messages.GET}
	handler2 := &Handler{method: messages.POST}
	handler3 := &Handler{method: messages.PUT}

	rt.AddHandler("/users", handler1)
	rt.AddHandler("/users/profile", handler2)
	rt.AddHandler("/users/profile/update/picture", handler3)

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/users")); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.POST, "/users/profile")); found != handler2 {
		t.Errorf("Expected to find handler2 for POST /users/profile")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.PUT, "/users/profile/update/picture")); found != handler3 {
		t.Errorf("Expected to find handler3 for PUT /users/profile/update/picture")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.PUT, "/users/profile/update")); found != nil {
		t.Errorf("Expected nil for PUT /users/profile/update")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/nonexistent")); found != nil {
		t.Errorf("Expected nil for nonexistent path")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.POST, "/users")); found != nil {
		t.Errorf("Expected nil for existing path but wrong method")
	}
}

func TestRoutingTreeAny(t *testing.T) {
	rt := &Router{}

	handler1 := &Handler{method: messages.GET}
	handler2 := &Handler{method: messages.GET}
	handler3 := &Handler{method: messages.GET}

	rt.AddHandler("/users/*", handler1)
	rt.AddHandler("/users/profile", handler2)
	rt.AddHandler("/*/delete", handler3)

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/users/eat")); found != handler1 {
		t.Errorf("Expected to find handler3 for GET /users/eat")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/drugs/delete")); found != handler3 {
		t.Errorf("Expected to find handler3 for GET /drugs/delete")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/users/profile")); found != handler1 {
		t.Errorf("Expected to find handler2 for GET /users/profile")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/users/delete")); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users/delete")
	}

	if found := rt.FindHandler(CreateEmptyRequest(messages.GET, "/rainbow/delete")); found != handler3 {
		t.Errorf("Expected to find handler3 for GET /rainbow/delete")
	}
}

func TestRoutingTreeArgumentSimple(t *testing.T) {
	rt := &Router{}

	handler1 := &Handler{method: messages.GET}

	rt.AddHandler("/users/:id", handler1)

	request1 := CreateEmptyRequest(messages.GET, "/users/123")
	if found := rt.FindHandler(request1); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users/123")
	}

	if request1.Args["id"] != "123" {
		t.Errorf("Expected path param id to be 123, got %s", request1.Args["id"])
	}
}

func TestRoutingTreeArgumentMultiple(t *testing.T) {
	rt := &Router{}

	handler1 := &Handler{method: messages.GET}

	rt.AddHandler("/users/:id/profile/:name/change", handler1)

	request1 := CreateEmptyRequest(messages.GET, "/users/123/profile/John/change")
	if found := rt.FindHandler(request1); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users/123/profile/John/change")
	}

	if request1.Args["id"] != "123" {
		t.Errorf("Expected path param id to be 123, got %s", request1.Args["id"])
	}

	if request1.Args["name"] != "John" {
		t.Errorf("Expected path param name to be John, got %s", request1.Args["name"])
	}

	if len(request1.Args) != 2 {
		t.Errorf("Expected 2 path params, got %d", len(request1.Args))
	}

}

func TestRoutingRouter(t *testing.T) {
	server := &Router{}

	rt := &Router{}

	handler1 := &Handler{method: messages.GET}
	handler2 := &Handler{method: messages.GET}

	rt.AddHandler("/users", handler1)
	rt.AddHandler("/users/:id", handler2)

	server.AddRouter("/api", rt)

	request1 := CreateEmptyRequest(messages.GET, "/api/users")
	if found := server.FindHandler(request1); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users")
	}

	request2 := CreateEmptyRequest(messages.GET, "/api/users/123")
	if found := server.FindHandler(request2); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users/123")
	}

	if request1.Args["id"] != "123" {
		t.Errorf("Expected path param id to be 123, got %s", request1.Args["id"])
	}
}
