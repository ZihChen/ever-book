package userrepo

import (
	"ever-book/app/global"
	"ever-book/app/global/structs"
	"ever-book/app/model"
	"ever-book/internal/database"
	"gorm.io/gorm"
	"log"
	"sync"
)

type Interface interface {
	GetUserByID(id int) (user model.User)
	GetUserList() (users []model.User)
	CreateUserByMap(userMap map[string]interface{})
	UpdateUserByID(id int, column global.UserColumn, value interface{})
	GetUserInfo(params structs.UserParams) (user model.User)
}

type repository struct {
	DB database.Interface
}

type Option func(*gorm.DB) *gorm.DB

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

func (r *repository) GetUserInfo(params structs.UserParams) (user model.User) {
	db := r.DB.GetConnection()

	if params.ID > 0 {
		db = db.Where("id = ?", params.ID)
	}

	if params.UUID != "" {
		db = db.Where("uuid = ?", params.UUID)
	}

	if err := db.Preload("Members").Find(&user).Error; err != nil {
		log.Fatalf("Get User Error:%v", err.Error())
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

func (r *repository) UpdateUserByID(id int, column global.UserColumn, value interface{}) {
	db := r.DB.GetConnection()

	if err := db.Model(&model.User{}).Where("id = ?", id).
		Update(string(column), value).Error; err != nil {
		log.Fatalf("Update User By ID Error:%v", err.Error())
	}
}
