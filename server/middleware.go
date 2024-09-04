package server

import "gohttplite/messages"

type Middleware func(request messages.Request, response messages.Response, next Middleware) messages.Response
