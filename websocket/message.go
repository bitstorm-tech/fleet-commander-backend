package websocket

import "gitlab.com/fleet-commander/fleet-commander-backend-go/game"

type MessageType string

const (
	Error      MessageType = "error"
	SignUp     MessageType = "sign_up"
	SignIn     MessageType = "sign_in"
	SignOut    MessageType = "sign_out"
	Correction MessageType = "correction"
	GameRules  MessageType = "game_rules"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewErrorMessage(message string) Message {
	return Message{
		Type:    Error,
		Payload: message,
	}
}

func NewSignInMessage() Message {
	return Message{
		Type: SignIn,
	}
}

func NewSignUpMessage() Message {
	return Message{
		Type: SignUp,
	}
}

func NewCorrectionMessage(r game.Resources) Message {
	return Message{
		Type:    Correction,
		Payload: r,
	}
}

func NewGameRulesMessage() Message {
	return Message{
		Type:    GameRules,
		Payload: ActiveRules,
	}
}
