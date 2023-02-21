package structs

type Trap struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MinPower    int64  `json:"min_power"`
	MaxPower    int64  `json:"max_power"`
	Price       int64  `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
