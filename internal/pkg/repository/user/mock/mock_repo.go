// Code generated by MockGen. DO NOT EDIT.
// Source: jirno/internal/pkg/domain/user (interfaces: IUserRepo)

// Package mock is a generated GoMock package.
package mock

import (
	user "jirno/internal/pkg/domain/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepo) Create(arg0 user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepo)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockIUserRepo) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepoMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepo)(nil).Delete), arg0)
}

// GetByID mocks base method.
func (m *MockIUserRepo) GetByID(arg0 int64) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIUserRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUserRepo)(nil).GetByID), arg0)
}

// GetByNickname mocks base method.
func (m *MockIUserRepo) GetByNickname(arg0 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNickname", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNickname indicates an expected call of GetByNickname.
func (mr *MockIUserRepoMockRecorder) GetByNickname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNickname", reflect.TypeOf((*MockIUserRepo)(nil).GetByNickname), arg0)
}

// GetMaxUserID mocks base method.
func (m *MockIUserRepo) GetMaxUserID() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxUserID")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxUserID indicates an expected call of GetMaxUserID.
func (mr *MockIUserRepoMockRecorder) GetMaxUserID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxUserID", reflect.TypeOf((*MockIUserRepo)(nil).GetMaxUserID))
}

// Update mocks base method.
func (m *MockIUserRepo) Update(arg0 user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepo)(nil).Update), arg0)
}
