package messaging

import (
	"encoding/json"
)

type Sender interface {
	Send([]byte) error
}

type Receiver interface {
	Start(h MessageHandler) error
	Stop() error
}

type MessageHandler interface {
	Handle(msg []byte) error
}

type Message struct {
	Id            string          `json:"id"`
	MessageType   string          `json:"message_type"`
	Payload       json.RawMessage `json:"payload"`
	CorrelationId string          `json:"correlationId"`
}
