package repository

import (
	"context"

	"example.com/internal/order/domain"
	"example.com/internal/order/infrastructure/dto"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, db *gorm.DB, order domain.Order) (*domain.Order, error) {
	dto := dto.Order{
		Name:  order.Name,
		Price: order.Price,
	}
	if err := db.Create(&dto).Error; err != nil {
		return nil, err
	}

	return convertOrderToDomain(dto), nil
}

func (r *OrderRepository) CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem domain.OrderItem) (*domain.OrderItem, error) {
	dto := dto.OrderItem{
		OrderID:  orderItem.OrderID,
		ImageURL: orderItem.ImageURL,
	}
	if err := db.Create(&dto).Error; err != nil {
		return nil, err
	}

	return convertOrderItemToDomain(dto), nil
}

func (r *OrderRepository) List(ctx context.Context) (domain.Orders, error) {
	var orders dto.Orders
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return convertOrdersToDomain(orders), nil
}

func (r *OrderRepository) Get(ctx context.Context, orderID int) (*domain.Order, error) {
	var order dto.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		return nil, err
	}

	return convertOrderToDomain(order), nil
}

func convertOrdersToDomain(dtos dto.Orders) domain.Orders {
	orders := make(domain.Orders, 0, len(dtos))
	for _, o := range dtos {
		order := convertOrderToDomain(o)
		orders = append(orders, *order)
	}
	return orders
}

func convertOrderToDomain(order dto.Order) *domain.Order {
	return &domain.Order{
		ID:    order.ID,
		Name:  order.Name,
		Price: order.Price,
	}
}

func convertOrderItemToDomain(orderItem dto.OrderItem) *domain.OrderItem {
	return &domain.OrderItem{
		ID:       orderItem.ID,
		OrderID:  orderItem.OrderID,
		ImageURL: orderItem.ImageURL,
	}
}
