// Code generated by MockGen. DO NOT EDIT.
// Source: api/controller/connectors.go

// Package controller is a generated GoMock package.
package controller

import (
	pulsar "github.com/apache/pulsar-client-go/pulsar"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	protocols "rdk.io/protocols"
	reflect "reflect"
)

// MockIGrpcClientConnectionWrp is a mock of IGrpcClientConnectionWrp interface
type MockIGrpcClientConnectionWrp struct {
	ctrl     *gomock.Controller
	recorder *MockIGrpcClientConnectionWrpMockRecorder
}

// MockIGrpcClientConnectionWrpMockRecorder is the mock recorder for MockIGrpcClientConnectionWrp
type MockIGrpcClientConnectionWrpMockRecorder struct {
	mock *MockIGrpcClientConnectionWrp
}

// NewMockIGrpcClientConnectionWrp creates a new mock instance
func NewMockIGrpcClientConnectionWrp(ctrl *gomock.Controller) *MockIGrpcClientConnectionWrp {
	mock := &MockIGrpcClientConnectionWrp{ctrl: ctrl}
	mock.recorder = &MockIGrpcClientConnectionWrpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGrpcClientConnectionWrp) EXPECT() *MockIGrpcClientConnectionWrpMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIGrpcClientConnectionWrp) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIGrpcClientConnectionWrpMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIGrpcClientConnectionWrp)(nil).Close))
}

// GetConnection mocks base method
func (m *MockIGrpcClientConnectionWrp) GetConnection() *grpc.ClientConn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection")
	ret0, _ := ret[0].(*grpc.ClientConn)
	return ret0
}

// GetConnection indicates an expected call of GetConnection
func (mr *MockIGrpcClientConnectionWrpMockRecorder) GetConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockIGrpcClientConnectionWrp)(nil).GetConnection))
}

// MockIConnectors is a mock of IConnectors interface
type MockIConnectors struct {
	ctrl     *gomock.Controller
	recorder *MockIConnectorsMockRecorder
}

// MockIConnectorsMockRecorder is the mock recorder for MockIConnectors
type MockIConnectorsMockRecorder struct {
	mock *MockIConnectors
}

// NewMockIConnectors creates a new mock instance
func NewMockIConnectors(ctrl *gomock.Controller) *MockIConnectors {
	mock := &MockIConnectors{ctrl: ctrl}
	mock.recorder = &MockIConnectorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConnectors) EXPECT() *MockIConnectorsMockRecorder {
	return m.recorder
}

// getGrpcConnection mocks base method
func (m *MockIConnectors) getGrpcConnection() IGrpcClientConnectionWrp {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getGrpcConnection")
	ret0, _ := ret[0].(IGrpcClientConnectionWrp)
	return ret0
}

// getGrpcConnection indicates an expected call of getGrpcConnection
func (mr *MockIConnectorsMockRecorder) getGrpcConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getGrpcConnection", reflect.TypeOf((*MockIConnectors)(nil).getGrpcConnection))
}

// getGrpcClient mocks base method
func (m *MockIConnectors) getGrpcClient(conn *IGrpcClientConnectionWrp) protocols.RandomizerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getGrpcClient", conn)
	ret0, _ := ret[0].(protocols.RandomizerClient)
	return ret0
}

// getGrpcClient indicates an expected call of getGrpcClient
func (mr *MockIConnectorsMockRecorder) getGrpcClient(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getGrpcClient", reflect.TypeOf((*MockIConnectors)(nil).getGrpcClient), conn)
}

// getPulsarClient mocks base method
func (m *MockIConnectors) getPulsarClient() pulsar.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPulsarClient")
	ret0, _ := ret[0].(pulsar.Client)
	return ret0
}

// getPulsarClient indicates an expected call of getPulsarClient
func (mr *MockIConnectorsMockRecorder) getPulsarClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPulsarClient", reflect.TypeOf((*MockIConnectors)(nil).getPulsarClient))
}

// getPulsarProducer mocks base method
func (m *MockIConnectors) getPulsarProducer(pulsarClient *pulsar.Client) pulsar.Producer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPulsarProducer", pulsarClient)
	ret0, _ := ret[0].(pulsar.Producer)
	return ret0
}

// getPulsarProducer indicates an expected call of getPulsarProducer
func (mr *MockIConnectorsMockRecorder) getPulsarProducer(pulsarClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPulsarProducer", reflect.TypeOf((*MockIConnectors)(nil).getPulsarProducer), pulsarClient)
}
