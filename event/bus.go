package event

type EventBus interface {
	Publish(e Event) error
	PublishEvents(e []Event) error
}

type EventPublisher interface {
	Events() []Event
}
