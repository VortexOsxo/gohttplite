package server

import (
	"fmt"
	"gohttplite/messages"
	"log"
	"net"
)

type Server struct {
	address string
	router  *Router
}

func CreateServer(address string) *Server {
	server := &Server{address: address}

	server.router = &Router{}
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

func (server *Server) AddRouter(router *Router) {
	server.router.AddRouter(router)
}

func (server *Server) handleRequest(request messages.Request) messages.Response {
	return server.router.handleRequest(request)
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	request := messages.GetRequest(conn)
	response := server.handleRequest(request)

	writeResponse(conn, response)
}

func writeResponse(conn net.Conn, response messages.Response) {
	_, err := conn.Write([]byte(response.ToString()))
	if err != nil {
		log.Println("Error writing:", err)
	}
}
