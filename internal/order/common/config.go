package common

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}
}

var instance *Config

// NewConfig は設定を初期化して返します
// シングルトンパターンを採用
func NewConfig() (*Config, error) {
	if instance == nil {
		// デフォルト値を設定
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		viper.SetDefault("DB_HOST", "localhost")
		viper.SetDefault("DB_PORT", 3306)
		viper.SetDefault("DB_USER", "root")
		viper.SetDefault("DB_PASSWORD", "password")
		viper.SetDefault("DB_ORDER_SERVICE_NAME", "order")

		// Config構造体にマッピング
		if err := viper.Unmarshal(&instance); err != nil {
			return nil, err
		}

		// 環境変数を読み込み
		viper.AutomaticEnv()

		// マッピング後の値を取得
		instance.MySQL.Host = viper.GetString("DB_HOST")
		instance.MySQL.Port = viper.GetInt("DB_PORT")
		instance.MySQL.Username = viper.GetString("DB_USER")
		instance.MySQL.Password = viper.GetString("DB_PASSWORD")
		instance.MySQL.Database = viper.GetString("DB_ORDER_SERVICE_NAME")
	}

	return instance, nil
}
