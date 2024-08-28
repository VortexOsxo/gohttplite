package messages

type StatusCode int

const (
	OK                    StatusCode = 200
	NOT_FOUND             StatusCode = 404
	INTERNAL_SERVER_ERROR StatusCode = 500
)

func (statusCode StatusCode) ToString() string {
	switch statusCode {
	case OK:
		return "OK"
	case NOT_FOUND:
		return "Not Found"
	case INTERNAL_SERVER_ERROR:
		return "Internal Server Error"
	}
	return "Unknown Status Code"
}
