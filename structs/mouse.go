package structs

type Mouse struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Power      string `json:"power"`
	Gold       int64  `json:"gold"`
	LocationID int64  `json:"location_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
