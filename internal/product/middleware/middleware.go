package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// ApplyMiddleware ミドルウェアを適用したハンドラーを返す
func ApplyMiddleware(handler http.Handler) http.Handler {
	// CORSミドルウェアの設定
	corsHandler := cors.Default().Handler(handler)

	return corsHandler
}
