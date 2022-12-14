package global

const (
	Income             string = "income"
	Expense            string = "expense"
	Cash               string = "cash"
	CreditCard         string = "credit_card"
	ConsumeGoods       string = "consume_goods"
	Fruit              string = "fruit"
	WaterBill          string = "water_bill"
	OilFee             string = "oil_fee"
	Breakfast          string = "breakfast"
	Lunch              string = "lunch"
	Dinner             string = "dinner"
	RepairReward       string = "repair_reward"
	GasFee             string = "gas_fee"
	Insurance          string = "insurance"
	LivingExpenses     string = "living_expenses"
	OrganicFood        string = "organic_food"
	DressFee           string = "dress_fee"
	HealthyFood        string = "healthy_food"
	AutomaticDeduction string = "automatic_deduction"
	ElectricBill       string = "electric_bill"
	Fish               string = "fish"
	Medical            string = "medical"
	Ticket             string = "ticket"
	Gardening          string = "gardening"
	GroceryShopping    string = "grocery_shopping"
	EasyCard           string = "easy_card"
	ManagementCost     string = "management_cost"
	PayBill            string = "pay_bill"
	PottedPlant        string = "potted_plant"
	Date               string = "date"
	Today              string = "today"
	Continue           string = "continue"
	Discard            string = "discard"
	Type               string = "type"
	Item               string = "item"
	Amount             string = "amount"
	Payment            string = "payment"
	Remark             string = "remark"
)

const (
	ContinueZhTw           string = "繼續"
	DiscardZhTw            string = "捨棄"
	NeedZhTw               string = "需要"
	SkipZhTw               string = "省略"
	RecordBalanceZhTw      string = "我要記帳"
	IncomeZhTw             string = "收入"
	ExpenseZhTw            string = "支出"
	DateZhTw               string = "日期"
	TodayZhTw              string = "今日"
	CashZhTw               string = "現金"
	CreditCardZhTw         string = "信用卡"
	ConsumeGoodsZhTw       string = "日用品"
	FruitZhTw              string = "水果"
	WaterBillZhTw          string = "水費"
	OilFeeZhTw             string = "加油"
	BreakfastZhTw          string = "早餐"
	LunchZhTw              string = "午餐"
	DinnerZhTw             string = "晚餐"
	RepairRewardZhTw       string = "外修費"
	GasFeeZhTw             string = "瓦斯費"
	InsuranceZhTw          string = "保險費"
	LivingExpensesZhTw     string = "生活費"
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
	EasyCardZhTw           string = "悠遊卡儲值"
	ManagementCostZhTw     string = "管理費"
	PayBillZhTw            string = "繳費"
	PottedPlantZhTw        string = "盆栽"
	TypeZhTw               string = "類型"
	ItemZhTw               string = "項目"
	AmountZhTw             string = "金額"
	PaymentZhTw            string = "付款方式"
	RemarkZhTw             string = "備註"
)

type TemporaryBalanceColumn string

const (
	TemporaryBalanceType    TemporaryBalanceColumn = "type"
	TemporaryBalanceItem    TemporaryBalanceColumn = "item"
	TemporaryBalanceAmount  TemporaryBalanceColumn = "amount"
	TemporaryBalancePayment TemporaryBalanceColumn = "payment"
	TemporaryBalanceRemark  TemporaryBalanceColumn = "remark"
)
