// Code generated by MockGen. DO NOT EDIT.
// Source: calculator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCalculator is a mock of Calculator interface.
type MockCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorMockRecorder
}

// MockCalculatorMockRecorder is the mock recorder for MockCalculator.
type MockCalculatorMockRecorder struct {
	mock *MockCalculator
}

// NewMockCalculator creates a new mock instance.
func NewMockCalculator(ctrl *gomock.Controller) *MockCalculator {
	mock := &MockCalculator{ctrl: ctrl}
	mock.recorder = &MockCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculator) EXPECT() *MockCalculatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculator) Add(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorMockRecorder) Add(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculator)(nil).Add), a, b)
}

// Mul mocks base method.
func (m *MockCalculator) Mul(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mul", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Mul indicates an expected call of Mul.
func (mr *MockCalculatorMockRecorder) Mul(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mul", reflect.TypeOf((*MockCalculator)(nil).Mul), a, b)
}
