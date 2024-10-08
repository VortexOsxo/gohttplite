package messages

import (
	"net"
	"strings"
)

type Request struct {
	Method  Verb
	Path    string
	Headers map[string]string
	Args    map[string]string
	Body    string
}

func getRequestString(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)

	request, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:request]), nil
}

func GetRequest(conn net.Conn) Request {
	request, _ := getRequestString(conn)

	headers := make(map[string]string)
	var method, path, body string

	lines := strings.Split(request, "\n")

	for index, line := range lines {
		if line == "\r" || line == "" {
			break
		}

		if index == 0 {
			arguments := strings.Split(line, " ")

			if len(arguments) < 2 {
				continue
			}
			method = arguments[0]
			path = arguments[1]
		} else {
			arguments := strings.Split(line, ": ")

			if len(arguments) < 2 {
				continue
			}
			headers[arguments[0]] = arguments[1]
		}
	}

	bodyStart := strings.Index(request, "\r\n\r\n")

	if bodyStart != -1 {
		body = request[bodyStart+4:]
	}

	return Request{Method: Verb(method), Path: path, Headers: headers, Args: make(map[string]string), Body: body}
}
