package messages

type Response struct {
	Body string
}

func CreateResponse() Response {
	return Response{
		Body: "HTTP/1.1 200 OK\r\n" + "Content-Type: text/plain\r\n" + "\r\n" + "Hello, World!",
	}
}
