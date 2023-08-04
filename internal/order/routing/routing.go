package routing

import (
	"net/http"

	"example.com/internal/order/di"
	"example.com/internal/order/handler"
	"example.com/internal/order/usecase"
	"github.com/gorilla/mux"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(router *mux.Router, deps *di.Dependencies) {
	setupOrderRouting(router, deps.OrderUsecase)
}

// setupOrderRouting 注文関連のルーティングのセットアップを行う
func setupOrderRouting(router *mux.Router, orderUsecase usecase.OrderUsecase) {
	orderHandler := handler.NewOrderHandler(orderUsecase)

	// 注文関連のハンドラーを設定
	orderRouter := router.PathPrefix("/orders").Subrouter()
	orderRouter.HandleFunc("", orderHandler.CreateHandler).Methods(http.MethodPost)
	orderRouter.HandleFunc("", orderHandler.OrdersHandler).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{order_id}", orderHandler.OrderHandler).Methods(http.MethodGet)
}
