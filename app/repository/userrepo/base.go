package userrepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	GetUserByID(id int) (user model.User)
	GetUserByUUID(uuid string) (user model.User)
	GetUserList() (users []model.User)
	CreateUserByMap(userMap map[string]interface{})
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

	if err := db.Where("uuid = ?", uuid).Preload("Members").Find(&user).Error; err != nil {
		log.Fatalf("Get User By UUID Error:%v", err.Error())
	}

	return
}

func (r *repository) CreateUserByMap(userMap map[string]interface{}) {
	db := r.DB.GetConnection()

	if err := db.Model(&model.User{}).Create(userMap).Error; err != nil {
		log.Fatalf("Create User By Map Error:%v", err.Error())
		return
	}

	return
}

func (r *repository) GetUserList() (users []model.User) {
	db := r.DB.GetConnection()

	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("Get User List Error:%v", err.Error())
		return
	}
	return
}

func (r *repository) GetUserByID(id int) (user model.User) {
	db := r.DB.GetConnection()
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Fatalf("Get User By ID Error:%v", err.Error())
	}

	return
}
