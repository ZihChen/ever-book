package tmpbalancerepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	CreateTemporaryBalanceByMap(fields map[string]interface{})
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
