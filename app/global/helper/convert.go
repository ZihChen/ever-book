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
	}
	return ""
}
