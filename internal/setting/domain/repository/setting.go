//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package repository

import (
	"context"

	"example.com/internal/setting/domain"
)

type SettingRepository interface {
	Get(ctx context.Context) (*domain.Setting, error)
}
