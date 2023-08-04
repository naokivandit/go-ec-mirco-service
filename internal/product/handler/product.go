package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/internal/product/usecase"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) ProductHandler {
	return ProductHandler{
		productUsecase: productUsecase,
	}
}

// 商品一覧を返すハンドラー関数
func (h *ProductHandler) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// 商品一覧を取得するためのユースケースを呼び出す
	products, err := h.productUsecase.List(r.Context())
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	// 商品データをJSONに変換
	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// 商品詳細を返すハンドラー関数
func (h *ProductHandler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを取得
	ctx := r.Context()

	// URLから商品IDを取得
	productID, err := getProductID(r)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// 商品詳細を取得するためのユースケースを呼び出す
	product, err := h.productUsecase.FindByID(ctx, productID)
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	// 商品データをJSONに変換
	jsonData, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// URLから商品IDを取得する
func getProductID(r *http.Request) (int, error) {
	// URLから商品IDを取得
	vars := mux.Vars(r)
	productID, ok := vars["productID"]
	if !ok {
		return 0, fmt.Errorf("productID not found")
	}

	// 商品IDを数値に変換
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return 0, fmt.Errorf("productID is not a number")
	}

	return productIDInt, nil
}
