package linebotservice

import (
	"ever-book/app/global"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"time"
)

func (s *service) ShowBalanceDateOptionTemplate() *linebot.TemplateMessage {
	tz, _ := time.LoadLocation("Asia/Taipei")
	now := time.Now().UTC().In(tz)
	nowDate := fmt.Sprintf("%02d-%02d-%02d", now.Year(), now.Month(), now.Day())

	return linebot.NewTemplateMessage("選擇支出/收入日期", &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: "選擇支出/收入日期",
				Actions: []linebot.TemplateAction{
					&linebot.MessageAction{
						Label: "今日",
						Text:  "今日",
					},
					&linebot.DatetimePickerAction{
						Label:   "日期",
						Mode:    "date",
						Data:    "data",
						Initial: nowDate,
						Max:     nowDate,
					},
				},
			},
		},
	})
}

func (s *service) ShowBalanceTypeOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(
		"請選擇收入或是支出",
		linebot.NewConfirmTemplate(
			"請選擇收入或是支出",
			linebot.NewMessageAction("income", global.Income),
			linebot.NewMessageAction("expense", global.Expense),
		),
	)
}
