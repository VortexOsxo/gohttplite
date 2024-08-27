package server

import (
	"gohttplite/messages"
	"testing"
)

func TestRoutingTreeDefault(t *testing.T) {
	rt := &RoutingTree{}

	handler1 := &Handler{method: messages.GET}
	handler2 := &Handler{method: messages.POST}
	handler3 := &Handler{method: messages.PUT}

	rt.AddHandler("/users", handler1)
	rt.AddHandler("/users/profile", handler2)
	rt.AddHandler("/users/profile/update/picture", handler3)

	if found := rt.FindHandler(messages.GET, "/users"); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users")
	}

	if found := rt.FindHandler(messages.POST, "/users/profile"); found != handler2 {
		t.Errorf("Expected to find handler2 for POST /users/profile")
	}

	if found := rt.FindHandler(messages.PUT, "/users/profile/update/picture"); found != handler3 {
		t.Errorf("Expected to find handler3 for PUT /users/profile/update/picture")
	}

	if found := rt.FindHandler(messages.PUT, "/users/profile/update"); found != nil {
		t.Errorf("Expected nil for PUT /users/profile/update")
	}

	if found := rt.FindHandler(messages.GET, "/nonexistent"); found != nil {
		t.Errorf("Expected nil for nonexistent path")
	}

	if found := rt.FindHandler(messages.POST, "/users"); found != nil {
		t.Errorf("Expected nil for existing path but wrong method")
	}
}

func TestRoutingTreeAny(t * testing.T) {
	rt := &RoutingTree{}

	handler1 := &Handler{method: messages.GET}
	handler2 := &Handler{method: messages.GET}
	handler3 := &Handler{method: messages.GET}

	rt.AddHandler("/users/*", handler1)
	rt.AddHandler("/users/profile", handler2)
	rt.AddHandler("/*/delete", handler3)

	if found := rt.FindHandler(messages.GET, "/users/eat"); found != handler1 {
		t.Errorf("Expected to find handler3 for GET /users/eat")
	}

	if found := rt.FindHandler(messages.GET, "/drugs/delete"); found != handler3 {
		t.Errorf("Expected to find handler3 for GET /drugs/delete")
	}

	if found := rt.FindHandler(messages.GET, "/users/profile"); found != handler1 {
		t.Errorf("Expected to find handler2 for GET /users/profile")
	}

	if found := rt.FindHandler(messages.GET, "/users/delete"); found != handler1 {
		t.Errorf("Expected to find handler1 for GET /users/delete")
	}

	if found := rt.FindHandler(messages.GET, "/rainbow/delete"); found != handler3 {
		t.Errorf("Expected to find handler3 for GET /rainbow/delete")
	}
}

func TestGetRouteFromPath(t *testing.T) {
	testCases := []struct {
		input          string
		expectedRoute  string
		expectedRemain string
	}{
		{"/users", "users", ""},
		{"/users/profile", "users", "/profile"},
		{"users/profile/", "users", "/profile"},
		{"/", "", ""},
		{"", "", ""},
	}

	for _, tc := range testCases {
		route, remain := getRouteFromPath(tc.input)
		if route != tc.expectedRoute || remain != tc.expectedRemain {
			t.Errorf("For input %s, expected (%s, %s) but got (%s, %s)",
				tc.input, tc.expectedRoute, tc.expectedRemain, route, remain)
		}
	}
}

