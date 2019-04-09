// Code generated by MockGen. DO NOT EDIT.
// Source: number-server/app/domain/service (interfaces: NumberService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "number-server/app/domain/model"
	reflect "reflect"
)

// MockNumberService is a mock of NumberService interface
type MockNumberService struct {
	ctrl     *gomock.Controller
	recorder *MockNumberServiceMockRecorder
}

// MockNumberServiceMockRecorder is the mock recorder for MockNumberService
type MockNumberServiceMockRecorder struct {
	mock *MockNumberService
}

// NewMockNumberService creates a new mock instance
func NewMockNumberService(ctrl *gomock.Controller) *MockNumberService {
	mock := &MockNumberService{ctrl: ctrl}
	mock.recorder = &MockNumberServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNumberService) EXPECT() *MockNumberServiceMockRecorder {
	return m.recorder
}

// GetCounters mocks base method
func (m *MockNumberService) GetCounters() *model.Report {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounters")
	ret0, _ := ret[0].(*model.Report)
	return ret0
}

// GetCounters indicates an expected call of GetCounters
func (mr *MockNumberServiceMockRecorder) GetCounters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounters", reflect.TypeOf((*MockNumberService)(nil).GetCounters))
}

// IsValidNumber mocks base method
func (m *MockNumberService) IsValidNumber(arg0 *model.Number) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidNumber", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidNumber indicates an expected call of IsValidNumber
func (mr *MockNumberServiceMockRecorder) IsValidNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidNumber", reflect.TypeOf((*MockNumberService)(nil).IsValidNumber), arg0)
}

// ResetCounters mocks base method
func (m *MockNumberService) ResetCounters() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetCounters")
}

// ResetCounters indicates an expected call of ResetCounters
func (mr *MockNumberServiceMockRecorder) ResetCounters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCounters", reflect.TypeOf((*MockNumberService)(nil).ResetCounters))
}

// Store mocks base method
func (m *MockNumberService) Store(arg0 *model.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockNumberServiceMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockNumberService)(nil).Store), arg0)
}
