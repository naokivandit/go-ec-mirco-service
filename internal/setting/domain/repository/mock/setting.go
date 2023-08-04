// Code generated by MockGen. DO NOT EDIT.
// Source: setting.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "example.com/internal/setting/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSettingRepository is a mock of SettingRepository interface.
type MockSettingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSettingRepositoryMockRecorder
}

// MockSettingRepositoryMockRecorder is the mock recorder for MockSettingRepository.
type MockSettingRepositoryMockRecorder struct {
	mock *MockSettingRepository
}

// NewMockSettingRepository creates a new mock instance.
func NewMockSettingRepository(ctrl *gomock.Controller) *MockSettingRepository {
	mock := &MockSettingRepository{ctrl: ctrl}
	mock.recorder = &MockSettingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingRepository) EXPECT() *MockSettingRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSettingRepository) Get(ctx context.Context) (*domain.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(*domain.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSettingRepositoryMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSettingRepository)(nil).Get), ctx)
}
