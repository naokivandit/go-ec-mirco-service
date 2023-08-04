package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/internal/order/domain"
)

type SettingRepository struct {
}

func NewSettingRepository() *SettingRepository {
	return &SettingRepository{}
}

func (r *SettingRepository) Get(ctx context.Context) (*domain.Setting, error) {
	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの作成
	req, err := http.NewRequest("GET", "http://setting:8082/setting", nil)
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

	var response domain.Setting
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Failed to deserialize response:", err)
		return nil, err
	}

	return &response, nil
}
