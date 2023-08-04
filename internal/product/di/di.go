package di

import (
	"example.com/internal/product/infrastructure/repository"
	"example.com/internal/product/usecase"
	"gorm.io/gorm"
)

type Dependencies struct {
	ProductUsecase usecase.ProductUsecase
}

// SetupDI 依存関係のセットアップを行う
func SetupDI(db *gorm.DB) *Dependencies {
	productRepo := repository.NewProductRepository(db)
	settingRepo := repository.NewSettingRepository()
	productUsecase := usecase.NewProductUsecase(productRepo, settingRepo)

	return &Dependencies{
		ProductUsecase: productUsecase,
	}
}
