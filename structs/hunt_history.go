package structs

type HuntHistory struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"gold"`
	MouseID    int64  `json:"mouse_id"`
	LocationID int64  `json:"location_id"`
	TrapID     int64  `json:"trap_id"`
	CreatedAt  string `json:"created_at"`
}
