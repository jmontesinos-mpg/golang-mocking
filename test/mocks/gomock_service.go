// Code generated by MockGen. DO NOT EDIT.
// Source: jmontesinos/golang-mocking/test (interfaces: APIServiceA)
//
// Generated by this command:
//
//	mockgen -destination gomock_service.go -package mocks jmontesinos/golang-mocking/test APIServiceA
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAPIServiceA is a mock of APIServiceA interface.
type MockAPIServiceA struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServiceAMockRecorder
}

// MockAPIServiceAMockRecorder is the mock recorder for MockAPIServiceA.
type MockAPIServiceAMockRecorder struct {
	mock *MockAPIServiceA
}

// NewMockAPIServiceA creates a new mock instance.
func NewMockAPIServiceA(ctrl *gomock.Controller) *MockAPIServiceA {
	mock := &MockAPIServiceA{ctrl: ctrl}
	mock.recorder = &MockAPIServiceAMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIServiceA) EXPECT() *MockAPIServiceAMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAPIServiceA) Add(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockAPIServiceAMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAPIServiceA)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockAPIServiceA) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAPIServiceAMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAPIServiceA)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockAPIServiceA) GetAll() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAPIServiceAMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAPIServiceA)(nil).GetAll))
}
