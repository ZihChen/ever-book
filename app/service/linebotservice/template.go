package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"strconv"
	"strings"
	"time"
)

func (s *service) ShowBalanceDateOptionTemplate() *linebot.TemplateMessage {
	nowDate := helper.GetNowDate()
	return linebot.NewTemplateMessage(global.BalanceDateOptionZhTw, &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: global.BalanceDateOptionZhTw,
				Actions: []linebot.TemplateAction{
					&linebot.DatetimePickerAction{
						Label:   global.ClickZhTw,
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
		str := strings.Split(name, "-")
		itemGroup := &linebot.CarouselColumn{
			Title:   str[0],
			Text:    str[1],
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

func (s *service) ShowMonthOptionTemplate() (template *linebot.TemplateMessage) {
	var carouselColumn []*linebot.CarouselColumn
	var monthsArr []int
	nowMonth := time.Now().Month()
	monthsArr = append(monthsArr, int(nowMonth))
	for i := 1; i <= 2; i++ {
		month := time.Now().AddDate(0, -i, 0).Month()
		monthsArr = append(monthsArr, int(month))
	}
	var actions []linebot.TemplateAction
	for _, month := range monthsArr {
		actions = append(actions, &linebot.PostbackAction{
			Label: helper.MonthConvertToZhTw(month),
			Text:  helper.MonthConvertToZhTw(month),
			Data:  helper.MonthConvertToKeyName(month),
		})
	}
	carouselColumn = append(carouselColumn, &linebot.CarouselColumn{
		Text:    global.BalanceMonthOptionZhTw,
		Actions: actions,
	})
	return linebot.NewTemplateMessage(global.BalanceMonthOptionZhTw, &linebot.CarouselTemplate{
		Columns: carouselColumn,
	})
}

func (s *service) ShowIsNeedDateAndRemarkOptionTemplate() *linebot.TemplateMessage {

	return linebot.NewTemplateMessage(global.IsNeedDateAndRemarkOptionZhTw, &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: global.IsNeedDateAndRemarkOptionZhTw,
				Actions: []linebot.TemplateAction{
					&linebot.PostbackAction{
						Label: global.TypeDateZhTw,
						Text:  global.TypeDateZhTw,
						Data:  global.TypeDate,
					},
					&linebot.PostbackAction{
						Label: global.TypeRemarkZhTw,
						Text:  global.TypeRemarkZhTw,
						Data:  global.TypeRemark,
					},
					&linebot.PostbackAction{
						Label: global.SkipZhTw,
						Text:  global.SkipZhTw,
						Data:  global.Skip,
					},
				},
			},
		},
	})
}

func (s *service) ShowUserGroupOption() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage(global.FamilyBalanceOptionZhTw, &linebot.CarouselTemplate{
		Columns: []*linebot.CarouselColumn{
			{
				Text: global.FamilyBalanceOptionZhTw,
				Actions: []linebot.TemplateAction{
					&linebot.PostbackAction{
						Label: global.CheckOtherBalanceZhTw,
						Text:  global.CheckOtherBalanceZhTw,
						Data:  global.CheckOtherBalance,
					},
					&linebot.PostbackAction{
						Label: global.BindOtherBalanceZhTw,
						Text:  global.BindOtherBalanceZhTw,
						Data:  global.BindOtherBalance,
					},
				},
			},
		},
	})
}

func (s *service) ShowUserListOption(userList []structs.UserObj) *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("選擇綁定成員", &linebot.ButtonsTemplate{
		Title: "選擇綁定成員",
		Text:  "如下:",
		Actions: func() (templateActions []linebot.TemplateAction) {
			for _, user := range userList {
				templateActions = append(templateActions, &linebot.PostbackAction{
					Label: user.Name,
					Text:  fmt.Sprintf("綁定：%s-%s", user.Name, strconv.Itoa(user.ID)),
					Data:  fmt.Sprintf("user-bind-%s", strconv.Itoa(user.ID)),
				})
			}
			return
		}(),
	})
}
