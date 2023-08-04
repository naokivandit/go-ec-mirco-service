//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package usecase

import (
	"context"

	"example.com/internal/setting/domain"
	"example.com/internal/setting/domain/repository"
)

type SettingUsecase interface {
	Get(ctx context.Context) (*domain.Setting, error)
}

type settingUsecase struct {
	settingRepository repository.SettingRepository
}

func NewSettingUsecase(settingRepository repository.SettingRepository) settingUsecase {
	// SettingUsecaseのインスタンスを作成して返す
	return settingUsecase{
		settingRepository: settingRepository,
	}
}

func (uc settingUsecase) Get(ctx context.Context) (*domain.Setting, error) {
	setting, err := uc.settingRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	return setting, nil
}
