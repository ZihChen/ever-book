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
		if event.Type == linebot.EventTypePostback {
			h.chooseDateAndShowBalanceType(user.ID, event.Postback.Params.Date, event.ReplyToken)
		}

		if event.Type == linebot.EventTypeMessage {
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
					dateTemplate := h.LineBotService.ShowBalanceDateOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, dateTemplate)
				// 選擇的日期為今日
				case global.TodayZhTw:
					h.chooseDateAndShowBalanceType(user.ID, helper.GetNowDate(), event.ReplyToken)
				// 選擇收入/支出
				case global.IncomeZhTw, global.ExpenseZhTw:
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceType,
						Value:  helper.ZhTwConvertToKeyName(message.Text),
					})
					itemTemplate := h.LineBotService.ShowBalanceItemOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, itemTemplate)
				// 選擇消費種類
				case global.ConsumeGoodsZhTw, global.FruitZhTw, global.WaterBillZhTw, global.OilFeeZhTw, global.BreakfastZhTw, global.LunchZhTw, global.DinnerZhTw, global.RepairRewardZhTw, global.GasFeeZhTw,
					global.InsuranceZhTw, global.LivingExpensesZhTw, global.OrganicFoodZhTw, global.DressFeeZhTw, global.HealthyFoodZhTw, global.AutomaticDeductionZhTw, global.ElectricBillZhTw, global.FishZhTW,
					global.MedicalZhTw, global.TicketZhTw, global.GardeningZhTw, global.GroceryShoppingZhTw, global.EasyCardZhTw, global.ManagementCostZhTw, global.PayBillZhTw, global.PottedPlantZhTw:
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceItem,
						Value:  helper.ZhTwConvertToKeyName(message.Text),
					})
					amountTemplate := linebot.NewTextMessage("請輸入金額：")
					h.replyMessageToUser(event.ReplyToken, amountTemplate)
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

func (h *Handler) replyMessageToUser(replyToken string, messages ...linebot.SendingMessage) {
	_, err := bot.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		log.Fatalf(err.Error())
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
