package database

import (
	"ever-book/app/global/errorcode"
	"ever-book/app/global/settings"
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

	connectPool, err = gorm.Open(mysql.Open(getDSN()), &gorm.Config{
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

	err = database.Set("gorm:table_options", "comment '暫存收支紀錄'").AutoMigrate(&model.TemporaryBalance{})
	if err != nil {
		log.Fatalf(errorcode.AutoMigrateError, err.Error())
	}

	err = database.Set("gorm:table_options", "comment '收支紀錄'").AutoMigrate(&model.DailyBalance{})
	if err != nil {
		log.Fatalf(errorcode.AutoMigrateError, err.Error())
	}
}

func getDSN() string {
	Host := settings.Config.DBConfig.Host
	Username := settings.Config.DBConfig.Username
	Password := settings.Config.DBConfig.Password
	Database := settings.Config.DBConfig.Database

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Database)
}
