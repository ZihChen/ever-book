package router

import (
	"ever-book/app/handler/linebothandler"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	lineBotHandler := linebothandler.New()
	r.POST("/callback", lineBotHandler.LineBotCallBack)
}
