package routing

import (
	"net/http"

	"example.com/internal/setting/di"
	"example.com/internal/setting/handler"
	"example.com/internal/setting/usecase"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupRouting ルーティングのセットアップを行う
func SetupRouting(app *newrelic.Application, router *mux.Router, deps *di.Dependencies) {
	setupSettingRouting(app, router, deps.SettingUsecase)
}

// setupSettingRouting 設定関連のルーティングのセットアップを行う
func setupSettingRouting(app *newrelic.Application, router *mux.Router, settingUsecase usecase.SettingUsecase) {
	settingHandler := handler.NewSettingHandler(settingUsecase)

	// 設定関連のハンドラーを設定
	settingRouter := router.PathPrefix("/setting").Subrouter()
	settingRouter.HandleFunc(newrelic.WrapHandleFunc(app, "", settingHandler.SettingHandler)).Methods(http.MethodGet)
}
