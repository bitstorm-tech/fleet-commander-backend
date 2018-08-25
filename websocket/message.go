package websocket

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

func NewErrorMessage(text string) *Message {
	return &Message{
		Type: ErrorType,
		Payload: struct {
			Message string `json:"message"`
		}{
			Message: text,
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

func NewCorrectionMessage(titanium int, fuel int, energy int) *Message {
	return &Message{
		Type: CorrectionType,
		Payload: struct {
			Titanium int
			Fuel     int
			Energy   int
		}{
			Titanium: titanium,
			Fuel:     fuel,
			Energy:   energy,
		},
	}
}
