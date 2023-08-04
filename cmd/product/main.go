package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/internal/product/common"
	"example.com/internal/product/di"
	"example.com/internal/product/middleware"
	"example.com/internal/product/routing"
	"github.com/gorilla/mux"
)

func main() {
	// 設定の初期化
	_, err := common.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// データベースへの接続
	db, err := common.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// DIのセットアップ
	deps := di.SetupDI(db)

	// ルーティングのセットアップ
	router := mux.NewRouter()
	routing.SetupRouting(router, deps)

	// ミドルウェアのセットアップ
	handler := middleware.ApplyMiddleware(router)

	// サーバーの起動
	fmt.Println("Server listening on port 8081...")
	http.ListenAndServe(":8081", handler)
}
