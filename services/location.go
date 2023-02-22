package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/dto"
)

func GetAllLocations() ([]dto.Location, error) {
	var locations, err = repository.GetAllLocations(database.DbConnection, )
	if err != nil {
		return locations, err
	} else {
		return locations, nil
	}
}

func GetLocation(location dto.Location) (dto.Location, error) {
	location, err = repository.GetLocation(database.DbConnection, location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func InsertLocation(location dto.Location) (dto.Location, error) {
	location, err = repository.InsertLocation(database.DbConnection, location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func UpdateLocation(location dto.Location) (dto.Location, error) {
	location, err = repository.UpdateLocation(database.DbConnection, location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func DeleteLocation(location dto.Location) (dto.Location, error) {
	location, err = repository.DeleteLocation(database.DbConnection, location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func TravelToLocation(location dto.Location, user dto.User) (dto.User, error) {
	if user.LocationID == location.ID {
		err := errors.New("you already on this location")
		return user, err
	} else if user.Gold < location.TravelCost {
		err := errors.New("not enough gold to travel")
		return user, err
	} else {
		user.Gold = user.Gold - location.TravelCost
		user.LocationID = location.ID

		user, err = UpdateUser(user)
		return user, nil
	}
}
