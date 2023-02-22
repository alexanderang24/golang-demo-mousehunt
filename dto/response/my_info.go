package response

type MyInfo struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	LocationName string `json:"location_name"`
	Gold         int64  `json:"gold"`
	Trap         MyTrap `json:"trap"`
}

type MyTrap struct {
	Name     string `json:"name"`
	MinPower int64  `json:"min_power"`
	MaxPower int64  `json:"max_power"`
}
