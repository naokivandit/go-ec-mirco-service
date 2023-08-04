package di

import (
	"example.com/internal/order/infrastructure/repository"
	"example.com/internal/order/usecase"
	"gorm.io/gorm"
)

type Dependencies struct {
	OrderUsecase usecase.OrderUsecase
}

// SetupDI 依存関係のセットアップを行う
func SetupDI(db *gorm.DB) *Dependencies {
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository()
	settingRepo := repository.NewSettingRepository()
	orderUsecase := usecase.NewOrderUsecase(orderRepo, productRepo, settingRepo, db)

	return &Dependencies{
		OrderUsecase: orderUsecase,
	}
}
