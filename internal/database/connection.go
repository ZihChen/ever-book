package database

import (
	"ever-book/app/global/errorcode"
	"ever-book/app/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var connectPool *gorm.DB

type Interface interface {
	GetConnection() *gorm.DB
	AutoMigrate()
}

type instance struct{}

func New() Interface {
	return &instance{}
}

func (db *instance) GetConnection() *gorm.DB {
	var err error
	if connectPool != nil {
		return connectPool
	}

	connectPool, err = gorm.Open(mysql.Open("root:root@tcp(ever-book-db:3306)/everbook?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf(errorcode.CheckDBConnectError, err.Error())
	}

	sqlPool, err := connectPool.DB()
	if err != nil {
		log.Fatalf(errorcode.DBConnectPoolError, err.Error())
	}

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	fmt.Println("db connect success!")

	return connectPool.Debug()
}

func (db *instance) AutoMigrate() {
	database := db.GetConnection()

	err := database.Set("gorm:table_options", "comment '使用者'").AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf(errorcode.AutoMigrateError, err.Error())
	}
}
