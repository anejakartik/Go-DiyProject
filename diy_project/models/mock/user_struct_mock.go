// Code generated by MockGen. DO NOT EDIT.
// Source: user_struct.go

// Package mock is a generated GoMock package.
package mock

import (
	models "diy_project/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// UpdateUserInput mocks base method.
func (m *MockUserInterface) UpdateUserInput(r models.UpdateUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInput", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInput indicates an expected call of UpdateUserInput.
func (mr *MockUserInterfaceMockRecorder) UpdateUserInput(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInput", reflect.TypeOf((*MockUserInterface)(nil).UpdateUserInput), r)
}

// Users mocks base method.
func (m *MockUserInterface) Users(r models.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockUserInterfaceMockRecorder) Users(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockUserInterface)(nil).Users), r)
}
