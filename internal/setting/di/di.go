package di

import (
	"example.com/internal/setting/infrastructure/repository"
	"example.com/internal/setting/usecase"
	"gorm.io/gorm"
)

type Dependencies struct {
	SettingUsecase usecase.SettingUsecase
}

// SetupDI 依存関係のセットアップを行う
func SetupDI(db *gorm.DB) *Dependencies {
	settingRepo := repository.NewSettingRepository(db)
	settingUsecase := usecase.NewSettingUsecase(settingRepo)

	return &Dependencies{
		SettingUsecase: settingUsecase,
	}
}
