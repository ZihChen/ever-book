package structs

import (
	"ever-book/app/global"
)

type CreateTmpBalanceFields struct {
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	Item    string `json:"item"`
	Amount  int    `json:"amount"`
	Payment string `json:"payment"`
	Remark  string `json:"remark"`
}

type UpdateTmpBalanceFields struct {
	UserID int                           `json:"user_id"`
	Column global.TemporaryBalanceColumn `json:"column"`
	Value  interface{}                   `json:"value"`
}

type TmpBalanceObj struct {
	Date    string `json:"date"`
	Type    string `json:"type"`
	Item    string `json:"item"`
	Amount  int    `json:"amount"`
	Payment string `json:"payment"`
	Remark  string `json:"remark"`
}

type BalanceObj struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	Item    string `json:"item"`
	Amount  int    `json:"amount"`
	Payment string `json:"payment"`
	Remark  string `json:"remark"`
}

type DateInterval struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type SummaryFlexMsg struct {
	StartDate     string       `json:"start_date"`
	EndDate       string       `json:"end_date"`
	TotalExpenses int          `json:"total_expenses"`
	TotalBalance  int          `json:"total_balance"`
	BalanceObjs   []BalanceObj `json:"balance_objs"`
}

type BalanceSummaryObj struct {
	TotalExpense int `json:"total_expenses"`
	TotalBalance int `json:"total_balance"`
}
