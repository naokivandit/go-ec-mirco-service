package common

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	// GORMの設定
	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // ログ出力先とフォーマット
			logger.Config{
				LogLevel: logger.Info, // ログレベル (Silent, Error, Warn, Info)
			},
		),
	}

	// MySQL接続文字列の取得
	dbConnectionString, err := getDBConnectionString()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dbConnectionString), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDBConnectionString() (string, error) {
	config, err := NewConfig()
	if err != nil {
		return "", err
	}

	// 環境変数の取得
	dbUser := config.MySQL.Username
	dbPassword := config.MySQL.Password
	dbHost := config.MySQL.Host
	dbPort := config.MySQL.Port
	dbName := config.MySQL.Database

	// MySQL接続文字列の構築
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dbConnectionString, nil
}
