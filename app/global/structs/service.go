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
