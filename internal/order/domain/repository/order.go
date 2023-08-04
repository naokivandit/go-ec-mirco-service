//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package repository

import (
	"context"

	"example.com/internal/order/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, db *gorm.DB, order domain.Order) (*domain.Order, error)
	CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem domain.OrderItem) (*domain.OrderItem, error)
	List(ctx context.Context) (domain.Orders, error)
	Get(ctx context.Context, orderID int) (*domain.Order, error)
}

type OrderItemRepository interface {
}
