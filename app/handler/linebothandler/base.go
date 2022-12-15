package linebothandler

import (
	"ever-book/app/service/dailybalanceservice"
	"ever-book/app/service/linebotservice"
	"ever-book/app/service/tmpbalanceservice"
	"ever-book/app/service/userservice"
)

type Handler struct {
	LineBotService      linebotservice.Interface
	UserService         userservice.Interface
	TmpBalanceService   tmpbalanceservice.Interface
	DailyBalanceService dailybalanceservice.Interface
}

func New() *Handler {
	return &Handler{
		LineBotService:      linebotservice.New(),
		UserService:         userservice.New(),
		TmpBalanceService:   tmpbalanceservice.New(),
		DailyBalanceService: dailybalanceservice.New(),
	}
}
