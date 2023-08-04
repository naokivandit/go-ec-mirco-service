package repository

import (
	"context"

	"example.com/internal/product/domain"
	"example.com/internal/product/infrastructure/dto"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) List(ctx context.Context) (domain.Products, error) {
	var products dto.Products
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return convertProductsToDomain(products), nil
}

func (r *ProductRepository) FindByID(ctx context.Context, productID int) (*domain.Product, error) {
	var product dto.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return convertProductToDomain(product), nil
}

func convertProductsToDomain(dtos dto.Products) domain.Products {
	products := make(domain.Products, 0, len(dtos))
	for _, dto := range dtos {
		product := convertProductToDomain(dto)
		products = append(products, *product)
	}
	return products
}

func convertProductToDomain(product dto.Product) *domain.Product {
	return &domain.Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		ImageURL: product.ImageURL,
	}
}
