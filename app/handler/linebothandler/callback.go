package linebothandler

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
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
				switch message.Text {
				// 記帳起始步驟
				case global.RecordBalanceZhTw:
					dateTemplate := h.LineBotService.ShowBalanceDateOptionTemplate()
					if _, err = bot.ReplyMessage(event.ReplyToken, dateTemplate).Do(); err != nil {
						log.Fatalf(err.Error())
					}
				// 選擇的日期為今日
				case global.TodayZhTw:
					h.chooseDateAndShowBalanceType(user.ID, helper.GetNowDate(), event.ReplyToken)
				// 選擇收入/支出
				case global.IncomeZhTw, global.ExpenseZhTw:
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceType,
						Value:  zhTwNameKeyConvert(message.Text),
					})
					itemTemplate := h.LineBotService.ShowBalanceItemOptionTemplate()
					if _, err = bot.ReplyMessage(event.ReplyToken, itemTemplate).Do(); err != nil {
						log.Fatalf(err.Error())
					}
				// 選擇消費種類
				case global.ConsumeGoodsZhTw, global.FruitZhTw, global.WaterBillZhTw, global.OilFeeZhTw, global.BreakfastZhTw, global.LunchZhTw, global.DinnerZhTw, global.RepairRewardZhTw, global.GasFeeZhTw,
					global.InsuranceZhTw, global.LivingExpensesZhTw, global.OrganicFoodZhTw, global.DressFeeZhTw, global.HealthyFoodZhTw, global.AutomaticDeductionZhTw, global.ElectricBillZhTw, global.FishZhTW,
					global.MedicalZhTw, global.TicketZhTw, global.GardeningZhTw, global.GroceryShoppingZhTw, global.EasyCardZhTw, global.ManagementCostZhTw, global.PayBillZhTw, global.PottedPlantZhTw:
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceItem,
						Value:  zhTwNameKeyConvert(message.Text),
					})
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

func (h *Handler) chooseDateAndShowBalanceType(userID int, recordDate, replyToken string) {
	// 選日期時開始暫存收支紀錄
	h.TmpBalanceService.CreateTemporaryBalance(structs.CreateTmpBalanceFields{
		Date:   recordDate,
		UserID: userID,
	})
	// 選擇日期的下一步:選擇收入/支出
	typeOption := h.LineBotService.ShowBalanceTypeOptionTemplate()

	_, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(recordDate), typeOption).Do()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func zhTwNameKeyConvert(str string) string {
	switch str {
	case global.IncomeZhTw:
		return global.Income
	case global.ExpenseZhTw:
		return global.Expense
	case global.ConsumeGoodsZhTw:
		return global.ConsumeGoods
	case global.FruitZhTw:
		return global.Fruit
	case global.WaterBillZhTw:
		return global.WaterBill
	case global.OilFeeZhTw:
		return global.OilFee
	case global.BreakfastZhTw:
		return global.Breakfast
	case global.LunchZhTw:
		return global.Lunch
	case global.DinnerZhTw:
		return global.Dinner
	case global.RepairRewardZhTw:
		return global.RepairReward
	case global.GasFeeZhTw:
		return global.GasFee
	case global.InsuranceZhTw:
		return global.Insurance
	case global.LivingExpensesZhTw:
		return global.LivingExpenses
	case global.OrganicFoodZhTw:
		return global.OrganicFood
	case global.DressFeeZhTw:
		return global.DressFee
	case global.HealthyFoodZhTw:
		return global.HealthyFood
	case global.AutomaticDeductionZhTw:
		return global.AutomaticDeduction
	case global.ElectricBillZhTw:
		return global.ElectricBill
	case global.FishZhTW:
		return global.Fish
	case global.MedicalZhTw:
		return global.Medical
	case global.TicketZhTw:
		return global.Ticket
	case global.GardeningZhTw:
		return global.Gardening
	case global.GroceryShoppingZhTw:
		return global.GroceryShopping
	case global.EasyCardZhTw:
		return global.EasyCard
	case global.ManagementCostZhTw:
		return global.ManagementCost
	case global.PayBillZhTw:
		return global.PayBill
	case global.PottedPlantZhTw:
		return global.PottedPlant
	}
	return ""
}
