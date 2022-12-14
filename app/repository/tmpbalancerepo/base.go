package tmpbalancerepo

import (
	"ever-book/app/global"
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	CheckTemporaryBalanceExistByUserID(userID int) (exist bool)
	GetTemporaryBalanceByUserID(userID int) (tmpBalance model.TemporaryBalance)
	CreateTemporaryBalanceByMap(fields map[string]interface{})
	GetLatestTemporaryBalanceByUserID(userID int) (tmpBalance model.TemporaryBalance)
	UpdateTemporaryBalanceByID(tmpBalanceID int, column global.TemporaryBalanceColumn, value interface{})
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

func (r *repository) CreateTemporaryBalanceByMap(fields map[string]interface{}) {
	db := r.DB.GetConnection()

	if err := db.Model(&model.TemporaryBalance{}).Create(fields).Error; err != nil {
		log.Fatalf("Create Temporary Balance By Map Error:%v", err.Error())
	}
}

func (r *repository) GetLatestTemporaryBalanceByUserID(userID int) (tmpBalance model.TemporaryBalance) {
	db := r.DB.GetConnection()

	if err := db.Where("user_id = ?", userID).Last(&tmpBalance).Error; err != nil {
		log.Fatalf("Get Latest Temporary Balance By UserID Error:%v", err.Error())
	}

	return
}

func (r *repository) UpdateTemporaryBalanceByID(tmpBalanceID int, column global.TemporaryBalanceColumn, value interface{}) {
	db := r.DB.GetConnection()

	if err := db.Model(&model.TemporaryBalance{}).Where("id = ?", tmpBalanceID).
		Update(string(column), value).Error; err != nil {
		log.Fatalf("Update Temporary Balance By ID Error:%v", err.Error())
	}
}

func (r *repository) GetTemporaryBalanceByUserID(userID int) (tmpBalance model.TemporaryBalance) {
	db := r.DB.GetConnection()

	if err := db.Where("user_id = ?", userID).Find(&tmpBalance).Error; err != nil {
		log.Fatalf("Get Temporary Balance By USer ID Error:%v", err.Error())
	}
	return
}

func (r *repository) CheckTemporaryBalanceExistByUserID(userID int) (exist bool) {
	db := r.DB.GetConnection()
	var count int64
	if err := db.Model(&model.TemporaryBalance{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		log.Fatalf("Check Temporary Balance Exist By USer ID Error:%v", err.Error())
	}
	if count > 0 {
		return true
	}
	return
}
