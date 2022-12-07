package router

import (
	"ever-book/app/handler/lineboth"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	lineBotHandler := lineboth.New()
	r.POST("/callback", lineBotHandler.LineBotCallBack)
}
