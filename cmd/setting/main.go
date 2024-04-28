package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/internal/setting/common"
	"example.com/internal/setting/di"
	"example.com/internal/setting/middleware"
	"example.com/internal/setting/routing"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	// 設定の初期化
	_, err := common.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("setting-service"),
		newrelic.ConfigLicense("YOUR_LICENSE_KEY"),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		log.Fatal("Error initializing New Relic:", err)
	}
	defer app.Shutdown(10 * time.Second)

	// データベースへの接続
	db, err := common.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// DIのセットアップ
	deps := di.SetupDI(db)

	// ルーティングのセットアップ
	router := mux.NewRouter()
	routing.SetupRouting(app, router, deps)

	// ミドルウェアのセットアップ
	handler := middleware.ApplyMiddleware(router)

	// サーバーの起動
	fmt.Println("Server listening on port 8082...")
	http.ListenAndServe(":8082", handler)
}
