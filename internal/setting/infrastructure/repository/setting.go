package repository

import (
	"context"

	"example.com/internal/setting/domain"
	"example.com/internal/setting/infrastructure/dto"
	"gorm.io/gorm"
)

type SettingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) *SettingRepository {
	return &SettingRepository{
		db: db,
	}
}

// Get 設定を取得する
func (r *SettingRepository) Get(ctx context.Context) (*domain.Setting, error) {
	var setting dto.Setting
	if err := r.db.First(&setting).Error; err != nil {
		return nil, err
	}
	return convertSettingToDomain(setting), nil
}

func convertSettingToDomain(setting dto.Setting) *domain.Setting {
	return &domain.Setting{
		ID:             setting.ID,
		MentenanceMode: setting.MentenanceMode,
	}
}
