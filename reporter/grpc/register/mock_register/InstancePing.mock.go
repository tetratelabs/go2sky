// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SkyAPM/go2sky/reporter/grpc/register (interfaces: ServiceInstancePingClient)

// Package mock_register is a generated GoMock package.
package mock_register

import (
	context "context"
	common "github.com/SkyAPM/go2sky/reporter/grpc/common"
	register "github.com/SkyAPM/go2sky/reporter/grpc/register"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServiceInstancePingClient is a mock of ServiceInstancePingClient interface
type MockServiceInstancePingClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstancePingClientMockRecorder
}

// MockServiceInstancePingClientMockRecorder is the mock recorder for MockServiceInstancePingClient
type MockServiceInstancePingClientMockRecorder struct {
	mock *MockServiceInstancePingClient
}

// NewMockServiceInstancePingClient creates a new mock instance
func NewMockServiceInstancePingClient(ctrl *gomock.Controller) *MockServiceInstancePingClient {
	mock := &MockServiceInstancePingClient{ctrl: ctrl}
	mock.recorder = &MockServiceInstancePingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceInstancePingClient) EXPECT() *MockServiceInstancePingClientMockRecorder {
	return m.recorder
}

// DoPing mocks base method
func (m *MockServiceInstancePingClient) DoPing(arg0 context.Context, arg1 *register.ServiceInstancePingPkg, arg2 ...grpc.CallOption) (*common.Commands, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoPing", varargs...)
	ret0, _ := ret[0].(*common.Commands)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoPing indicates an expected call of DoPing
func (mr *MockServiceInstancePingClientMockRecorder) DoPing(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoPing", reflect.TypeOf((*MockServiceInstancePingClient)(nil).DoPing), varargs...)
}
