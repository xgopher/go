package core

import (
	_ "app/conf" // 配置
	"os"

	_ "github.com/go-sql-driver/mysql" // 导入驱动程序
	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB	

func init() {
	initDB()
}

func initDB() {

	conn := os.Getenv("BEEGO_SQLCONN")
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	db.DB().SetMaxIdleConns(10)  // Golang 原生方法: 设置闲置的连接数
	db.DB().SetMaxOpenConns(100) // Golang 原生方法: 最大打开的连接数，默认值为0表示不限制
	db.DB().Ping()				 // Golang 原生方法: ping

	// 表前辍
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "pre_" + defaultTableName
	}

	DB = db
}
