package messages

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	StatusCode  StatusCode
	Body        string
	contentType string
}

func (response *Response) SetStatusCode(statusCode StatusCode) {
	response.StatusCode = statusCode
}

func (response *Response) SetMessage(message string) {
	response.Body = message
	response.contentType = "text/plain"
}

func (response *Response) SetJson(object any) error {
	jsonData, err := json.Marshal(object)
	if err != nil {
		return err
	}

	response.Body = string(jsonData)
	response.contentType = "application/json"

	return nil
}

func (response *Response) ToString() string {
	return fmt.Sprintf("HTTP/1.1 %d %s\r\n"+"Content-Type: %s\r\n"+"\r\n"+response.Body, response.StatusCode, response.StatusCode.ToString(), response.contentType)
}
