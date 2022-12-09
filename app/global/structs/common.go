package structs

type Balance struct {
	Date    string `json:"date"`
	Type    string `json:"type"`
	Item    string `json:"item"`
	Amount  int    `json:"amount"`
	Payment string `json:"payment"`
	Remark  string `json:"remark"`
}
