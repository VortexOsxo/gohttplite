package main

import (
	e "gohttplite/example"
	s "gohttplite/server"
)

func main() {
	server := s.CreateServer("localhost:8080")

	server.AddRouter(e.CreateUsersRouter())

	server.Start()
}
