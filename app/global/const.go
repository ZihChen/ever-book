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
	NeedRemark         string = "need_remark"
	SkipRemark         string = "skip_remark"
	ConfirmDelete      string = "confirm_delete"
	CancelDelete       string = "cancel_delete"
)

const (
	ContinueZhTw             string = "繼續"
	DiscardZhTw              string = "捨棄"
	NeedRemarkZhTw           string = "需要"
	SkipRemarkZhTw           string = "省略"
	RecordBalanceZhTw        string = "我要記帳"
	IncomeZhTw               string = "收入"
	ExpenseZhTw              string = "支出"
	DateZhTw                 string = "日期"
	TodayZhTw                string = "今日"
	CashZhTw                 string = "現金"
	CreditCardZhTw           string = "信用卡"
	ConsumeGoodsZhTw         string = "日用品"
	FruitZhTw                string = "水果"
	WaterBillZhTw            string = "水費"
	OilFeeZhTw               string = "加油"
	BreakfastZhTw            string = "早餐"
	LunchZhTw                string = "午餐"
	DinnerZhTw               string = "晚餐"
	RepairRewardZhTw         string = "外修費"
	GasFeeZhTw               string = "瓦斯費"
	InsuranceZhTw            string = "保險費"
	LivingExpensesZhTw       string = "生活費"
	OrganicFoodZhTw          string = "有機食品"
	DressFeeZhTw             string = "衣服"
	HealthyFoodZhTw          string = "保養品"
	AutomaticDeductionZhTw   string = "帳簿自動扣款"
	ElectricBillZhTw         string = "電費"
	FishZhTW                 string = "魚"
	MedicalZhTw              string = "醫療"
	TicketZhTw               string = "買車票"
	GardeningZhTw            string = "園藝"
	GroceryShoppingZhTw      string = "買菜"
	EasyCardZhTw             string = "悠遊卡儲值"
	ManagementCostZhTw       string = "管理費"
	PayBillZhTw              string = "繳費"
	PottedPlantZhTw          string = "盆栽"
	TypeZhTw                 string = "類型"
	ItemZhTw                 string = "項目"
	AmountZhTw               string = "金額"
	PaymentZhTw              string = "付款方式"
	RemarkZhTw               string = "備註"
	SuccessRecordZhTw        string = "成功新增一筆收支紀錄!"
	SuccessDiscardZhTw       string = "捨棄收支紀錄"
	EnterRemarkZhTw          string = "請輸入備註："
	EnterAmountZhTw          string = "請輸入金額："
	UnfinishedBalanceZhTw    string = "你有一筆紀錄尚未填寫完!"
	BalanceDateOptionZhTw    string = "選擇支出/收入日期"
	BalanceTypeOptionZhTw    string = "請選擇收入或是支出"
	BalanceItemOptionZhTw    string = "選擇消費項目"
	BalancePaymentOptionZhTw string = "選擇付款方式"
	BalanceRemarkOptionZhTw  string = "需不需要填寫備註"
	BalanceMonthOptionZhTw   string = "選擇月份"
	ContinueOrNotOptionZhTw  string = "是否要繼續輸入或是捨棄"
	AccountBookSummaryZhTw   string = "查看帳本"
	OtherAccountBookZhTw     string = "對方帳本"
	DeletePreviousRecordZhTw string = "刪除上一筆"
	CancelOrNotZhTw          string = "是否刪除?"
	ConfirmZhTw              string = "確認"
	CancelZhTw               string = "取消"
	SuccessDeleteZhTw        string = "刪除成功"
)

type TemporaryBalanceColumn string

const (
	TemporaryBalanceType    TemporaryBalanceColumn = "type"
	TemporaryBalanceItem    TemporaryBalanceColumn = "item"
	TemporaryBalanceAmount  TemporaryBalanceColumn = "amount"
	TemporaryBalancePayment TemporaryBalanceColumn = "payment"
	TemporaryBalanceRemark  TemporaryBalanceColumn = "remark"
)

const (
	JanZhTw string = "一月"
	FebZhTw string = "二月"
	MarZhTw string = "三月"
	AprZhTw string = "四月"
	MayZhTw string = "五月"
	JunZhTw string = "六月"
	JulZhTw string = "七月"
	AugZhTw string = "八月"
	SepZhTw string = "九月"
	OctZhTw string = "十月"
	NovZhTw string = "十一月"
	DecZhTw string = "十二月"
)

const (
	JanSummary string = "january_sum"
	FebSummary string = "february_sum"
	MarSummary string = "march_sum"
	AprSummary string = "april_sum"
	MaySummary string = "may_sum"
	JunSummary string = "june_sum"
	JulSummary string = "july_sum"
	AugSummary string = "august_sum"
	SepSummary string = "september_sum"
	OctSummary string = "october_sum"
	NovSummary string = "november_sum"
	DecSummary string = "december_sum"
)
