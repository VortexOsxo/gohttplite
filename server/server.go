package server

import (
	"fmt"
	"gohttplite/messages"
	"log"
	"net"
)

type Server struct {
	address  string
	handlers []Handler
}

func CreateServer(address string) Server {
	return Server{address: address}
}

func (server *Server) AddHandler(handler Handler) {
	server.handlers = append(server.handlers, handler)
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", server.address)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

	defer listener.Close()

	fmt.Println("Server listening on", server.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	request := messages.GetRequest(conn)

	fmt.Println("Path ffs:")
	fmt.Println(request.Path)

	for _, handler := range server.handlers {
		if handler.path == request.Path {
			response := handler.handler(request)
			writeResponse(conn, response)
			return
		}
	}
}

func writeResponse(conn net.Conn, response messages.Response) {
	_, err := conn.Write([]byte(response.Body))
	if err != nil {
		log.Println("Error writing:", err)
	}
}
