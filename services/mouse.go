package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/dto"
	"strconv"
)

func GetAllMice() ([]dto.Mouse, error) {
	var mice, err = repository.GetAllMice(database.DbConnection)
	if err != nil {
		return mice, err
	} else {
		return mice, nil
	}
}

func GetAllMiceInLocation(locationID int64) ([]dto.Mouse, error) {
	var mice, err = repository.GetAllMiceInLocation(database.DbConnection, locationID)
	if err != nil {
		return mice, err
	} else {
		return mice, nil
	}
}

func GetMouse(mouse dto.Mouse) (dto.Mouse, error) {
	mouse, err = repository.GetMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func InsertMouse(mouse dto.Mouse) (dto.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	var location = dto.Location{
		ID: mouse.LocationID,
	}
	_, err = repository.GetLocation(database.DbConnection, location)
	if err != nil {
		err = errors.New("location with ID " + strconv.Itoa(int(mouse.LocationID)) + " not found")
		return mouse, err
	}

	mouse, err = repository.InsertMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func UpdateMouse(mouse dto.Mouse) (dto.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	var location = dto.Location{
		ID: mouse.LocationID,
	}
	_, err = repository.GetLocation(database.DbConnection, location)
	if err != nil {
		err = errors.New("location with ID " + strconv.Itoa(int(mouse.LocationID)) + " not found")
		return mouse, err
	}

	_, err = GetMouse(mouse)
	if err != nil {
		return mouse, err
	}

	mouse, err = repository.UpdateMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func DeleteMouse(mouse dto.Mouse) (dto.Mouse, error) {
	mouse, err = GetMouse(mouse)
	if err != nil {
		return mouse, err
	}

	mouse, err = repository.DeleteMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}