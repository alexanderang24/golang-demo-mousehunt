package structs

type Trap struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Power     int64  `json:"power"`
	Price     int64  `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
