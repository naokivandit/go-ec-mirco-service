package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// ApplyMiddleware ミドルウェアを適用したハンドラーを返す
func ApplyMiddleware(handler http.Handler) http.Handler {
	// CORSミドルウェアのカスタム設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // すべてのアクセス元からのアクセスを許可
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// CORSミドルウェアをハンドラーに適用
	return c.Handler(handler)
}
