package server

import (
	"fmt"
	"gohttplite/messages"
	"log"
	"net"
)

type Server struct {
	address         string
	routing_tree    *RoutingTree
	default_handler Handler
}

func CreateServer(address string) *Server {
	server := &Server{address: address}

	server.routing_tree = &RoutingTree{}

	server.default_handler = CreateHandler(messages.Verb(""), func(request messages.Request) messages.Response {
		return messages.Response{Body: "HTTP/1.1 404 Not Found\r\n" + "Content-Type: text/plain\r\n" + "\r\n" + "Not Found"}
	})

	return server
}

func (server *Server) AddHandler(path string, handler Handler) {
	server.routing_tree.AddHandler(path, &handler)
}

func (server *Server) findHandler(method messages.Verb, path string) *Handler {
	handler := server.routing_tree.FindHandler(method, path)
	if handler == nil {
		return &server.default_handler
	}
	return handler
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

	handler := server.findHandler(request.Method, request.Path)

	response := (*handler).Handle(request)
	writeResponse(conn, response)
}

func writeResponse(conn net.Conn, response messages.Response) {
	_, err := conn.Write([]byte(response.Body))
	if err != nil {
		log.Println("Error writing:", err)
	}
}
