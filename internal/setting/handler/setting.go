package handler

import (
	"encoding/json"
	"net/http"

	"example.com/internal/setting/usecase"
)

type SettingHandler struct {
	settingUsecase usecase.SettingUsecase
}

func NewSettingHandler(settingUsecase usecase.SettingUsecase) SettingHandler {
	return SettingHandler{
		settingUsecase: settingUsecase,
	}
}

// SettingHandler 設定を返す
func (h *SettingHandler) SettingHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを取得
	ctx := r.Context()

	// 設定を取得するためのユースケースを呼び出す
	setting, err := h.settingUsecase.Get(ctx)
	if err != nil {
		// エラーハンドリング
		http.Error(w, "Failed to get setting", http.StatusInternalServerError)
		return
	}

	// 設定データをJSONに変換
	jsonData, err := json.Marshal(setting)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
