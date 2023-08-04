package repository

import (
	"context"
	"reflect"
	"testing"

	"example.com/internal/setting/domain"
	"gorm.io/gorm"
)

func TestSettingRepository_Get(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Setting
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SettingRepository{
				db: tt.fields.db,
			}
			got, err := r.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SettingRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SettingRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
