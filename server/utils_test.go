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
