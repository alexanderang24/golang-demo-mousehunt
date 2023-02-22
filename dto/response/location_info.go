package response

type LocationInfo struct {
	LocationName        string       `json:"location_name"`
	LocationDescription string       `json:"location_description"`
	TravelCost          int64        `json:"travel_cost"`
	Mice                []MiceInHere `json:"mice_in_here"`
}

type MiceInHere struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MinPower    int64  `json:"min_power"`
	MaxPower    int64  `json:"max_power"`
	GoldReward  int64  `json:"gold_reward"`
}
