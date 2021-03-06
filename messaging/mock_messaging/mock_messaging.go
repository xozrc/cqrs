// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/xozrc/cqrs/messaging (interfaces: Receiver)

package mock_messaging

import (
	gomock "github.com/golang/mock/gomock"
	messaging "github.com/xozrc/cqrs/messaging"
)

// Mock of Receiver interface
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *_MockReceiverRecorder
}

// Recorder for MockReceiver (not exported)
type _MockReceiverRecorder struct {
	mock *MockReceiver
}

func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &_MockReceiverRecorder{mock}
	return mock
}

func (_m *MockReceiver) EXPECT() *_MockReceiverRecorder {
	return _m.recorder
}

func (_m *MockReceiver) Start(_param0 messaging.Handler) error {
	ret := _m.ctrl.Call(_m, "Start", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReceiverRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockReceiver) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReceiverRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
