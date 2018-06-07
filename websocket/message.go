package websocket

const (
	ErrorType = iota
	SignUpType
	SignInType
	SignOutType
)

type Message struct {
	Type    int         `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type ErrorPayload struct {
	Text string `json:"text"`
}

func NewErrorMessage(text string) *Message {
	return &Message{
		Type: ErrorType,
		Payload: ErrorPayload{
			Text: text,
		},
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
