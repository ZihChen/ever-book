package global

const (
	Income  string = "income"
	Expense string = "expense"

	RecordBalanceZhTw      string = "我要記帳"
	IncomeZhTw             string = "收入"
	ExpenseZhTw            string = "支出"
	TodayZhTw              string = "今日"
	ConsumeGoodsZhTw       string = "日用品"
	FruitZhTw              string = "水果"
	WaterBillZhTw          string = "水費"
	OilFeeZhTw             string = "加油"
	LunchZhTw              string = "外食(午餐)"
	DinnerZhTw             string = "外食(晚餐)"
	RepairRewardZhTw       string = "外修費"
	GasFeeZhTw             string = "瓦斯費"
	LivingExpensesZhTw     string = "生活開銷"
	OrganicFoodZhTw        string = "有機食品"
	DressFeeZhTw           string = "衣服"
	HealthyFoodZhTw        string = "保養品"
	AutomaticDeductionZhTw string = "帳簿自動扣款"
	ElectricBillZhTw       string = "電費"
	FishZhTW               string = "魚"
	MedicalZhTw            string = "醫療"
	TicketZhTw             string = "買車票"
	GardeningZhTw          string = "園藝"
	GroceryShoppingZhTw    string = "買菜"
	EasyCardZhTw           string = "悠遊卡除值"
	ManagementCostZhTw     string = "管理費"
	PayBill                string = "繳費"
	PottedPlantZhTw        string = "盆栽"
)

const (
	DailyConsume string = "日用品"
	Fruit        string = "水果"
	WaterFee     string = "水費"
	GasFee       string = "瓦斯費"
	Launch       string = "外食(午餐)"
	Dinner       string = "外食(晚餐)"
)

type TemporaryBalanceColumn string

const (
	TemporaryBalanceType    TemporaryBalanceColumn = "type"
	TemporaryBalanceItem    TemporaryBalanceColumn = "item"
	TemporaryBalanceAmount  TemporaryBalanceColumn = "amount"
	TemporaryBalancePayment TemporaryBalanceColumn = "payment"
	TemporaryBalanceRemark  TemporaryBalanceColumn = "remark"
)
