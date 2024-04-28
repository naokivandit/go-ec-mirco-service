package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/internal/product/common"
	"example.com/internal/product/di"
	"example.com/internal/product/middleware"
	"example.com/internal/product/routing"
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
		newrelic.ConfigAppName("product-service"),
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
	fmt.Println("Server listening on port 8081...")
	http.ListenAndServe(":8081", handler)
}
