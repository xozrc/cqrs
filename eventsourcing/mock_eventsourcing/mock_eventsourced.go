// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/xozrc/cqrs/eventsourcing (interfaces: EventSourced)

package mock_eventsourcing

import (
	gomock "github.com/golang/mock/gomock"
	eventsourcing "github.com/xozrc/cqrs/eventsourcing"
	types "github.com/xozrc/cqrs/types"
)

// Mock of EventSourced interface
type MockEventSourced struct {
	ctrl     *gomock.Controller
	recorder *_MockEventSourcedRecorder
}

// Recorder for MockEventSourced (not exported)
type _MockEventSourcedRecorder struct {
	mock *MockEventSourced
}

func NewMockEventSourced(ctrl *gomock.Controller) *MockEventSourced {
	mock := &MockEventSourced{ctrl: ctrl}
	mock.recorder = &_MockEventSourcedRecorder{mock}
	return mock
}

func (_m *MockEventSourced) EXPECT() *_MockEventSourcedRecorder {
	return _m.recorder
}

func (_m *MockEventSourced) ApplyEvent(_param0 eventsourcing.VersionedEvent) error {
	ret := _m.ctrl.Call(_m, "ApplyEvent", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockEventSourcedRecorder) ApplyEvent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ApplyEvent", arg0)
}

func (_m *MockEventSourced) Events() []eventsourcing.VersionedEvent {
	ret := _m.ctrl.Call(_m, "Events")
	ret0, _ := ret[0].([]eventsourcing.VersionedEvent)
	return ret0
}

func (_mr *_MockEventSourcedRecorder) Events() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Events")
}

func (_m *MockEventSourced) Id() types.Guid {
	ret := _m.ctrl.Call(_m, "Id")
	ret0, _ := ret[0].(types.Guid)
	return ret0
}

func (_mr *_MockEventSourcedRecorder) Id() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Id")
}

func (_m *MockEventSourced) Payload() []byte {
	ret := _m.ctrl.Call(_m, "Payload")
	ret0, _ := ret[0].([]byte)
	return ret0
}

func (_mr *_MockEventSourcedRecorder) Payload() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Payload")
}

func (_m *MockEventSourced) Version() int64 {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockEventSourcedRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}
