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

type ResourcesPayload struct {
	Titanium int `json:"titanium"`
	Fuel     int `json:"fuel"`
	Energy   int `json:"energy"`
}

type MessagePayload struct {
	Message string `json:"message"`
}

func NewErrorMessage(message string) *Message {
	return &Message{
		Type: ErrorType,
		Payload: MessagePayload{
			Message: message,
		},
	}
}

func NewSignInMessage(titanium int, fuel int, energy int) *Message {
	return &Message{
		Type: SignInType,
		Payload: ResourcesPayload{
			Titanium: titanium,
			Fuel:     fuel,
			Energy:   energy,
		},
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
		Payload: ResourcesPayload{
			Titanium: titanium,
			Fuel:     fuel,
			Energy:   energy,
		},
	}
}
