package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/internal/order/usecase"
)

type OrderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase) OrderHandler {
	return OrderHandler{
		orderUsecase: orderUsecase,
	}
}

// 注文を作成するハンドラー関数
func (h *OrderHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを取得
	ctx := r.Context()

	// リクエストボディから商品IDを取得
	var input usecase.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 注文を作成するためのユースケースを呼び出す
	order, err := h.orderUsecase.Create(ctx, input)
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// 注文データをJSONに変換
	jsonData, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// 注文一覧を返すハンドラー関数
func (h *OrderHandler) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを取得
	ctx := r.Context()

	// 注文一覧を取得するためのユースケースを呼び出す
	orders, err := h.orderUsecase.List(ctx)
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	// 注文データをJSONに変換
	jsonData, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// 注文を返すハンドラー関数
func (h *OrderHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを取得
	ctx := r.Context()

	// リクエストパスから注文IDを取得
	paramOrderID := r.URL.Path[len("/orders/"):]

	// string型からint型に変換
	orderID, err := strconv.Atoi(paramOrderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// 注文一覧を取得するためのユースケースを呼び出す
	order, err := h.orderUsecase.Get(ctx, orderID)
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	// 注文データをJSONに変換
	jsonData, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
