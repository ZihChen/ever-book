package global

var BalanceItems map[string][]string = map[string][]string{
	"支出-日常": {
		LivingExpensesZhTw,
		ConsumeGoodsZhTw,
		OilFeeZhTw,
	},
	"支出-菜市場": {
		GroceryShoppingZhTw,
		FishZhTW,
		FruitZhTw,
	},
	"支出-外食": {
		BreakfastZhTw,
		LunchZhTw,
		DinnerZhTw,
	},
	"支出-繳費": {
		WaterBillZhTw,
		ElectricBillZhTw,
		GasFeeZhTw,
	},
	"支出-繳費2": {
		TelephoneFeeZhTw,
		ManagementCostZhTw,
		InsuranceZhTw,
	},
	"支出-其他": {
		TicketZhTw,
		EasyCardZhTw,
		DressFeeZhTw,
	},
	"支出-其他2": {
		AutomaticDeductionZhTw,
		OtherExpenseZhTw,
		PayBillZhTw,
	},
	"支出-其他3": {
		OrganicFoodZhTw,
		GardeningZhTw,
		MedicalZhTw,
	},
	"收入-其他": {
		OtherIncomeZhTw,
		RepairRewardZhTw,
		SalaryZhTw,
	},
}
