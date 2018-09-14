package websocket

import "github.com/bugjoe/fleet-commander-backend/game"

type MessageType string

const (
	Error      MessageType = "error"
	SignUp     MessageType = "sign_up"
	SignIn     MessageType = "sign_in"
	SignOut    MessageType = "sign_out"
	GameRules  MessageType = "game_rules"
	Resources  MessageType = "resources"
	Ships      MessageType = "ships"
	MotherShip MessageType = "mother_ship"
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

func NewGameRulesMessage() Message {
	return Message{
		Type:    GameRules,
		Payload: game.ActiveRules,
	}
}

func NewResourcesMessage(r game.Resources) Message {
	return Message{
		Type:    Resources,
		Payload: r,
	}
}

func NewShipsMessage(s game.Ships) Message {
	return Message{
		Type:    Ships,
		Payload: s,
	}
}

func NewMotherShipMessage(m game.MotherShip) Message {
	return Message{
		Type:    MotherShip,
		Payload: m,
	}
}
