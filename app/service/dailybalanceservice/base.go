package dailybalanceservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"ever-book/app/repository/dailybalancerepo"
	"ever-book/app/repository/userrepo"
	"fmt"
	"sync"
)

type Interface interface {
	CreateDailyBalanceByTmpBalance(userID int, tmpBalanceObj structs.TmpBalanceObj)
	GetLatestDailyBalance(userID int) (tmpBalanceObj structs.TmpBalanceObj, exist bool)
	DeletePreviousDailyBalance(userID int)
	GetDailyBalancesByMonth(userID int, month int) (balanceObjs []structs.BalanceObj)
	GetTotalBalanceByMonth(userID int, month int) (balanceSum structs.BalanceSummaryObj)
}
type service struct {
	DailyBalanceRepo dailybalancerepo.Interface
	UserRepo         userrepo.Interface
}

var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{
			DailyBalanceRepo: dailybalancerepo.New(),
			UserRepo:         userrepo.New(),
		}
	})
	return singleton
}

func (s *service) CreateDailyBalanceByTmpBalance(userID int, tmpBalanceObj structs.TmpBalanceObj) {
	tpMap := helper.StructToMap(tmpBalanceObj)
	tpMap["user_id"] = userID
	s.DailyBalanceRepo.CreateDailyBalanceByMap(tpMap)
}

func (s *service) GetLatestDailyBalance(userID int) (tmpBalanceObj structs.TmpBalanceObj, exist bool) {
	if exist = s.DailyBalanceRepo.CheckDailyBalanceExistByUserID(userID); !exist {
		return
	}
	dailyBalance := s.DailyBalanceRepo.GatLatestDailyBalanceByUserID(userID)
	tmpBalanceObj.Date = func() string {
		dateTime := dailyBalance.Date
		return fmt.Sprintf("%02d-%02d-%02d", dateTime.Year(), dateTime.Month(), dateTime.Day())
	}()
	tmpBalanceObj.Type = dailyBalance.Type
	tmpBalanceObj.Item = dailyBalance.Item
	tmpBalanceObj.Amount = dailyBalance.Amount
	tmpBalanceObj.Payment = dailyBalance.Payment
	tmpBalanceObj.Remark = dailyBalance.Remark
	return
}

func (s *service) DeletePreviousDailyBalance(userID int) {
	dailyBalance := s.DailyBalanceRepo.GatLatestDailyBalanceByUserID(userID)
	s.DailyBalanceRepo.DeleteDailyBalanceByID(dailyBalance.ID)
}

func (s *service) GetDailyBalancesByMonth(userID int, month int) (balanceObjs []structs.BalanceObj) {
	dailyBalances := s.DailyBalanceRepo.GetDailyBalanceByDateInterval(userID, helper.GetIntervalDate(month))
	for _, dailyBalance := range dailyBalances {
		balanceObjs = append(balanceObjs, structs.BalanceObj{
			ID: dailyBalance.ID,
			Date: func() string {
				dateTime := dailyBalance.Date
				return fmt.Sprintf("%02d", dateTime.Day())
			}(),
			Type:    dailyBalance.Type,
			Item:    dailyBalance.Item,
			Amount:  dailyBalance.Amount,
			Payment: dailyBalance.Payment,
			Remark:  dailyBalance.Remark,
		})
	}
	return
}

func (s *service) GetTotalBalanceByMonth(userID int, month int) (balanceSum structs.BalanceSummaryObj) {
	user := s.UserRepo.GetUserInfo(structs.UserParams{
		ID: userID,
	})
	totalExpense := s.DailyBalanceRepo.GetTotalAmountByDateInterval(userID, global.Expense, helper.GetIntervalDate(month))
	totalIncome := s.DailyBalanceRepo.GetTotalAmountByDateInterval(userID, global.Income, helper.GetIntervalDate(month))

	return structs.BalanceSummaryObj{
		TotalExpense: totalExpense,
		TotalBalance: user.Budget + totalIncome - totalExpense,
	}
}
