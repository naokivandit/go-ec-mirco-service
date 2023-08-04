package routing

import (
	"net/http"

	"example.com/internal/product/di"
	"example.com/internal/product/handler"
	"example.com/internal/product/usecase"
	"github.com/gorilla/mux"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(router *mux.Router, deps *di.Dependencies) {
	setupProductRouting(router, deps.ProductUsecase)
}

// setupProductRouting 商品関連のルーティングのセットアップを行う
func setupProductRouting(router *mux.Router, productUsecase usecase.ProductUsecase) {
	productHandler := handler.NewProductHandler(productUsecase)

	// 商品関連のハンドラーを設定
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", productHandler.ProductsHandler).Methods(http.MethodGet)
	productRouter.HandleFunc("/{productID:[0-9]+}", productHandler.ProductHandler).Methods(http.MethodGet)
}
