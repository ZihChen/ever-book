package model

import "time"

type TemporaryBalance struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;not null;primary_key"`
	UserID    int       `json:"user_id" gorm:"column:user_id;type:int(11);not null"`
	Type      string    `json:"type" gorm:"column:type;type:varchar(20)"`
	Item      string    `json:"item" gorm:"column:item;type:varchar(20)"`
	Payment   string    `json:"payment" gorm:"column:payment;type:varchar(20)"`
	Amount    int       `json:"amount" gorm:"column:amount;type:int(20)"`
	Remark    string    `json:"remark" gorm:"column:remark;type:varchar(50)"`
	Date      time.Time `json:"date" gorm:"column:date;type:datetime;not null;default:current_timestamp"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp; default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp; not null;default:current_timestamp ON update current_timestamp"`
}

func (TemporaryBalance) TableName() string {
	return "temporary_balance"
}
