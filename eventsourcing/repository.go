package eventsourcing

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

var (
	EventSourcedNoFound = errors.New("eventsoucing no found")
)

type Repository interface {
	Find(id types.Guid, es EventSourced) (err error)
	Save(es EventSourced, correlationId string) error
}

type EventSourcedRepository struct {
	es EventStore //event store
	//event sender
}

func (esr *EventSourcedRepository) Find(id types.Guid, es EventSourced) (err error) {

	//todo: read from cache
	var tv int64 = 0

	st := cqrspkg.TypeName(es)

	partitionKey := GetPartitionKey(st, id)

	teds, err := esr.es.Load(partitionKey, tv)
	if err != nil {
		return
	}

	if len(teds) == 0 {
		err = EventSourcedNoFound
		return
	}

	tes := make([]VersionedEvent, 0, len(teds))
	//convert to event
	for _, ted := range teds {

		factory := GetVersionEventFactory(ted.EventType)
		if factory == nil {
			return errors.New("no found factory")
		}
		sourceId, err := strconv.ParseInt(ted.SourceId, 10, 64)
		if err != nil {
			return err
		}

		te := factory.NewVersionEvent(types.Guid(sourceId), ted.Version)
		err = json.Unmarshal([]byte(ted.Payload), te)

		if err != nil {
			//todo: do extra action
			return err
		}

		tes = append(tes, te)

	}

	//load events
	for _, e := range tes {
		err = es.ApplyEvent(e)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return
	}
	return
}

func (esr *EventSourcedRepository) Save(es EventSourced, correlationId string) error {

	st := cqrspkg.TypeName(es)
	partitionKey := GetPartitionKey(st, es.Id())

	tes := es.Events()

	eds := make([]*EventEntity, 0, len(tes))
	for _, e := range tes {
		ed, err := ToData(st, partitionKey, e)
		if err != nil {
			return err
		}
		ed.CorrelationId = correlationId
		eds = append(eds, ed)
	}
	fmt.Printf("%#v\n", eds)
	//save in store
	err := esr.es.Save(partitionKey, eds)
	if err != nil {
		//todo: do extra action
		return err
	}

	//publish async

	//cache snapshot
	return nil
}

func NewRepository(es EventStore) *EventSourcedRepository {
	esr := &EventSourcedRepository{}
	esr.es = es
	return esr
}
