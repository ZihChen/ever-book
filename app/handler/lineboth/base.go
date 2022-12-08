package lineboth

import (
	"ever-book/app/service/linebots"
	"ever-book/app/service/users"
)

type Handler struct {
	LineBotService linebots.Interface
	UserService    users.Interface
}

func New() *Handler {
	return &Handler{
		LineBotService: linebots.NewService(),
		UserService:    users.NewService(),
	}
}
