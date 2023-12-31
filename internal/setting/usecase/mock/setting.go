// Code generated by MockGen. DO NOT EDIT.
// Source: setting.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "example.com/internal/setting/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSettingUsecase is a mock of SettingUsecase interface.
type MockSettingUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockSettingUsecaseMockRecorder
}

// MockSettingUsecaseMockRecorder is the mock recorder for MockSettingUsecase.
type MockSettingUsecaseMockRecorder struct {
	mock *MockSettingUsecase
}

// NewMockSettingUsecase creates a new mock instance.
func NewMockSettingUsecase(ctrl *gomock.Controller) *MockSettingUsecase {
	mock := &MockSettingUsecase{ctrl: ctrl}
	mock.recorder = &MockSettingUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingUsecase) EXPECT() *MockSettingUsecaseMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSettingUsecase) Get(ctx context.Context) (*domain.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(*domain.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSettingUsecaseMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSettingUsecase)(nil).Get), ctx)
}
