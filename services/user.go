package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/middleware"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

func GetAllUsers() ([]structs.User, error) {
	var users, err = repository.GetAllUsers(database.DbConnection)
	if err != nil {
		return users, err
	} else {
		return users, nil
	}
}

func GetUser(user structs.User) (structs.User, error) {
	user, err = repository.GetUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func Register(user structs.User) (structs.User, error) {
	user.Role = "player"
	user.Gold = 0
	user.LocationID = 1
	user.TrapID = 1
	return InsertUser(user)
}

func InsertUser(user structs.User) (structs.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err = errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	user, err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func UpdateUser(user structs.User) (structs.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err = errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	user, err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func DeleteUser(user structs.User) (structs.User, error) {
	user, err = repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func Login(user structs.User) (structs.User, error) {
	// check username and password correct
	userInDb, err := repository.GetUserByUsername(database.DbConnection, user.Username)
	if err != nil {
		return user, err
	}
	if user.Password != userInDb.Password {
		err = errors.New("incorrect password")
		return user, err
	}

	token, err := middleware.GenerateJWT(userInDb.Username, userInDb.Role)
	if err != nil {
		return user, err
	} else {
		user.Token = token
		return user, nil
	}
}

func GetByUsername(username string) (structs.User, error) {
	user, err := repository.GetUserByUsername(database.DbConnection, username)
	if err != nil {
		return user, err
	}
	return user, nil
}