package tmpbalanceservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/repository/tmpbalancerepo"
	"fmt"
	"sync"
)

type Interface interface {
	GetTemporaryBalanceByUserID(userID int) (tmpBalanceObj structs.TmpBalanceObj, exist bool)
	CreateTemporaryBalance(fields structs.CreateTmpBalanceFields)
	UpdateTemporaryBalance(fields structs.UpdateTmpBalanceFields)
	DeleteTemporaryBalance(userID int)
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

func (s *service) GetTemporaryBalanceByUserID(userID int) (tmpBalanceObj structs.TmpBalanceObj, exist bool) {
	if exist = s.TmpBalanceRepo.CheckTemporaryBalanceExistByUserID(userID); exist == false {
		return
	}
	tmpBalance := s.TmpBalanceRepo.GetTemporaryBalanceByUserID(userID)
	tmpBalanceObj.Date = func() string {
		dateTime := tmpBalance.Date
		return fmt.Sprintf("%02d-%02d-%02d", dateTime.Year(), dateTime.Month(), dateTime.Day())
	}()
	tmpBalanceObj.Type = tmpBalance.Type
	tmpBalanceObj.Item = tmpBalance.Item
	tmpBalanceObj.Amount = tmpBalance.Amount
	tmpBalanceObj.Payment = tmpBalance.Payment
	tmpBalanceObj.Remark = tmpBalance.Remark
	return
}

func (s *service) DeleteTemporaryBalance(userID int) {
	s.TmpBalanceRepo.DeleteTemporaryBalanceByUserID(userID)
}
