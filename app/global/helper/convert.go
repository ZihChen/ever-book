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
	default:
		return ""
	}
}
