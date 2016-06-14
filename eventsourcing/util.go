package eventsourcing

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xozrc/cqrs/types"
)

func GetPartitionKey1(sourceType string, id types.Guid) string {
	return fmt.Sprintf("%s_%d", sourceType, id)
}

func GetPartitionKey(es EventSourced) string {
	st := reflect.TypeOf(es).Name()
	return fmt.Sprintf("%s_%d", st, es.Id())
}

func snapShotEventSourced(es *EventSourced) (bs []byte, err error) {
	return
}

func ToData(st string, partitionKey string, e VersionedEvent) (*EventEntity, error) {
	ed := &EventEntity{}
	ed.PartitionKey = partitionKey
	ed.SourceType = st
	ed.EventType = reflect.TypeOf(e).Name()
	//json endcode event
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	ed.Payload = string(payload)
	ed.SourceId = fmt.Sprintf("%d", e.SourceId())
	ed.Version = e.Version()
	return ed, nil
}
