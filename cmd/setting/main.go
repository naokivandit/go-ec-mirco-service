package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/internal/setting/common"
	"example.com/internal/setting/di"
	"example.com/internal/setting/middleware"
	"example.com/internal/setting/routing"
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
	fmt.Println("Server listening on port 8082...")
	http.ListenAndServe(":8082", handler)
}
