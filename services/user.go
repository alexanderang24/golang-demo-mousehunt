package services

import (
	"errors"
	"golang-demo-mousehunt/middleware"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

type UserService interface {
	GetAllUsers() ([]structs.User, error)
	GetUser(user structs.User) (structs.User, error)
	Register(user structs.User) (structs.User, error)
	InsertUser(user structs.User) (structs.User, error)
	UpdateUser(user structs.User) (structs.User, error)
	DeleteUser(user structs.User) (structs.User, error)
	Login(user structs.User) (structs.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]structs.User, error) {
	var users, err = s.repository.GetAllUsers()
	if err != nil {
		return users, err
	} else {
		return users, nil
	}
}

func (s *userService) GetUser(user structs.User) (structs.User, error) {
	user, err = s.repository.GetUser(user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (s* userService) Register(user structs.User) (structs.User, error) {
	user.Role = "player"
	user.Gold = 0
	user.LocationID = 1
	user.TrapID = 1
	return s.InsertUser(user)
}

func (s *userService) InsertUser(user structs.User) (structs.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err = errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	user, err = s.repository.InsertUser(user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (s *userService) UpdateUser(user structs.User) (structs.User, error) {
	if user.Role != "player" && user.Role != "admin" {
		err = errors.New("invalid role: " + user.Role + ". Value must be player or admin")
		return user, err
	}

	user, err = s.repository.UpdateUser(user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (s *userService) DeleteUser(user structs.User) (structs.User, error) {
	user, err = s.repository.DeleteUser(user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (s *userService) Login(user structs.User) (structs.User, error) {
	// check username and password correct
	userInDb, err := s.repository.GetUserByUsername(user.Username)
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