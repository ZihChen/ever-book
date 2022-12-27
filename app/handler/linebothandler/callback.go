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
		sourceID := event.Source.UserID
		user := h.UserService.GetOrCreateUser(structs.UserFields{
			UUID: sourceID,
			Name: func() string {
				profile, _ := bot.GetProfile(sourceID).Do()
				return profile.DisplayName
			}(),
		})

		switch event.Type {
		// Message Type
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// 紀錄金額
				balanceAmount, err := strconv.Atoi(message.Text)
				if err == nil {
					h.TmpBalanceService.CreateTemporaryBalance(structs.CreateTmpBalanceFields{
						UserID: user.ID,
						Amount: balanceAmount,
						Date:   helper.GetNowDate(),
					})
					template := h.LineBotService.ShowBalanceItemOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, template)
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
						amountTemplate := linebot.NewTextMessage(global.EnterAmountZhTw)
						h.replyMessageToUser(event.ReplyToken, amountTemplate)
					}
					return
				//查看帳本
				case global.AccountBookSummaryZhTw:
					monthOption := h.LineBotService.ShowMonthOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, monthOption)
					return
				case global.TodayZhTw, global.IncomeZhTw, global.ExpenseZhTw, global.CashZhTw, global.CreditCardZhTw, global.ConsumeGoodsZhTw, global.FruitZhTw, global.WaterBillZhTw, global.OilFeeZhTw,
					global.BreakfastZhTw, global.LunchZhTw, global.DinnerZhTw, global.RepairRewardZhTw, global.GasFeeZhTw, global.InsuranceZhTw, global.LivingExpensesZhTw, global.OrganicFoodZhTw, global.DressFeeZhTw,
					global.HealthyFoodZhTw, global.AutomaticDeductionZhTw, global.ElectricBillZhTw, global.FishZhTW, global.MedicalZhTw, global.TicketZhTw, global.GardeningZhTw, global.GroceryShoppingZhTw,
					global.EasyCardZhTw, global.ManagementCostZhTw, global.PayBillZhTw, global.PottedPlantZhTw, global.ContinueZhTw, global.DiscardZhTw, global.NeedRemarkZhTw, global.SkipRemarkZhTw, global.ConfirmZhTw, global.CancelZhTw,
					global.JanZhTw, global.FebZhTw, global.MarZhTw, global.AprZhTw, global.MayZhTw, global.JunZhTw, global.JulZhTw, global.AugZhTw, global.SepZhTw, global.OctZhTw, global.NovZhTw, global.DecZhTw, global.TypeDateZhTw, global.TypeRemarkZhTw:
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

					if tmpRecord.Date == helper.GetNowDate() {
						template := h.LineBotService.ShowBalanceDateOptionTemplate()
						h.replyMessageToUser(event.ReplyToken, template)
						return
					}

					tmpRecord, _ = h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)

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

		// Postback Type
		case linebot.EventTypePostback:
			data := event.Postback.Data
			switch data {
			// 選擇消費種類
			case global.ConsumeGoods, global.Fruit, global.WaterBill, global.OilFee, global.Breakfast, global.Lunch, global.Dinner, global.RepairReward, global.GasFee,
				global.Insurance, global.LivingExpenses, global.OrganicFood, global.DressFee, global.HealthyFood, global.AutomaticDeduction, global.ElectricBill, global.Fish,
				global.Medical, global.Ticket, global.Gardening, global.GroceryShopping, global.EasyCard, global.ManagementCost, global.PayBill, global.Salary, global.OtherIncome:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalanceItem,
					Value:  data,
				})
				balanceType := func() string {
					switch data {
					case global.Salary, global.OtherIncome, global.RepairReward:
						return global.Income
					default:
						return global.Expense
					}
				}()
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalanceType,
					Value:  balanceType,
				})
				template := h.LineBotService.ShowBalancePaymentOptionTemplate()
				h.replyMessageToUser(event.ReplyToken, template)
				return
			// 選擇付款方式
			case global.Cash, global.CreditCard:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					UserID: user.ID,
					Column: global.TemporaryBalancePayment,
					Value:  data,
				})
				template := h.LineBotService.ShowIsNeedDateAndRemarkOptionTemplate()
				h.replyMessageToUser(event.ReplyToken, template)
				return
			// 填寫指定日期
			case global.TypeDate:
				template := h.LineBotService.ShowBalanceDateOptionTemplate()
				h.replyMessageToUser(event.ReplyToken, template)
				return
			// 選擇記帳日期
			case global.Date:
				h.TmpBalanceService.UpdateTemporaryBalance(structs.UpdateTmpBalanceFields{
					Column: global.TemporaryBalanceDate,
					Value:  event.Postback.Params.Date,
					UserID: user.ID,
				})
				tmpBalance, exist := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				if exist && tmpBalance.Remark == "" {
					template := h.LineBotService.ShowBalanceRemarkOptionTemplate()
					h.replyMessageToUser(event.ReplyToken, template)
					return
				}
				// 新增一筆收支紀錄
				h.DailyBalanceService.CreateDailyBalanceByTmpBalance(user.ID, tmpBalance)
				// 將暫存紀錄刪除
				h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
				balanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(global.SuccessRecordZhTw, tmpBalance)
				h.replyMessageToUser(event.ReplyToken, balanceFlexMsg)
				return
			case global.NeedRemark, global.TypeRemark:
				template := linebot.NewTextMessage(global.EnterRemarkZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
				return
			case global.Skip, global.SkipRemark:
				tmpRecord, _ := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				// 新增一筆收支紀錄
				h.DailyBalanceService.CreateDailyBalanceByTmpBalance(user.ID, tmpRecord)
				// 將暫存紀錄刪除
				h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
				balanceFlexMsg := h.LineBotService.ShowTmpBalanceFlexMessage(global.SuccessRecordZhTw, tmpRecord)
				h.replyMessageToUser(event.ReplyToken, balanceFlexMsg)
				return
			case global.Continue:
				// 依照順序先檢查種類、項目、金額、付款方式、備註
				tmpRecord, _ := h.TmpBalanceService.GetTemporaryBalanceByUserID(user.ID)
				template := h.checkColumnsIsFilled(tmpRecord)
				h.replyMessageToUser(event.ReplyToken, template)
				return
			// 是否繼續步驟:捨棄
			case global.Discard:
				// 刪除暫存紀錄回到選擇日期步驟
				h.TmpBalanceService.DeleteTemporaryBalance(user.ID)
				template := linebot.NewTextMessage(global.SuccessDiscardZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
				return
			// 確認刪除上一筆資料
			case global.ConfirmDelete:
				h.DailyBalanceService.DeletePreviousDailyBalance(user.ID)
				template := linebot.NewTextMessage(global.SuccessDeleteZhTw)
				h.replyMessageToUser(event.ReplyToken, template)
				return
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
				return
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

func (h *Handler) checkColumnsIsFilled(tmpRecord structs.TmpBalanceObj) (template linebot.SendingMessage) {
	if tmpRecord.Item == "" {
		template = h.LineBotService.ShowBalanceItemOptionTemplate()
	} else if tmpRecord.Payment == "" {
		template = h.LineBotService.ShowBalancePaymentOptionTemplate()
	} else if tmpRecord.Remark == "" {
		template = linebot.NewTextMessage(global.EnterRemarkZhTw)
	}
	return
}
