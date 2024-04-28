package routing

import (
	"net/http"

	"example.com/internal/product/di"
	"example.com/internal/product/handler"
	"example.com/internal/product/usecase"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(app *newrelic.Application, router *mux.Router, deps *di.Dependencies) {
	setupProductRouting(app, router, deps.ProductUsecase)
}

// setupProductRouting 商品関連のルーティングのセットアップを行う
func setupProductRouting(app *newrelic.Application, router *mux.Router, productUsecase usecase.ProductUsecase) {
	productHandler := handler.NewProductHandler(productUsecase)

	// 商品関連のハンドラーを設定
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc(newrelic.WrapHandleFunc(app, "", productHandler.ProductsHandler)).Methods(http.MethodGet)
	productRouter.HandleFunc(newrelic.WrapHandleFunc(app, "/{productID:[0-9]+}", productHandler.ProductsHandler)).Methods(http.MethodGet)
}
