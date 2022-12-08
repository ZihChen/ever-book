package userrepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	GetUserByUUID(uuid string) (user model.User)
	CreateUserByUUID(uuid string) (user model.User)
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

func (r *repository) GetUserByUUID(uuid string) (user model.User) {
	db := r.DB.GetConnection()

	if err := db.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		log.Fatalf("Get User By UUID Error:%v", err.Error())
	}

	return
}

func (r *repository) CreateUserByUUID(uuid string) (user model.User) {
	db := r.DB.GetConnection()

	if err := db.Model(&user).Create(map[string]interface{}{
		"uuid": uuid,
	}).Error; err != nil {
		log.Fatalf("Create User By UUID Error:%v", err.Error())
		return
	}

	return
}
