package lineboth

import "ever-book/app/service/linebots"

type Handler struct {
	LineBotService linebots.Interface
}

func New() *Handler {
	return &Handler{
		LineBotService: linebots.NewService(),
	}
}
