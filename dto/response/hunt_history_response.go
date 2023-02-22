package response

type HuntHistoryResponse struct {
	Success    bool      `json:"success"`
	Date       string	 `json:"date"`
	Location   string    `json:"location"`
	MouseName  string    `json:"mouse"`
	GoldGained int64     `json:"gold_gained"`
}
