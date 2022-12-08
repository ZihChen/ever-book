package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;not null;primary_key"`
	UUID      string    `json:"uuid" gorm:"column:uuid;type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp; default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp; not null;default:current_timestamp ON update current_timestamp"`
}

func (User) TableName() string {
	return "user"
}