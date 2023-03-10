package dto

type Mouse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MinPower    int64  `json:"min_power"`
	MaxPower    int64  `json:"max_power"`
	Gold        int64  `json:"gold"`
	LocationID  int64  `json:"location_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
