package websocket

import "gitlab.com/fleet-commander/fleet-commander-backend-go/game"

const (
	ErrorType      = 0
	SignUpType     = 1
	SignInType     = 2
	SignOutType    = 3
	CorrectionType = 4
)

type Message struct {
	Type    int         `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewErrorMessage(message string) *Message {
	return &Message{
		Type:    ErrorType,
		Payload: message,
	}
}

func NewSignInMessage() *Message {
	return &Message{
		Type: SignInType,
	}
}

func NewSignUpMessage() *Message {
	return &Message{
		Type: SignUpType,
	}
}

func NewCorrectionMessage(r game.Resources) *Message {
	return &Message{
		Type:    CorrectionType,
		Payload: r,
	}
}
