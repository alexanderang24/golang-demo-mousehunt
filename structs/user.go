package structs

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Gold       int64  `json:"gold"`
	LocationID int64  `json:"location_id"`
	TrapID     int64  `json:"trap_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
