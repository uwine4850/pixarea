package messages

import (
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type SingleErrorResponse struct {
	rest.InmplementDTOMessage
	Error string
}

type AuthLoginMessageRequest struct {
	rest.InmplementDTOMessage
	Username string
	Password string
}

type CSRFTokenResponse struct {
	rest.InmplementDTOMessage
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
	"frontend/src/messages/messages.ts": {
		SingleErrorResponse{},
		AuthLoginMessageRequest{},
	},
	"frontend/src/messages/csrf.ts": {
		CSRFTokenResponse{},
	},
}
