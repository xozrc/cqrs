package messaging

import (
	"encoding/json"
)

type Sender interface {
	Send([]byte) error
}

type Receiver interface {
	Start(h Handler) error
	Stop() error
}

type Handler interface {
	Handle(msg []byte) error
}

type Message struct {
	Id            string `json:"id"`
	MessageType   string `json:"message_type"`
	Payload       []byte `json:"payload"`
	CorrelationId string `json:"correlation_id"`
}

func (m *Message) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(m)
	return
}

func (m *Message) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, m)
	return
}
