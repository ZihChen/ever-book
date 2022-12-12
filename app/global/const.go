package global

const (
	Income  string = "income"
	Expense string = "expense"

	RecordBalanceZhTw string = "我要記帳"
	IncomeZhTw        string = "收入"
	ExpenseZhTw       string = "支出"
	TodayZhTw         string = "今日"
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
