package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/dto"
	"golang-demo-mousehunt/dto/response"
	"golang-demo-mousehunt/repository"
	"math/rand"
	"strconv"
)

func GetAllHistories(user dto.User) ([]response.HuntHistoryResponse, error) {
	histories, err := repository.GetHuntHistoriesByUserId(database.DbConnection, user.ID)

	var responses []response.HuntHistoryResponse
	for _, history := range histories {
		location, _ := GetLocation(dto.Location{ID: history.LocationID})
		mouse, _ := GetMouse(dto.Mouse{ID: history.MouseID})
		var gold int64 = 0
		if history.Success == true {
			gold = mouse.Gold
		}
		//date := history.CreatedAt.Format("2006-01-02 15:04:05")
		date := history.CreatedAt.Format("Monday, 02 January 2006 at 15:04:05")

		responses = append(responses, response.HuntHistoryResponse{
			Date:       date,
			Location:   location.Name,
			MouseName:  mouse.Name,
			Success:    history.Success,
			GoldGained: gold,
		})
	}

	if err != nil {
		return responses, err
	} else {
		return responses, nil
	}
}

func DoHunt(user dto.User) (response.HuntResponse, error) {
	// hunt logics here
	var huntResponse response.HuntResponse

	// get mice in location
	mice, err := GetAllMiceInLocation(user.LocationID)
	if err != nil {
		err = errors.New("error when getting mice data")
		return huntResponse, err
	}

	// random mouse encounter and generate its power
	var index = rand.Intn(len(mice))
	var mouse = mice[index]
	var mMinPower = int(mouse.MinPower)
	var mMaxPower = int(mouse.MaxPower)
	var mPower = rand.Intn(mMaxPower-mMinPower+1) + mMinPower

	// get user trap
	trap, err := GetTrap(dto.Trap{ID: user.TrapID})
	if err != nil {
		err = errors.New("error when getting trap data")
		return huntResponse, err
	}

	// random trap power
	var tMinPower = int(trap.MinPower)
	var tMaxPower = int(trap.MaxPower)
	var tPower = rand.Intn(tMaxPower-tMinPower+1) + tMinPower

	// calculate hunt
	var history dto.HuntHistory
	if tPower < mPower {
		history.Success = false
	} else {
		history.Success = true
		user.Gold = user.Gold + mouse.Gold
	}
	history.UserID = user.ID
	history.MouseID = mouse.ID
	history.LocationID = user.LocationID
	history.TrapID = user.TrapID

	// update user
	user, err = UpdateUser(user)
	if err != nil {
		return huntResponse, err
	}

	// add history
	history, err = repository.InsertHuntHistory(database.DbConnection, history)
	if err != nil {
		return huntResponse, err
	}

	// set huntResponse
	location, _ := GetLocation(dto.Location{ID: user.LocationID})
	huntResponse.Location = location.Name
	huntResponse.MouseName = mouse.Name
	huntResponse.MousePower = strconv.Itoa(mPower)
	huntResponse.TrapPower = strconv.Itoa(tPower)
	huntResponse.GoldGained = strconv.Itoa(int(mouse.Gold))
	huntResponse.GoldTotal = strconv.Itoa(int(user.Gold))
	huntResponse.Success = history.Success
	return huntResponse, nil
}
