package routing

import (
	"net/http"

	"example.com/internal/setting/di"
	"example.com/internal/setting/handler"
	"example.com/internal/setting/usecase"
	"github.com/gorilla/mux"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(router *mux.Router, deps *di.Dependencies) {
	setupSettingRouting(router, deps.SettingUsecase)
}

// setupSettingRouting 設定関連のルーティングのセットアップを行う
func setupSettingRouting(router *mux.Router, settingUsecase usecase.SettingUsecase) {
	settingHandler := handler.NewSettingHandler(settingUsecase)

	// 設定関連のハンドラーを設定
	settingRouter := router.PathPrefix("/setting").Subrouter()
	settingRouter.HandleFunc("", settingHandler.SettingHandler).Methods(http.MethodGet)
}
