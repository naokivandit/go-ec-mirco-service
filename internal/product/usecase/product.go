//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package usecase

import (
	"context"
	"fmt"

	"example.com/internal/product/domain"
	"example.com/internal/product/domain/repository"
)

type ProductUsecase interface {
	List(ctx context.Context) (domain.Products, error)
	FindByID(ctx context.Context, productID int) (*domain.Product, error)
}

type productUsecase struct {
	productRepository  repository.ProductRepository
	setttingRepository repository.SettingRepository
}

func NewProductUsecase(productRepository repository.ProductRepository, setttingRepository repository.SettingRepository) productUsecase {
	// ProductUsecaseのインスタンスを作成して返す
	return productUsecase{
		productRepository:  productRepository,
		setttingRepository: setttingRepository,
	}
}

func (uc productUsecase) List(ctx context.Context) (domain.Products, error) {
	// 設定の取得
	setting, err := uc.setttingRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	// メンテナンスモードだったらエラーを返す
	if setting.IsMaintenance() {
		return nil, fmt.Errorf("maintenance mode")
	}

	products, err := uc.productRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc productUsecase) FindByID(ctx context.Context, productID int) (*domain.Product, error) {
	// 設定の取得
	setting, err := uc.setttingRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	// メンテナンスモードだったらエラーを返す
	if setting.IsMaintenance() {
		return nil, fmt.Errorf("maintenance mode")
	}

	product, err := uc.productRepository.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
