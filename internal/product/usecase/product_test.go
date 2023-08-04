//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package usecase

import (
	"context"
	"reflect"
	"testing"

	mock_repository "example.com/internal/product/domain/repository/mock"

	"example.com/internal/product/domain"
	"github.com/golang/mock/gomock"
)

type mockProducts struct {
	products domain.Products
	err      error
}
type mockSetting struct {
	setting *domain.Setting
	err     error
}

func Test_productUsecase_List(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name            string
		mockListProduct mockProducts
		mockSetting     mockSetting
		args            args
		want            domain.Products
		wantErr         bool
	}{
		{
			name: "正常系",
			mockListProduct: mockProducts{
				products: domain.Products{
					{
						ID:   1,
						Name: "product1",
					},
					{
						ID:   2,
						Name: "product2",
					},
				},
				err: nil,
			},
			mockSetting: mockSetting{
				setting: &domain.Setting{
					MentenanceMode: false,
				},
				err: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want: domain.Products{
				{
					ID:   1,
					Name: "product1",
				},
				{
					ID:   2,
					Name: "product2",
				},
			},
			wantErr: false,
		},
		{
			name: "メンテナンスモード",
			mockListProduct: mockProducts{
				products: nil,
				err:      nil,
			},
			mockSetting: mockSetting{
				setting: &domain.Setting{
					MentenanceMode: true,
				},
				err: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			productRepo := mock_repository.NewMockProductRepository(ctrl)
			if tt.mockListProduct.products != nil {
				productRepo.EXPECT().List(tt.args.ctx).Return(tt.mockListProduct.products, tt.mockListProduct.err)
			}

			settingRepo := mock_repository.NewMockSettingRepository(ctrl)
			if tt.mockSetting.setting != nil {
				settingRepo.EXPECT().Get(tt.args.ctx).Return(tt.mockSetting.setting, tt.mockSetting.err)
			}

			uc := NewProductUsecase(productRepo, settingRepo)

			got, err := uc.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("productUsecase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productUsecase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
