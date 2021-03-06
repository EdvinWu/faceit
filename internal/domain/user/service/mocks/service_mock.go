// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	model "faceit-test/internal/domain/user/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(ctx context.Context, user model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), ctx, userID)
}

// PaginatedUsers mocks base method.
func (m *MockUser) PaginatedUsers(ctx context.Context, user model.PagedUserRequest) ([]*model.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginatedUsers", ctx, user)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PaginatedUsers indicates an expected call of PaginatedUsers.
func (mr *MockUserMockRecorder) PaginatedUsers(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginatedUsers", reflect.TypeOf((*MockUser)(nil).PaginatedUsers), ctx, user)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), ctx, user)
}
