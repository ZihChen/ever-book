package userservice

import (
	"ever-book/app/model"
	"ever-book/app/repository/userrepo"
	"sync"
)

type Interface interface {
	GetOrCreateUser(uuid string) (user model.User)
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

func (s *service) GetOrCreateUser(uuid string) (user model.User) {
	user = s.UserRepo.GetUserByUUID(uuid)
	if user.ID != 0 {
		return
	}

	s.UserRepo.CreateUserByUUID(uuid)
	user = s.UserRepo.GetUserByUUID(uuid)

	return
}
