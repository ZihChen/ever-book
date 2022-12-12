package tmpbalanceservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/repository/tmpbalancerepo"
	"sync"
)

type Interface interface {
	CreateTemporaryBalance(fields structs.CreateTmpBalanceFields)
	UpdateTemporaryBalance(fields structs.UpdateTmpBalanceFields)
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

func (s *service) UpdateTemporaryBalance(fields structs.UpdateTmpBalanceFields) {
	tmpBalance := s.TmpBalanceRepo.GetLatestTemporaryBalanceByUserID(fields.UserID)
	s.TmpBalanceRepo.UpdateTemporaryBalanceByID(tmpBalance.ID, fields.Column, fields.Value)
}
