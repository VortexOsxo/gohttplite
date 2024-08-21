package main

import (
	"gohttplite/messages"
	s "gohttplite/server"
)

func main() {
	server := s.CreateServer("localhost:8080")

	server.AddHandler(s.CreateHandler("/api", func(request messages.Request) messages.Response {
		return messages.CreateResponse()
	}))
	server.Start()
}
