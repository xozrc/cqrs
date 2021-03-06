#! /bin/bash -e
mockgen github.com/xozrc/cqrs/messaging Sender \
 > messaging/mock_messaging/mock_messaging.go
mockgen github.com/xozrc/cqrs/messaging Receiver \
 > messaging/mock_messaging/mock_messaging.go
gofmt -w messaging/mock_messaging/mock_messaging.go


mockgen github.com/xozrc/cqrs/eventsourcing VersionedEvent \
 > eventsourcing/mock_eventsourcing/mock_event.go
 mockgen github.com/xozrc/cqrs/eventsourcing EventSourced \
 > eventsourcing/mock_eventsourcing/mock_eventsourced.go
mockgen github.com/xozrc/cqrs/eventsourcing EventStore \
  > eventsourcing/mock_eventsourcing/mock_store.go
gofmt -w eventsourcing/mock_eventsourcing/mock_event.go eventsourcing/mock_eventsourcing/mock_store.go
echo >&2 "OK"
