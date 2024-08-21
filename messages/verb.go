package messages

type Verb string

const (
	GET    Verb = "GET"
	POST   Verb = "POST"
	PUT    Verb = "PUT"
	DELETE Verb = "DELETE"
	PATCH  Verb = "PATCH"
	ANY    Verb = "ANY"
)
