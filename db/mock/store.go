// Code generated by MockGen. DO NOT EDIT.
// Source: barbershop/db/sqlc (interfaces: StoreMain)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "barbershop/db/sqlc"
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
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

// CreateBarber mocks base method.
func (m *MockStoreMain) CreateBarber(arg0 context.Context, arg1 db.CreateBarberParams) (db.Barber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBarber", arg0, arg1)
	ret0, _ := ret[0].(db.Barber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBarber indicates an expected call of CreateBarber.
func (mr *MockStoreMainMockRecorder) CreateBarber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBarber", reflect.TypeOf((*MockStoreMain)(nil).CreateBarber), arg0, arg1)
}

// CreateSessionBarber mocks base method.
func (m *MockStoreMain) CreateSessionBarber(arg0 context.Context, arg1 db.CreateSessionBarberParams) (db.SessionsBarber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSessionBarber", arg0, arg1)
	ret0, _ := ret[0].(db.SessionsBarber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSessionBarber indicates an expected call of CreateSessionBarber.
func (mr *MockStoreMainMockRecorder) CreateSessionBarber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSessionBarber", reflect.TypeOf((*MockStoreMain)(nil).CreateSessionBarber), arg0, arg1)
}

// GetBarber mocks base method.
func (m *MockStoreMain) GetBarber(arg0 context.Context, arg1 string) (db.GetBarberRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBarber", arg0, arg1)
	ret0, _ := ret[0].(db.GetBarberRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBarber indicates an expected call of GetBarber.
func (mr *MockStoreMainMockRecorder) GetBarber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBarber", reflect.TypeOf((*MockStoreMain)(nil).GetBarber), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStoreMain) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.SessionsBarber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.SessionsBarber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMainMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStoreMain)(nil).GetSession), arg0, arg1)
}
