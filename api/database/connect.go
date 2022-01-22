package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	schema = "%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// docker-compose.ymlに設定した環境変数を取得
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(schema, username, password, dbName)
	DB             *gorm.DB // DBインスタンス
)

func Connect() {
	connection, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	// DNSサーバが見つからなかった場合
	if err != nil {
		panic("Could not connect")
	}

	// コネクション情報を追加
	DB = connection

	// connection.AutoMigrate(
	// )
}
