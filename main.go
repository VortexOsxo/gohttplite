package main

import (
	"gohttplite/messages"
	s "gohttplite/server"
)

func main() {
	server := s.CreateServer("localhost:8080")

	server.AddHandler("/lol", messages.GET, func(request messages.Request, response messages.Response) messages.Response {
		response.SetStatusCode(messages.OK)
		response.Body = "Hello, World! from lol"
		return response
	})

	server.Start()
}
