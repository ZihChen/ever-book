package usermemberrepo

import (
	"ever-book/app/model"
	"ever-book/internal/database"
	"log"
	"sync"
)

type Interface interface {
	CreateUserMembersRelation(user model.User, members []model.User)
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

func (r *repository) CreateUserMembersRelation(user model.User, members []model.User) {
	db := r.DB.GetConnection()

	if err := db.Model(&user).Association("Members").Replace(&members); err != nil {
		log.Fatalf("Create User Members Relation Error:%v", err.Error())
	}
}
