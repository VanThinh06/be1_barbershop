// Code generated by MockGen. DO NOT EDIT.
// Source: barbershop/db/sqlc (interfaces: StoreMain)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "barbershop/db/sqlc"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStoreMain is a mock of StoreMain interface.
type MockStoreMain struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMainMockRecorder
}

// MockStoreMainMockRecorder is the mock recorder for MockStoreMain.
type MockStoreMainMockRecorder struct {
	mock *MockStoreMain
}

// NewMockStoreMain creates a new mock instance.
func NewMockStoreMain(ctrl *gomock.Controller) *MockStoreMain {
	mock := &MockStoreMain{ctrl: ctrl}
	mock.recorder = &MockStoreMainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreMain) EXPECT() *MockStoreMainMockRecorder {
	return m.recorder
}

// CreateUsers mocks base method.
func (m *MockStoreMain) CreateUsers(arg0 context.Context, arg1 db.CreateUsersParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUsers indicates an expected call of CreateUsers.
func (mr *MockStoreMainMockRecorder) CreateUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockStoreMain)(nil).CreateUsers), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStoreMain) GetUsers(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMainMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStoreMain)(nil).GetUsers), arg0, arg1)
}
