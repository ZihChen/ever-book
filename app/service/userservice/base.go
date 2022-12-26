package userservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/model"
	"ever-book/app/repository/userrepo"
	"sync"
)

type Interface interface {
	GetOrCreateUser(userFields structs.UserFields) (user model.User)
}

type service struct {
	UserRepo userrepo.Interface
}

var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{
			UserRepo: userrepo.New(),
		}
	})
	return singleton
}

func (s *service) GetOrCreateUser(userFields structs.UserFields) (user model.User) {
	user = s.UserRepo.GetUserByUUID(userFields.UUID)
	if user.ID != 0 {
		return
	}
	s.UserRepo.CreateUserByMap(helper.StructToMap(userFields))
	user = s.UserRepo.GetUserByUUID(userFields.UUID)
	return
}
