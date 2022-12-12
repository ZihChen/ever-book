package linebothandler

import (
	"ever-book/app/global/structs"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
)

func (h *Handler) LineBotCallBack(ctx *gin.Context) {

	bot := h.LineBotService.GetClient()
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, event := range events {
		user := h.UserService.GetOrCreateUser(event.Source.UserID)
		if event.Type == linebot.EventTypePostback {
			recordDate := event.Postback.Params.Date

			// 選日期時開始暫存收支紀錄
			h.TmpBalanceService.CreateTemporaryBalance(structs.CreateTmpBalanceFields{
				Date:   recordDate,
				UserID: user.ID,
			})
			// 選鑿日期的下一步:選擇收入/支出
			typeOption := h.LineBotService.ShowBalanceTypeOptionTemplate()

			_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(recordDate), typeOption).Do()
			if err != nil {
				log.Fatalf(err.Error())
			}
		}

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "我要記帳":
					dateTemplate := h.LineBotService.ShowBalanceDateOptionTemplate()
					if _, err = bot.ReplyMessage(event.ReplyToken, dateTemplate).Do(); err != nil {
						log.Fatalf(err.Error())
					}
				case "收入":
					fallthrough
				case "支出":
					// TODO: 先找到最新的一筆TmpRecord，再對type做寫入

				case "查看當日統計":
				case "查看當月統計":
				case "刪除上一筆資料":
				case "對方當月統計":
				case "範例教學":
				default:

				}
			}
		}
	}
}
