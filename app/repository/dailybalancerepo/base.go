package dailybalancerepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	CreateDailyBalanceByMap(fields map[string]interface{})
	GatLatestDailyBalanceByUserID(userID int) (dailyBalance model.DailyBalance)
	CheckDailyBalanceExistByUserID(userID int) (exist bool)
	DeleteDailyBalanceByID(id int)
}

type repository struct {
	DB database.Interface
}

var singleton *repository
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &repository{
			DB: database.New(),
		}
	})
	return singleton
}

func (r *repository) CreateDailyBalanceByMap(fields map[string]interface{}) {
	db := r.DB.GetConnection()

	if err := db.Model(&model.DailyBalance{}).Create(fields).Error; err != nil {
		log.Fatalf("Create Daily Balance By Map Error:%v", err.Error())
	}
}

func (r *repository) GatLatestDailyBalanceByUserID(userID int) (dailyBalance model.DailyBalance) {
	db := r.DB.GetConnection()

	if err := db.Where("user_id = ?", userID).Last(&dailyBalance).Error; err != nil {
		log.Fatalf("Get Latest Daily Balance By UserID Error:%v", err.Error())
	}
	return
}

func (r *repository) CheckDailyBalanceExistByUserID(userID int) (exist bool) {
	db := r.DB.GetConnection()
	var count int64
	if err := db.Model(&model.DailyBalance{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		log.Fatalf("Check Daily Balance Exist By USer ID Error:%v", err.Error())
	}
	if count > 0 {
		return true
	}
	return
}

func (r *repository) DeleteDailyBalanceByID(id int) {
	db := r.DB.GetConnection()

	if err := db.Where("id = ?", id).Delete(&model.DailyBalance{}).Error; err != nil {
		log.Fatalf("Delete Daily Balance By ID Error:%v", err.Error())
	}
}
