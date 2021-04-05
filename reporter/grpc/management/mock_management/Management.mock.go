// Code generated by MockGen. DO NOT EDIT.
// Source: skywalking.apache.org/repo/goapi/collect/management/v3 (interfaces: ManagementServiceClient)

// Package mock_v3 is a generated GoMock package.
package mock_v3

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
	v30 "skywalking.apache.org/repo/goapi/collect/management/v3"
)

// MockManagementServiceClient is a mock of ManagementServiceClient interface.
type MockManagementServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagementServiceClientMockRecorder
}

// MockManagementServiceClientMockRecorder is the mock recorder for MockManagementServiceClient.
type MockManagementServiceClientMockRecorder struct {
	mock *MockManagementServiceClient
}

// NewMockManagementServiceClient creates a new mock instance.
func NewMockManagementServiceClient(ctrl *gomock.Controller) *MockManagementServiceClient {
	mock := &MockManagementServiceClient{ctrl: ctrl}
	mock.recorder = &MockManagementServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagementServiceClient) EXPECT() *MockManagementServiceClientMockRecorder {
	return m.recorder
}

// KeepAlive mocks base method.
func (m *MockManagementServiceClient) KeepAlive(arg0 context.Context, arg1 *v30.InstancePingPkg, arg2 ...grpc.CallOption) (*v3.Commands, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "KeepAlive", varargs...)
	ret0, _ := ret[0].(*v3.Commands)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockManagementServiceClientMockRecorder) KeepAlive(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockManagementServiceClient)(nil).KeepAlive), varargs...)
}

// ReportInstanceProperties mocks base method.
func (m *MockManagementServiceClient) ReportInstanceProperties(arg0 context.Context, arg1 *v30.InstanceProperties, arg2 ...grpc.CallOption) (*v3.Commands, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportInstanceProperties", varargs...)
	ret0, _ := ret[0].(*v3.Commands)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportInstanceProperties indicates an expected call of ReportInstanceProperties.
func (mr *MockManagementServiceClientMockRecorder) ReportInstanceProperties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportInstanceProperties", reflect.TypeOf((*MockManagementServiceClient)(nil).ReportInstanceProperties), varargs...)
}
