// Code generated by MockGen. DO NOT EDIT.
// Source: order.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "example.com/internal/order/domain"
	usecase "example.com/internal/order/usecase"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderUsecase is a mock of OrderUsecase interface.
type MockOrderUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockOrderUsecaseMockRecorder
}

// MockOrderUsecaseMockRecorder is the mock recorder for MockOrderUsecase.
type MockOrderUsecaseMockRecorder struct {
	mock *MockOrderUsecase
}

// NewMockOrderUsecase creates a new mock instance.
func NewMockOrderUsecase(ctrl *gomock.Controller) *MockOrderUsecase {
	mock := &MockOrderUsecase{ctrl: ctrl}
	mock.recorder = &MockOrderUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderUsecase) EXPECT() *MockOrderUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderUsecase) Create(ctx context.Context, input usecase.CreateOrderInput) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderUsecaseMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderUsecase)(nil).Create), ctx, input)
}

// Get mocks base method.
func (m *MockOrderUsecase) Get(ctx context.Context, orderID int) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, orderID)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderUsecaseMockRecorder) Get(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderUsecase)(nil).Get), ctx, orderID)
}

// List mocks base method.
func (m *MockOrderUsecase) List(ctx context.Context) (domain.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(domain.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderUsecaseMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderUsecase)(nil).List), ctx)
}