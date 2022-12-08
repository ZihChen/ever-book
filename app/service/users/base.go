package users

import (
	"ever-book/app/model"
	"ever-book/app/repository/userr"
	"sync"
)

type Interface interface {
	GetOrCreateUser(uuid string) (user model.User)
}

type service struct {
	UserRepo userr.Interface
}

var singleton *service
var once sync.Once

func NewService() Interface {
	once.Do(func() {
		singleton = &service{
			UserRepo: userr.NewRepository(),
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