package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/dto"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/util"
	"strconv"
)

func GetAllUsers() ([]dto.User, error) {
	var users, err = repository.GetAllUsers(database.DbConnection)
	if err != nil {
		return users, err
	} else {
		return users, nil
	}
}

func GetUser(user dto.User) (dto.User, error) {
	user, err := repository.GetUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func Register(user dto.User) (dto.User, error) {
	existingUser, err := GetByUsername(user.Username)
	if existingUser != (dto.User{}) {
		err = errors.New("username already exist")
		return user, err
	}

	user.Role = "player"
	user.Gold = 0
	user.LocationID = 1
	user.TrapID = 1
	return InsertUser(user)
}

func InsertUser(user dto.User) (dto.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err := errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	var location = dto.Location{
		ID: user.LocationID,
	}
	_, err := repository.GetLocation(database.DbConnection, location)
	if err != nil {
		err = errors.New("location with ID " + strconv.Itoa(int(user.LocationID)) + " not found")
		return user, err
	}

	var trap = dto.Trap{
		ID: user.TrapID,
	}
	_, err = repository.GetTrap(database.DbConnection, trap)
	if err != nil {
		err = errors.New("trap with ID " + strconv.Itoa(int(user.TrapID)) + " not found")
		return user, err
	}

	user, err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func UpdateUser(user dto.User) (dto.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err := errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	var location = dto.Location{
		ID: user.LocationID,
	}
	_, err := repository.GetLocation(database.DbConnection, location)
	if err != nil {
		err = errors.New("location with ID " + strconv.Itoa(int(user.LocationID)) + " not found")
		return user, err
	}

	var trap = dto.Trap{
		ID: user.TrapID,
	}
	_, err = repository.GetTrap(database.DbConnection, trap)
	if err != nil {
		err = errors.New("trap with ID " + strconv.Itoa(int(user.TrapID)) + " not found")
		return user, err
	}

	user, err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func DeleteUser(user dto.User) (dto.User, error) {
	user, err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func Login(user dto.User) (string, error) {
	// check username and password correct
	userInDb, err := repository.GetUserByUsername(database.DbConnection, user.Username)
	if userInDb == (dto.User{}) {
		err = errors.New("username not found")
		return "", err
	}
	if user.Password != userInDb.Password {
		err = errors.New("incorrect password")
		return "", err
	}

	token, err := util.GenerateJWT(userInDb.Username, userInDb.Role)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func GetByUsername(username string) (dto.User, error) {
	user, err := repository.GetUserByUsername(database.DbConnection, username)
	if err != nil {
		return user, err
	}
	return user, nil
}