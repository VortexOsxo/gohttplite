package example

import (
	"gohttplite/messages"
	"gohttplite/server"
)

func CreateUsersRouter() *server.Router {
	var users_router = server.CreateRouter("api/users")

	users_router.AddHandler("", server.CreateHandler(messages.GET, func(request messages.Request, response messages.Response) messages.Response {
		return response
	}))

	return users_router
}
