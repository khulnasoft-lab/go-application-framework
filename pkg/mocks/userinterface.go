// Code generated by MockGen. DO NOT EDIT.
// Source: userinterface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ui "github.com/khulnasoft-lab/go-application-framework/pkg/ui"
)

// MockUserInterface is a mock of UserInterface interface.
type MockUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserInterfaceMockRecorder
}

// MockUserInterfaceMockRecorder is the mock recorder for MockUserInterface.
type MockUserInterfaceMockRecorder struct {
	mock *MockUserInterface
}

// NewMockUserInterface creates a new mock instance.
func NewMockUserInterface(ctrl *gomock.Controller) *MockUserInterface {
	mock := &MockUserInterface{ctrl: ctrl}
	mock.recorder = &MockUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInterface) EXPECT() *MockUserInterfaceMockRecorder {
	return m.recorder
}

// Input mocks base method.
func (m *MockUserInterface) Input(prompt string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Input", prompt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Input indicates an expected call of Input.
func (mr *MockUserInterfaceMockRecorder) Input(prompt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Input", reflect.TypeOf((*MockUserInterface)(nil).Input), prompt)
}

// NewProgressBar mocks base method.
func (m *MockUserInterface) NewProgressBar() ui.ProgressBar {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProgressBar")
	ret0, _ := ret[0].(ui.ProgressBar)
	return ret0
}

// NewProgressBar indicates an expected call of NewProgressBar.
func (mr *MockUserInterfaceMockRecorder) NewProgressBar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProgressBar", reflect.TypeOf((*MockUserInterface)(nil).NewProgressBar))
}

// Output mocks base method.
func (m *MockUserInterface) Output(output string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output", output)
	ret0, _ := ret[0].(error)
	return ret0
}

// Output indicates an expected call of Output.
func (mr *MockUserInterfaceMockRecorder) Output(output interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockUserInterface)(nil).Output), output)
}

// OutputError mocks base method.
func (m *MockUserInterface) OutputError(err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutputError", err)
	ret0, _ := ret[0].(error)
	return ret0
}

// OutputError indicates an expected call of OutputError.
func (mr *MockUserInterfaceMockRecorder) OutputError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputError", reflect.TypeOf((*MockUserInterface)(nil).OutputError), err)
}
