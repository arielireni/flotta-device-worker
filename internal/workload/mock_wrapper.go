// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-device-worker/internal/workload (interfaces: WorkloadWrapper)

// Package workload is a generated GoMock package.
package workload

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/project-flotta/flotta-device-worker/internal/workload/api"
	v1 "k8s.io/api/core/v1"
)

// MockWorkloadWrapper is a mock of WorkloadWrapper interface.
type MockWorkloadWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadWrapperMockRecorder
}

// MockWorkloadWrapperMockRecorder is the mock recorder for MockWorkloadWrapper.
type MockWorkloadWrapperMockRecorder struct {
	mock *MockWorkloadWrapper
}

// NewMockWorkloadWrapper creates a new mock instance.
func NewMockWorkloadWrapper(ctrl *gomock.Controller) *MockWorkloadWrapper {
	mock := &MockWorkloadWrapper{ctrl: ctrl}
	mock.recorder = &MockWorkloadWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadWrapper) EXPECT() *MockWorkloadWrapperMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockWorkloadWrapper) CreateSecret(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockWorkloadWrapperMockRecorder) CreateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockWorkloadWrapper)(nil).CreateSecret), arg0, arg1)
}

// Init mocks base method.
func (m *MockWorkloadWrapper) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockWorkloadWrapperMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockWorkloadWrapper)(nil).Init))
}

// List mocks base method.
func (m *MockWorkloadWrapper) List() ([]api.WorkloadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]api.WorkloadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkloadWrapperMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkloadWrapper)(nil).List))
}

// ListSecrets mocks base method.
func (m *MockWorkloadWrapper) ListSecrets() (map[string]struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets")
	ret0, _ := ret[0].(map[string]struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockWorkloadWrapperMockRecorder) ListSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockWorkloadWrapper)(nil).ListSecrets))
}

// Logs mocks base method.
func (m *MockWorkloadWrapper) Logs(arg0 string, arg1 io.Writer) (context.CancelFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", arg0, arg1)
	ret0, _ := ret[0].(context.CancelFunc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockWorkloadWrapperMockRecorder) Logs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockWorkloadWrapper)(nil).Logs), arg0, arg1)
}

// PersistConfiguration mocks base method.
func (m *MockWorkloadWrapper) PersistConfiguration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistConfiguration")
	ret0, _ := ret[0].(error)
	return ret0
}

// PersistConfiguration indicates an expected call of PersistConfiguration.
func (mr *MockWorkloadWrapperMockRecorder) PersistConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistConfiguration", reflect.TypeOf((*MockWorkloadWrapper)(nil).PersistConfiguration))
}

// RegisterObserver mocks base method.
func (m *MockWorkloadWrapper) RegisterObserver(arg0 Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterObserver", arg0)
}

// RegisterObserver indicates an expected call of RegisterObserver.
func (mr *MockWorkloadWrapperMockRecorder) RegisterObserver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterObserver", reflect.TypeOf((*MockWorkloadWrapper)(nil).RegisterObserver), arg0)
}

// Remove mocks base method.
func (m *MockWorkloadWrapper) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockWorkloadWrapperMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockWorkloadWrapper)(nil).Remove), arg0)
}

// RemoveMappingFile mocks base method.
func (m *MockWorkloadWrapper) RemoveMappingFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMappingFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMappingFile indicates an expected call of RemoveMappingFile.
func (mr *MockWorkloadWrapperMockRecorder) RemoveMappingFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMappingFile", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveMappingFile))
}

// RemoveSecret mocks base method.
func (m *MockWorkloadWrapper) RemoveSecret(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecret", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret.
func (mr *MockWorkloadWrapperMockRecorder) RemoveSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecret", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveSecret), arg0)
}

// RemoveServicesFile mocks base method.
func (m *MockWorkloadWrapper) RemoveServicesFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServicesFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveServicesFile indicates an expected call of RemoveServicesFile.
func (mr *MockWorkloadWrapperMockRecorder) RemoveServicesFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServicesFile", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveServicesFile))
}

// RemoveTable mocks base method.
func (m *MockWorkloadWrapper) RemoveTable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTable")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTable indicates an expected call of RemoveTable.
func (mr *MockWorkloadWrapperMockRecorder) RemoveTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTable", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveTable))
}

// Run mocks base method.
func (m *MockWorkloadWrapper) Run(arg0 *v1.Pod, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWorkloadWrapperMockRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWorkloadWrapper)(nil).Run), arg0, arg1, arg2)
}

// Start mocks base method.
func (m *MockWorkloadWrapper) Start(arg0 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockWorkloadWrapperMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWorkloadWrapper)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockWorkloadWrapper) Stop(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockWorkloadWrapperMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWorkloadWrapper)(nil).Stop), arg0)
}

// UpdateSecret mocks base method.
func (m *MockWorkloadWrapper) UpdateSecret(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockWorkloadWrapperMockRecorder) UpdateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockWorkloadWrapper)(nil).UpdateSecret), arg0, arg1)
}
