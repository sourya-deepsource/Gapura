// Code generated by MockGen. DO NOT EDIT.
// Source: models/users.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/vinhut/gapura/models"
	reflect "reflect"
)

// MockUserDatabase is a mock of UserDatabase interface
type MockUserDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockUserDatabaseMockRecorder
}

// MockUserDatabaseMockRecorder is the mock recorder for MockUserDatabase
type MockUserDatabaseMockRecorder struct {
	mock *MockUserDatabase
}

// NewMockUserDatabase creates a new mock instance
func NewMockUserDatabase(ctrl *gomock.Controller) *MockUserDatabase {
	mock := &MockUserDatabase{ctrl: ctrl}
	mock.recorder = &MockUserDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDatabase) EXPECT() *MockUserDatabaseMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockUserDatabase) Find(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockUserDatabaseMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserDatabase)(nil).Find), arg0, arg1, arg2)
}

// Create mocks base method
func (m *MockUserDatabase) Create(arg0 *models.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserDatabaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDatabase)(nil).Create), arg0)
}

// Update mocks base method
func (m *MockUserDatabase) Update() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserDatabaseMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDatabase)(nil).Update))
}

// Delete mocks base method
func (m *MockUserDatabase) Delete(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockUserDatabaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDatabase)(nil).Delete), arg0)
}