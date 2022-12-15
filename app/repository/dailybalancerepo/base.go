package dailybalancerepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	CreateDailyBalanceByMap(fields map[string]interface{})
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
