// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dyleme/Notifier/internal/service/service (interfaces: DefaultNotificationParamsRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/notification_params_mocks.go -package=mocks . DefaultNotificationParamsRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domains "github.com/Dyleme/Notifier/internal/domains"
	gomock "go.uber.org/mock/gomock"
)

// MockDefaultNotificationParamsRepository is a mock of DefaultNotificationParamsRepository interface.
type MockDefaultNotificationParamsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDefaultNotificationParamsRepositoryMockRecorder
}

// MockDefaultNotificationParamsRepositoryMockRecorder is the mock recorder for MockDefaultNotificationParamsRepository.
type MockDefaultNotificationParamsRepositoryMockRecorder struct {
	mock *MockDefaultNotificationParamsRepository
}

// NewMockDefaultNotificationParamsRepository creates a new mock instance.
func NewMockDefaultNotificationParamsRepository(ctrl *gomock.Controller) *MockDefaultNotificationParamsRepository {
	mock := &MockDefaultNotificationParamsRepository{ctrl: ctrl}
	mock.recorder = &MockDefaultNotificationParamsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDefaultNotificationParamsRepository) EXPECT() *MockDefaultNotificationParamsRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDefaultNotificationParamsRepository) Get(arg0 context.Context, arg1 int) (domains.NotificationParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(domains.NotificationParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDefaultNotificationParamsRepositoryMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDefaultNotificationParamsRepository)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockDefaultNotificationParamsRepository) Set(arg0 context.Context, arg1 int, arg2 domains.NotificationParams) (domains.NotificationParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(domains.NotificationParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockDefaultNotificationParamsRepositoryMockRecorder) Set(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDefaultNotificationParamsRepository)(nil).Set), arg0, arg1, arg2)
}
