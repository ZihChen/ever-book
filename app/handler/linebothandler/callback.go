package linebothandler

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"strconv"
)

var bot *linebot.Client

func (h *Handler) LineBotCallBack(ctx *gin.Context) {
	bot = h.LineBotService.GetClient()
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, event := range events {
		user := h.UserService.GetOrCreateUser(event.Source.UserID)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// 紀錄金額
				balanceAmount, err := strconv.Atoi(message.Text)
				if err == nil {
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceAmount,
						Value:  balanceAmount,
					})
					paymentTemplate := h.LineBotService.ShowBalancePaymentOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, paymentTemplate)
				}
				switch message.Text {
				// 記帳起始步驟
				case global.RecordBalanceZhTw:
					tmpRecord, exist := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
					// 如果此用戶已經存在一筆暫存紀錄，則詢問是否繼續步驟
					if exist {
						tmpBalanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(tmpRecord)
						isContinueTemplate := h.LineBotService.ShowContinueOrDiscardOptionTemplate()
						h.replyMessageToUser(event.ReplyToken, tmpBalanceFlexMsg, isContinueTemplate)
					} else {
						dateTemplate := h.LineBotService.ShowBalanceDateOptionTemplate()
						h.replyMessageToUser(event.ReplyToken, dateTemplate)
					}
				case "查看當日統計":
				case "查看當月統計":
				case "刪除上一筆資料":
				case "對方當月統計":
				case "範例教學":
				default:
				}
			}
		case linebot.EventTypePostback:
			data := event.Postback.Data
			switch data {
			// 選擇記帳日期
			case global.Date:
				h.chooseDateAndShowBalanceType(user.ID, event.Postback.Params.Date, event.ReplyToken)
			// 選擇的日期為今日
			case global.Today:
				h.chooseDateAndShowBalanceType(user.ID, helper.GetNowDate(), event.ReplyToken)
			// 選擇收入/支出
			case global.Expense, global.Income:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalanceType,
					Value:  data,
				})
				itemTemplate := h.LineBotService.ShowBalanceItemOptionTemplate()
				h.replyMessageToUser(event.ReplyToken, itemTemplate)
			// 選擇消費種類
			case global.ConsumeGoods, global.Fruit, global.WaterBill, global.OilFee, global.Breakfast, global.Lunch, global.Dinner, global.RepairReward, global.GasFee,
				global.Insurance, global.LivingExpenses, global.OrganicFood, global.DressFee, global.HealthyFood, global.AutomaticDeduction, global.ElectricBill, global.Fish,
				global.Medical, global.Ticket, global.Gardening, global.GroceryShopping, global.EasyCard, global.ManagementCost, global.PayBill, global.PottedPlant:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalanceItem,
					Value:  data,
				})
				amountTemplate := linebot.NewTextMessage("請輸入金額：")
				h.replyMessageToUser(event.ReplyToken, amountTemplate)
			// 選擇付款方式
			case global.Cash, global.CreditCard:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalancePayment,
					Value:  data,
				})
				remarkTemplate := h.LineBotService.ShowBalanceRemarkOptionTemplate()
				h.replyMessageToUser(event.ReplyToken, remarkTemplate)
			// 是否繼續步驟:繼續
			case global.Continue:
				// 依照順序先檢查種類、項目、金額、付款方式、備註
				tmpRecord, _ := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				var template linebot.SendingMessage
				if tmpRecord.Type == "" {
					template = h.LineBotService.ShowBalanceTypeOptionTemplate()
				} else if tmpRecord.Item == "" {
					template = h.LineBotService.ShowBalanceItemOptionTemplate()
				} else if tmpRecord.Amount == 0 {
					template = linebot.NewTextMessage("請輸入金額：")
				} else if tmpRecord.Payment == "" {
					template = h.LineBotService.ShowBalancePaymentOptionTemplate()
				} else if tmpRecord.Remark == "" {

				}
				h.replyMessageToUser(event.ReplyToken, template)
			// 是否繼續步驟:捨棄
			case global.Discard:
				// TODO: 刪除暫存紀錄回到選擇日期步驟
			}
		}
	}
}

func (h *Handler) replyMessageToUser(replyToken string, messages ...linebot.SendingMessage) {
	_, err := bot.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		log.Fatalf("Reply Message Error:%v", err.Error())
	}
}

func (h *Handler) chooseDateAndShowBalanceType(userID int, recordDate, replyToken string) {
	// 選日期時開始暫存收支紀錄
	h.TmpBalanceService.CreateTemporaryBalance(structs.CreateTmpBalanceFields{
		Date:   recordDate,
		UserID: userID,
	})
	// 選擇日期的下一步:選擇收入/支出
	typeOption := h.LineBotService.ShowBalanceTypeOptionTemplate()

	h.replyMessageToUser(replyToken, linebot.NewTextMessage(recordDate), typeOption)
}
