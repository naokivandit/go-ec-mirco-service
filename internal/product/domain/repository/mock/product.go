// Code generated by MockGen. DO NOT EDIT.
// Source: product.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "example.com/internal/product/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockProductRepository) FindByID(ctx context.Context, productID int) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, productID)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockProductRepositoryMockRecorder) FindByID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockProductRepository)(nil).FindByID), ctx, productID)
}

// List mocks base method.
func (m *MockProductRepository) List(ctx context.Context) (domain.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(domain.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductRepository)(nil).List), ctx)
}
