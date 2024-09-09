package server

import "testing"

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

func TestDecomposePath(t *testing.T) {
	path := "/users/robert/profile"

	decomposedPath := decomposePath(path, false)

	if len(decomposedPath) != 3 {
		t.Errorf("Decompose path returns a path of wrong size")
	}

	if decomposedPath[0] != "users" || decomposedPath[1] != "robert" || decomposedPath[2] != "profile" {
		t.Errorf("Decompose path did not return the proper value in the path")
	}
}
