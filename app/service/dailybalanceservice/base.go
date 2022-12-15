package dailybalanceservice

import (
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/repository/dailybalancerepo"
	"sync"
)

type Interface interface {
	CreateDailyBalanceByTmpBalance(userID int, tmpBalanceObj structs.TmpBalanceObj)
}
type service struct {
	DailyBalanceRepo dailybalancerepo.Interface
}

var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{
			DailyBalanceRepo: dailybalancerepo.New(),
		}
	})
	return singleton
}

func (s *service) CreateDailyBalanceByTmpBalance(userID int, tmpBalanceObj structs.TmpBalanceObj) {
	tpMap := helper.StructToMap(tmpBalanceObj)
	tpMap["user_id"] = userID
	s.DailyBalanceRepo.CreateDailyBalanceByMap(tpMap)
}
