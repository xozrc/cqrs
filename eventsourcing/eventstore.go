package eventsourcing

//event entity
type EventEntity struct {
	Id           int64 `gorm:"primary_key"`
	PartitionKey string
	SourceType   string
	EventType    string

	SourceId      string
	CorrelationId string
	Version       int64
	Payload       string
}

func (ed *EventEntity) TableName() string {
	return "t_event"
}

type EventStore interface {
	Load(partitionKey string, version int64) ([]*EventEntity, error)
	Save(partitionKey string, events []*EventEntity) error
}
