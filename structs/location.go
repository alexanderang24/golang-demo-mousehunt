package structs

type Location struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	TravelCost int64  `json:"travel_cost"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
