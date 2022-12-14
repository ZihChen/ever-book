package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (s *service) ShowBalanceDateOptionTemplate() *linebot.TemplateMessage {
	nowDate := helper.GetNowDate()
	return linebot.NewTemplateMessage("選擇支出/收入日期", &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: "選擇支出/收入日期",
				Actions: []linebot.TemplateAction{
					&linebot.PostbackAction{
						Label: global.TodayZhTw,
						Text:  global.TodayZhTw,
						Data:  global.Today,
					},
					&linebot.DatetimePickerAction{
						Label:   global.DateZhTw,
						Mode:    global.Date,
						Data:    global.Date,
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
			&linebot.PostbackAction{
				Label: global.IncomeZhTw,
				Text:  global.IncomeZhTw,
				Data:  global.Income,
			},
			&linebot.PostbackAction{
				Label: global.ExpenseZhTw,
				Text:  global.ExpenseZhTw,
				Data:  global.Expense,
			},
		),
	)
}

func (s *service) ShowBalanceItemOptionTemplate() (template *linebot.TemplateMessage) {
	var carouselColumn []*linebot.CarouselColumn
	for name, group := range global.BalanceItems {
		var actions []linebot.TemplateAction
		for _, item := range group {
			actions = append(actions, &linebot.MessageAction{
				Label: item,
				Text:  item,
			})
		}
		itemGroup := &linebot.CarouselColumn{
			Text:    name,
			Actions: actions,
		}
		carouselColumn = append(carouselColumn, itemGroup)
	}

	return linebot.NewTemplateMessage("選擇消費項目", &linebot.CarouselTemplate{
		Columns: carouselColumn,
	})
}

func (s *service) ShowBalancePaymentOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("選擇支出/收入日期", &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: "選擇付款方式",
				Actions: []linebot.TemplateAction{
					&linebot.MessageAction{
						Label: global.CashZhTw,
						Text:  global.CashZhTw,
					},
					&linebot.MessageAction{
						Label: global.CreditCardZhTw,
						Text:  global.CreditCardZhTw,
					},
				},
			},
		},
	})
}

func (s *service) ShowBalanceRemarkOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(
		"需不需要填寫備註",
		linebot.NewConfirmTemplate(
			"需不需要填寫備註",
			linebot.NewMessageAction(global.NeedZhTw, global.NeedZhTw),
			linebot.NewMessageAction(global.SkipZhTw, global.SkipZhTw),
		),
	)
}
