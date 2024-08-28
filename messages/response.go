package messages

import "fmt"

type Response struct {
	StatusCode StatusCode
	Body       string
}

func (response *Response) SetStatusCode(statusCode StatusCode) {
	response.StatusCode = statusCode
}

func (response *Response) ToString() string {
	return fmt.Sprintf("HTTP/1.1 %d %s\r\n"+"Content-Type: text/plain\r\n"+"\r\n"+response.Body, response.StatusCode, response.StatusCode.ToString())
}
