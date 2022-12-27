package helper

import "ever-book/app/global"

func ZhTwConvertToKeyName(str string) string {
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
	case global.CreditCardZhTw:
		return global.CreditCard
	case global.CashZhTw:
		return global.Cash
	case global.SalaryZhTw:
		return global.Salary
	case global.OtherIncomeZhTw:
		return global.OtherIncome
	default:
		return ""
	}
}

func KeyNameConvertToZhTw(str string) string {
	switch str {
	case global.Income:
		return global.IncomeZhTw
	case global.Expense:
		return global.ExpenseZhTw
	case global.ConsumeGoods:
		return global.ConsumeGoodsZhTw
	case global.Fruit:
		return global.FruitZhTw
	case global.WaterBill:
		return global.WaterBillZhTw
	case global.OilFee:
		return global.OilFeeZhTw
	case global.Breakfast:
		return global.BreakfastZhTw
	case global.Lunch:
		return global.LunchZhTw
	case global.Dinner:
		return global.DinnerZhTw
	case global.RepairReward:
		return global.RepairRewardZhTw
	case global.GasFee:
		return global.GasFeeZhTw
	case global.Insurance:
		return global.InsuranceZhTw
	case global.LivingExpenses:
		return global.LivingExpensesZhTw
	case global.OrganicFood:
		return global.OrganicFoodZhTw
	case global.DressFee:
		return global.DressFeeZhTw
	case global.HealthyFood:
		return global.HealthyFoodZhTw
	case global.AutomaticDeduction:
		return global.AutomaticDeductionZhTw
	case global.ElectricBill:
		return global.ElectricBillZhTw
	case global.Fish:
		return global.FishZhTW
	case global.Medical:
		return global.MedicalZhTw
	case global.Ticket:
		return global.TicketZhTw
	case global.Gardening:
		return global.GardeningZhTw
	case global.GroceryShopping:
		return global.GroceryShoppingZhTw
	case global.EasyCard:
		return global.EasyCardZhTw
	case global.ManagementCost:
		return global.ManagementCostZhTw
	case global.PayBill:
		return global.PayBillZhTw
	case global.PottedPlant:
		return global.PottedPlantZhTw
	case global.CreditCard:
		return global.CreditCardZhTw
	case global.Cash:
		return global.CashZhTw
	case global.Salary:
		return global.SalaryZhTw
	case global.OtherIncome:
		return global.OtherIncomeZhTw
	default:
		return ""
	}
}

func MonthConvertToZhTw(month int) string {
	switch month {
	case 1:
		return global.JanZhTw
	case 2:
		return global.FebZhTw
	case 3:
		return global.MarZhTw
	case 4:
		return global.AprZhTw
	case 5:
		return global.MayZhTw
	case 6:
		return global.JunZhTw
	case 7:
		return global.JulZhTw
	case 8:
		return global.AugZhTw
	case 9:
		return global.SepZhTw
	case 10:
		return global.OctZhTw
	case 11:
		return global.NovZhTw
	case 12:
		return global.DecZhTw
	default:
		return ""
	}
}

func MonthConvertToKeyName(month int) string {
	switch month {
	case 1:
		return global.JanSummary
	case 2:
		return global.FebSummary
	case 3:
		return global.MarSummary
	case 4:
		return global.AprSummary
	case 5:
		return global.MaySummary
	case 6:
		return global.JunSummary
	case 7:
		return global.JulSummary
	case 8:
		return global.AugSummary
	case 9:
		return global.SepSummary
	case 10:
		return global.OctSummary
	case 11:
		return global.NovSummary
	case 12:
		return global.DecSummary
	default:
		return ""
	}
}

func KeyNameConvertToMonth(key string) int {
	switch key {
	case global.JanSummary:
		return 1
	case global.FebSummary:
		return 2
	case global.MarSummary:
		return 3
	case global.AprSummary:
		return 4
	case global.MaySummary:
		return 5
	case global.JunSummary:
		return 6
	case global.JulSummary:
		return 7
	case global.AugSummary:
		return 8
	case global.SepSummary:
		return 9
	case global.OctSummary:
		return 10
	case global.NovSummary:
		return 11
	case global.DecSummary:
		return 12
	default:
		return 0
	}
}
