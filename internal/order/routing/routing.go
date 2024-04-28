package routing

import (
	"net/http"

	"example.com/internal/order/di"
	"example.com/internal/order/handler"
	"example.com/internal/order/usecase"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(app *newrelic.Application, router *mux.Router, deps *di.Dependencies) {
	setupOrderRouting(app, router, deps.OrderUsecase)
}

// setupOrderRouting 注文関連のルーティングのセットアップを行う
func setupOrderRouting(app *newrelic.Application, router *mux.Router, orderUsecase usecase.OrderUsecase) {
	orderHandler := handler.NewOrderHandler(orderUsecase)

	// 注文関連のハンドラーを設定
	orderRouter := router.PathPrefix("/orders").Subrouter()
	orderRouter.HandleFunc(newrelic.WrapHandleFunc(app, "", orderHandler.CreateHandler)).Methods(http.MethodPost)
	orderRouter.HandleFunc(newrelic.WrapHandleFunc(app, "", orderHandler.OrdersHandler)).Methods(http.MethodGet)
	orderRouter.HandleFunc(newrelic.WrapHandleFunc(app, "/{order_id}", orderHandler.OrderHandler)).Methods(http.MethodGet)
}
