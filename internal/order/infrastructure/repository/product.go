package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/internal/order/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) List(ctx context.Context) (domain.Products, error) {
	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの作成
	req, err := http.NewRequest("GET", "http://product:8081/products", nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return nil, err
	}

	// リクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return nil, err
	}

	// jsonからdomain.Productに変換
	var response domain.Products
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Failed to deserialize response:", err)
		return nil, err
	}

	return response, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, productID int) (*domain.Product, error) {
	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの作成
	req, err := http.NewRequest("GET", fmt.Sprintf("http://product:8081/products/%d", productID), nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return nil, err
	}

	// リクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return nil, err
	}

	// jsonからdomain.Productに変換
	var response domain.Product
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Failed to deserialize response:", err)
		return nil, err
	}

	return &response, nil
}
