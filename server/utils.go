package server

import "strings"

func getRouteFromPath(path string) (string, string) {
	if len(path) > 1 && path[0] == '/' {
		path = path[1:]
	}

	if len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	index := strings.Index(path, "/")

	if index == -1 {
		return path, ""
	}

	route := path[:index]
	remainingPath := path[index:]

	return route, remainingPath
}
