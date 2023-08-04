//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package repository

import (
	"context"

	"example.com/internal/order/domain"
)

type ProductRepository interface {
	List(ctx context.Context) (domain.Products, error)
	FindByID(ctx context.Context, productID int) (*domain.Product, error)
}
