package tmpbalanceservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/repository/tmpbalancerepo"
	"sync"
)

type Interface interface {
	CreateTemporaryBalance(fields structs.CreateTmpBalanceFields)
}

type service struct {
	TmpBalanceRepo tmpbalancerepo.Interface
}

var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{
			TmpBalanceRepo: tmpbalancerepo.New(),
		}
	})
	return singleton
}

func (s *service) CreateTemporaryBalance(fields structs.CreateTmpBalanceFields) {
	tpMap := helper.StructToMap(fields)
	s.TmpBalanceRepo.CreateTemporaryBalanceByMap(tpMap)
}
