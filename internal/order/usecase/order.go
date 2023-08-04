//go:generate mockgen -source $GOFILE -destination mock/$GOFILE
package usecase

import (
	"context"

	"example.com/internal/order/common"
	"example.com/internal/order/domain"
	"example.com/internal/order/domain/repository"
	"gorm.io/gorm"
)

type OrderUsecase interface {
	Create(ctx context.Context, input CreateOrderInput) (*domain.Order, error)
	List(ctx context.Context) (domain.Orders, error)
	Get(ctx context.Context, orderID int) (*domain.Order, error)
}

type orderUsecase struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
	settingRepository repository.SettingRepository
	db                *gorm.DB
}

func NewOrderUsecase(orderRepository repository.OrderRepository, productRepository repository.ProductRepository, settingRepository repository.SettingRepository, db *gorm.DB) orderUsecase {
	return orderUsecase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
		settingRepository: settingRepository,
		db:                db,
	}
}

type CreateOrderInput struct {
	ProductID int `json:"product_id"`
}

func (uc orderUsecase) Create(ctx context.Context, input CreateOrderInput) (*domain.Order, error) {
	// 設定を取得
	setting, err := uc.settingRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	// メンテナンス中の場合はエラー
	if setting.IsMaintenance() {
		return nil, common.ErrMaintenance
	}

	// 商品の一覧を取得
	product, err := uc.productRepository.FindByID(ctx, input.ProductID)
	if err != nil {
		return nil, err
	}

	// トランザクションを開始
	tx := uc.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 注文を保存
	order, err := uc.orderRepository.Create(ctx, tx, domain.Order{
		Name:  product.Name,
		Price: product.Price,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 注文商品を保存
	orderItem, err := uc.orderRepository.CreateOrderItem(ctx, tx, domain.OrderItem{
		OrderID:  order.ID,
		ImageURL: product.ImageURL,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	order.OrderItem = *orderItem

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (uc orderUsecase) List(ctx context.Context) (domain.Orders, error) {
	return uc.orderRepository.List(ctx)
}

func (uc orderUsecase) Get(ctx context.Context, orderID int) (*domain.Order, error) {
	return uc.orderRepository.Get(ctx, orderID)
}
