package userservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/model"
	"ever-book/app/repository/userrepo"
	"sync"
)

type Interface interface {
	GetUserList(userID int) (userList []structs.UserObj)
	GetUserInfo(uuid string) (user model.User)
	CreateUser(userFields structs.UserFields)
	UpdateUser(fields structs.UpdateUserFields)
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

func (s *service) GetUserList(userID int) (userList []structs.UserObj) {
	users := s.UserRepo.GetUserList()
	for _, user := range users {
		if user.ID != userID {
			userList = append(userList, structs.UserObj{
				ID:   user.ID,
				Name: user.Name,
			})
		}
	}
	return
}

func (s *service) GetUserInfo(uuid string) (user model.User) {
	user = s.UserRepo.GetUserByUUID(uuid)
	return
}

func (s *service) CreateUser(userFields structs.UserFields) {
	s.UserRepo.CreateUserByMap(helper.StructToMap(userFields))
}

func (s *service) UpdateUser(fields structs.UpdateUserFields) {
	s.UserRepo.UpdateUserByID(fields.ID, fields.Column, fields.Value)
}
