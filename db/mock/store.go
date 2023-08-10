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

// CreateService mocks base method.
func (m *MockStoreMain) CreateService(arg0 context.Context, arg1 db.CreateServiceParams) (db.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1)
	ret0, _ := ret[0].(db.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockStoreMainMockRecorder) CreateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockStoreMain)(nil).CreateService), arg0, arg1)
}

// CreateServiceCategory mocks base method.
func (m *MockStoreMain) CreateServiceCategory(arg0 context.Context, arg1 db.CreateServiceCategoryParams) (db.ServiceCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ServiceCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServiceCategory indicates an expected call of CreateServiceCategory.
func (mr *MockStoreMainMockRecorder) CreateServiceCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceCategory", reflect.TypeOf((*MockStoreMain)(nil).CreateServiceCategory), arg0, arg1)
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

// CreateStore mocks base method.
func (m *MockStoreMain) CreateStore(arg0 context.Context, arg1 db.CreateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStore indicates an expected call of CreateStore.
func (mr *MockStoreMainMockRecorder) CreateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockStoreMain)(nil).CreateStore), arg0, arg1)
}

// DeleteService mocks base method.
func (m *MockStoreMain) DeleteService(arg0 context.Context, arg1 uuid.UUID) (db.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(db.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockStoreMainMockRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockStoreMain)(nil).DeleteService), arg0, arg1)
}

// DeleteServiceCategory mocks base method.
func (m *MockStoreMain) DeleteServiceCategory(arg0 context.Context, arg1 uuid.UUID) (db.ServiceCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ServiceCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteServiceCategory indicates an expected call of DeleteServiceCategory.
func (mr *MockStoreMainMockRecorder) DeleteServiceCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceCategory", reflect.TypeOf((*MockStoreMain)(nil).DeleteServiceCategory), arg0, arg1)
}

// DeleteServicewithStoreCategory mocks base method.
func (m *MockStoreMain) DeleteServicewithStoreCategory(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServicewithStoreCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServicewithStoreCategory indicates an expected call of DeleteServicewithStoreCategory.
func (mr *MockStoreMainMockRecorder) DeleteServicewithStoreCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServicewithStoreCategory", reflect.TypeOf((*MockStoreMain)(nil).DeleteServicewithStoreCategory), arg0, arg1)
}

// GetBarber mocks base method.
func (m *MockStoreMain) GetBarber(arg0 context.Context, arg1 string) (db.Barber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBarber", arg0, arg1)
	ret0, _ := ret[0].(db.Barber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBarber indicates an expected call of GetBarber.
func (mr *MockStoreMainMockRecorder) GetBarber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBarber", reflect.TypeOf((*MockStoreMain)(nil).GetBarber), arg0, arg1)
}

// GetListServiceCategorywithStore mocks base method.
func (m *MockStoreMain) GetListServiceCategorywithStore(arg0 context.Context, arg1 uuid.UUID) ([]db.ServiceCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListServiceCategorywithStore", arg0, arg1)
	ret0, _ := ret[0].([]db.ServiceCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListServiceCategorywithStore indicates an expected call of GetListServiceCategorywithStore.
func (mr *MockStoreMainMockRecorder) GetListServiceCategorywithStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListServiceCategorywithStore", reflect.TypeOf((*MockStoreMain)(nil).GetListServiceCategorywithStore), arg0, arg1)
}

// GetListServicewithCategory mocks base method.
func (m *MockStoreMain) GetListServicewithCategory(arg0 context.Context, arg1 db.GetListServicewithCategoryParams) ([]db.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListServicewithCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListServicewithCategory indicates an expected call of GetListServicewithCategory.
func (mr *MockStoreMainMockRecorder) GetListServicewithCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListServicewithCategory", reflect.TypeOf((*MockStoreMain)(nil).GetListServicewithCategory), arg0, arg1)
}

// GetListStore mocks base method.
func (m *MockStoreMain) GetListStore(arg0 context.Context, arg1 db.GetListStoreParams) ([]db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListStore", arg0, arg1)
	ret0, _ := ret[0].([]db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListStore indicates an expected call of GetListStore.
func (mr *MockStoreMainMockRecorder) GetListStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListStore", reflect.TypeOf((*MockStoreMain)(nil).GetListStore), arg0, arg1)
}

// GetService mocks base method.
func (m *MockStoreMain) GetService(arg0 context.Context, arg1 uuid.UUID) (db.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(db.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockStoreMainMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockStoreMain)(nil).GetService), arg0, arg1)
}

// GetServiceCategory mocks base method.
func (m *MockStoreMain) GetServiceCategory(arg0 context.Context, arg1 uuid.UUID) (db.ServiceCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ServiceCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceCategory indicates an expected call of GetServiceCategory.
func (mr *MockStoreMainMockRecorder) GetServiceCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCategory", reflect.TypeOf((*MockStoreMain)(nil).GetServiceCategory), arg0, arg1)
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

// GetStore mocks base method.
func (m *MockStoreMain) GetStore(arg0 context.Context, arg1 uuid.UUID) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockStoreMainMockRecorder) GetStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockStoreMain)(nil).GetStore), arg0, arg1)
}

// UpdateService mocks base method.
func (m *MockStoreMain) UpdateService(arg0 context.Context, arg1 db.UpdateServiceParams) (db.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1)
	ret0, _ := ret[0].(db.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockStoreMainMockRecorder) UpdateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockStoreMain)(nil).UpdateService), arg0, arg1)
}
