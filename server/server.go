package server

import (
	"fmt"
	"gohttplite/messages"
	"log"
	"net"
)

type Server struct {
	address         string
	router          *Router
	default_handler Handler
}

func CreateServer(address string) *Server {
	server := &Server{address: address}

	server.router = &Router{}

	server.default_handler = CreateHandler(messages.Verb(""), func(request messages.Request, response messages.Response) messages.Response {
		return messages.Response{StatusCode: messages.NOT_FOUND, Body: "Not Found"}
	})

	return server
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

func (server *Server) AddHandler(path string, method messages.Verb, handler_func func(messages.Request, messages.Response) messages.Response) {
	handler := CreateHandler(method, handler_func)
	server.router.AddHandler(path, &handler)
}

func (server *Server) AddRouter(path string, router *Router) {
	server.router.AddRouter(path, router)
}

func (server *Server) findHandler(request messages.Request) *Handler {
	handler := server.router.FindHandler(request)
	if handler == nil {
		return &server.default_handler
	}
	return handler
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	request := messages.GetRequest(conn)

	fmt.Println("Path ffs:")
	fmt.Println(request.Path)

	handler := server.findHandler(request)

	response := (*handler).Handle(request, messages.Response{})
	writeResponse(conn, response)
}

func writeResponse(conn net.Conn, response messages.Response) {
	_, err := conn.Write([]byte(response.ToString()))
	if err != nil {
		log.Println("Error writing:", err)
	}
}
