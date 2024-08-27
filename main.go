package main

import (
	"gohttplite/messages"
	s "gohttplite/server"
)

func main() {
	server := s.CreateServer("localhost:8080")

	server.AddHandler("/lol", s.CreateHandler(messages.GET, func(request messages.Request) messages.Response {
		return messages.Response{
			Body: "HTTP/1.1 200 OK\r\n" + "Content-Type: text/plain\r\n" + "\r\n" + "Hello, World! from lol",
		}
	}))

	server.Start()

}
