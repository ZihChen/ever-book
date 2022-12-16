package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (s *service) ShowBalanceDateOptionTemplate() *linebot.TemplateMessage {
	nowDate := helper.GetNowDate()
	return linebot.NewTemplateMessage(global.BalanceDateOptionZhTw, &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: global.BalanceDateOptionZhTw,
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
		global.BalanceTypeOptionZhTw,
		linebot.NewConfirmTemplate(
			global.BalanceTypeOptionZhTw,
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
			actions = append(actions, &linebot.PostbackAction{
				Label: item,
				Text:  item,
				Data:  helper.ZhTwConvertToKeyName(item),
			})
		}
		itemGroup := &linebot.CarouselColumn{
			Text:    name,
			Actions: actions,
		}
		carouselColumn = append(carouselColumn, itemGroup)
	}

	return linebot.NewTemplateMessage(global.BalanceItemOptionZhTw, &linebot.CarouselTemplate{
		Columns: carouselColumn,
	})
}

func (s *service) ShowBalancePaymentOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(global.BalancePaymentOptionZhTw, &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: global.BalancePaymentOptionZhTw,
				Actions: []linebot.TemplateAction{
					&linebot.PostbackAction{
						Label: global.CashZhTw,
						Text:  global.CashZhTw,
						Data:  global.Cash,
					},
					&linebot.PostbackAction{
						Label: global.CreditCardZhTw,
						Text:  global.CreditCardZhTw,
						Data:  global.CreditCard,
					},
				},
			},
		},
	})
}

func (s *service) ShowBalanceRemarkOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(
		global.BalanceRemarkOptionZhTw,
		linebot.NewConfirmTemplate(
			global.BalanceRemarkOptionZhTw,
			&linebot.PostbackAction{
				Label: global.NeedRemarkZhTw,
				Text:  global.NeedRemarkZhTw,
				Data:  global.NeedRemark,
			},
			&linebot.PostbackAction{
				Label: global.SkipRemarkZhTw,
				Text:  global.SkipRemarkZhTw,
				Data:  global.SkipRemark,
			},
		),
	)
}

func (s *service) ShowContinueOrDiscardOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(
		global.ContinueOrNotOptionZhTw,
		linebot.NewConfirmTemplate(
			global.ContinueOrNotOptionZhTw,
			&linebot.PostbackAction{
				Label: global.ContinueZhTw,
				Text:  global.ContinueZhTw,
				Data:  global.Continue,
			},
			&linebot.PostbackAction{
				Label: global.DiscardZhTw,
				Text:  global.DiscardZhTw,
				Data:  global.Discard,
			},
		),
	)
}

func (s *service) ShowCancelOrNotOptionTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(
		global.CancelOrNotZhTw,
		linebot.NewConfirmTemplate(
			global.CancelOrNotZhTw,
			&linebot.PostbackAction{
				Label: global.ConfirmZhTw,
				Text:  global.ConfirmZhTw,
				Data:  global.ConfirmDelete,
			},
			&linebot.PostbackAction{
				Label: global.CancelZhTw,
				Text:  global.CancelZhTw,
				Data:  global.CancelDelete,
			},
		),
	)
}
