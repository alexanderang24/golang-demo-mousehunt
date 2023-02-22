package dto

import "time"

type HuntHistory struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	MouseID    int64  `json:"mouse_id"`
	LocationID int64  `json:"location_id"`
	TrapID     int64  `json:"trap_id"`
	Success    bool   `json:"success"`
	CreatedAt  time.Time `json:"created_at"`
}
