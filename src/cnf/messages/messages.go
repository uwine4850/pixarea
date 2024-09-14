package messages

import (
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type SingleErrorResponse struct {
	rest.ImplementDTOMessage
	Error    string
	Redirect string
}

type AuthLoginMessageRequest struct {
	rest.ImplementDTOMessage
	Username string
	Password string
}

type CSRFTokenResponse struct {
	rest.ImplementDTOMessage
	Token string
	Error string
}

var AllowedMessages = []rest.AllowMessage{
	{
		Package: "messages",
		Name:    "SingleErrorResponse",
	},
	{
		Package: "messages",
		Name:    "AuthLoginMessageRequest",
	},
	{
		Package: "messages",
		Name:    "CSRFTokenResponse",
	},
}

var MessagesList = map[string]*[]irest.IMessage{
	"frontend/src/messages/error.ts": {
		SingleErrorResponse{},
	},
	"frontend/src/messages/csrf.ts": {
		CSRFTokenResponse{},
	},
	"frontend/src/messages/auth.ts": {
		AuthLoginMessageRequest{},
	},
}
