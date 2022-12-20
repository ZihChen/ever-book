package linebothandler

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"strconv"
	"strings"
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
					return
				}
				switch message.Text {
				// 記帳起始步驟
				case global.RecordBalanceZhTw:
					tmpRecord, exist := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
					// 如果此用戶已經存在一筆暫存紀錄，則詢問是否繼續步驟
					if exist {
						tmpBalanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(global.UnfinishedBalanceZhTw, tmpRecord)
						isContinueTemplate := h.LineBotService.ShowContinueOrDiscardOptionTemplate()
						h.replyMessageToUser(event.ReplyToken, tmpBalanceFlexMsg, isContinueTemplate)
					} else {
						dateTemplate := h.LineBotService.ShowBalanceDateOptionTemplate()
						h.replyMessageToUser(event.ReplyToken, dateTemplate)
					}
				// 查看帳本
				case global.AccountBookSummaryZhTw:
					monthOption := h.LineBotService.ShowMonthOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, monthOption)
				case global.OtherAccountBookZhTw:
				case global.DeletePreviousRecordZhTw:
					balance, exist := h.DailyBalanceService.GetLatestDailyBalance(user.ID)
					if !exist {
						// TODO: 回傳沒有資料可以刪除
					}
					cancelTemplate := h.LineBotService.ShowCancelBalanceFlexMessage("您要刪除下列資料:", balance)
					cancelOption := h.LineBotService.ShowCancelOrNotOptionTemplate()

					h.replyMessageToUser(event.ReplyToken, cancelTemplate, cancelOption)

				case global.TodayZhTw, global.IncomeZhTw, global.ExpenseZhTw, global.CashZhTw, global.CreditCardZhTw, global.ConsumeGoodsZhTw, global.FruitZhTw, global.WaterBillZhTw, global.OilFeeZhTw,
					global.BreakfastZhTw, global.LunchZhTw, global.DinnerZhTw, global.RepairRewardZhTw, global.GasFeeZhTw, global.InsuranceZhTw, global.LivingExpensesZhTw, global.OrganicFoodZhTw, global.DressFeeZhTw,
					global.HealthyFoodZhTw, global.AutomaticDeductionZhTw, global.ElectricBillZhTw, global.FishZhTW, global.MedicalZhTw, global.TicketZhTw, global.GardeningZhTw, global.GroceryShoppingZhTw,
					global.EasyCardZhTw, global.ManagementCostZhTw, global.PayBillZhTw, global.PottedPlantZhTw, global.ContinueZhTw, global.DiscardZhTw, global.NeedRemarkZhTw, global.SkipRemarkZhTw, global.ConfirmZhTw, global.CancelZhTw,
					global.JanZhTw, global.FebZhTw, global.MarZhTw, global.AprZhTw, global.MayZhTw, global.JunZhTw, global.JulZhTw, global.AugZhTw, global.SepZhTw, global.OctZhTw, global.NovZhTw, global.DecZhTw:
					// 避免觸發寫入備註
					return
				default:
					tmpRecord, exist := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
					if !exist {
						// 如果還沒有暫存收支紀錄亂打資料就不動作
						return
					}
					h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
						UserID: user.ID,
						Column: global.TemporaryBalanceRemark,
						Value:  message.Text,
					})
					// 檢查是否除了備註外其他必填欄位都已經填寫好
					template := h.checkColumnsIsFilled(tmpRecord)
					if template != nil {
						h.replyMessageToUser(event.ReplyToken, template)
						return
					}
					// 新增一筆收支紀錄
					h.DailyBalanceService.CreateDailyBalanceByTmpBalance(user.ID, tmpRecord)
					// 將暫存紀錄刪除
					h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
					balanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(global.SuccessRecordZhTw, tmpRecord)
					h.replyMessageToUser(event.ReplyToken, balanceFlexMsg)
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
				amountTemplate := linebot.NewTextMessage(global.EnterAmountZhTw)
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
			// 填寫備註
			case global.NeedRemark:
				template := linebot.NewTextMessage(global.EnterRemarkZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
			case global.SkipRemark:
				tmpRecord, _ := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				// 新增一筆收支紀錄
				h.DailyBalanceService.CreateDailyBalanceByTmpBalance(user.ID, tmpRecord)
				// 將暫存紀錄刪除
				h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
				balanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(global.SuccessRecordZhTw, tmpRecord)
				h.replyMessageToUser(event.ReplyToken, balanceFlexMsg)
			// 是否繼續步驟:繼續
			case global.Continue:
				// 依照順序先檢查種類、項目、金額、付款方式、備註
				tmpRecord, _ := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				template := h.checkColumnsIsFilled(tmpRecord)
				h.replyMessageToUser(event.ReplyToken, template)
			// 是否繼續步驟:捨棄
			case global.Discard:
				// 刪除暫存紀錄回到選擇日期步驟
				h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
				template := linebot.NewTextMessage(global.SuccessDiscardZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
			// 確認刪除上一筆資料
			case global.ConfirmDelete:
				h.DailyBalanceService.DeletePreviousDailyBalance(user.ID)
				template := linebot.NewTextMessage(global.SuccessDeleteZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
			case global.JanSummary, global.FebSummary, global.MarSummary, global.AprSummary, global.MaySummary, global.JunSummary, global.JulSummary, global.AugSummary, global.SepSummary, global.OctSummary, global.NovSummary, global.DecSummary:
				month := helper.KeyNameConvertToMonth(data)
				balanceObjs := h.DailyBalanceService.GetDailyBalancesByMonth(user.ID, month)
				balanceSum := h.DailyBalanceService.GetTotalBalanceByMonth(user.ID, month)
				dateFormat := helper.GetIntervalDateFormat(month)
				template := h.LineBotService.ShowSummaryFlexMessage(func() string {
					var msgTitle strings.Builder
					msgTitle.WriteString(strconv.Itoa(month))
					msgTitle.WriteString("月總結")
					return msgTitle.String()
				}(), structs.SummaryFlexMsg{
					StartDate:     dateFormat.StartDate,
					EndDate:       dateFormat.EndDate,
					BalanceObjs:   balanceObjs,
					TotalExpenses: balanceSum.TotalExpense,
					TotalBalance:  balanceSum.TotalBalance,
				})
				h.replyMessageToUser(event.ReplyToken, template)
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

func (h *Handler) checkColumnsIsFilled(tmpRecord structs.TmpBalanceObj) (template linebot.SendingMessage) {
	if tmpRecord.Type == "" {
		template = h.LineBotService.ShowBalanceTypeOptionTemplate()
	} else if tmpRecord.Item == "" {
		template = h.LineBotService.ShowBalanceItemOptionTemplate()
	} else if tmpRecord.Amount == 0 {
		template = linebot.NewTextMessage(global.EnterAmountZhTw)
	} else if tmpRecord.Payment == "" {
		template = h.LineBotService.ShowBalancePaymentOptionTemplate()
	} else if tmpRecord.Remark == "" {
		template = linebot.NewTextMessage(global.EnterRemarkZhTw)
	}
	return
}
