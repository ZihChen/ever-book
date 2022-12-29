package usermemberservice

import (
	"ever-book/app/model"
	"ever-book/app/repository/usermemberrepo"
	"ever-book/app/repository/userrepo"
	"sync"
)

type Interface interface {
	BindUserMember(user model.User, memberId int)
}

type service struct {
	UserMapRepo usermemberrepo.Interface
	UserRepo    userrepo.Interface
}

var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{
			UserMapRepo: usermemberrepo.New(),
			UserRepo:    userrepo.New(),
		}
	})
	return singleton
}

func (s *service) BindUserMember(user model.User, memberId int) {
	var members []model.User
	member := s.UserRepo.GetUserByID(memberId)
	members = append(user.Members, member)
	s.UserMapRepo.CreateUserMembersRelation(user, members)
}
