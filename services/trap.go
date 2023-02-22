package services

import (
	"errors"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

var (
	err error
)

type TrapService interface {
	GetAllTraps() ([]structs.Trap, error)
	GetTrap(trap structs.Trap) (structs.Trap, error)
	InsertTrap(trap structs.Trap) (structs.Trap, error)
	UpdateTrap(trap structs.Trap) (structs.Trap, error)
	DeleteTrap(trap structs.Trap) (structs.Trap, error)
	BuyTrap(trap structs.Trap, user structs.User) (structs.User, error)
}

type trapService struct {
	repository repository.TrapRepository
}

func NewTrapService(repo repository.TrapRepository) *trapService {
	return &trapService{repo}
}

func (s *trapService) GetAllTraps() ([]structs.Trap, error) {
	var traps, err = s.repository.GetAllTraps()
	if err != nil {
		return traps, err
	} else {
		return traps, nil
	}
}

func (s *trapService) GetTrap(trap structs.Trap) (structs.Trap, error) {
	trap, err = s.repository.GetTrap(trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func (s *trapService) InsertTrap(trap structs.Trap) (structs.Trap, error) {
	if trap.MaxPower < trap.MinPower {
		err = errors.New("max power should not be lower than min power")
		return trap, err
	}

	trap, err = s.repository.InsertTrap(trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func (s *trapService) UpdateTrap(trap structs.Trap) (structs.Trap, error) {
	if trap.MaxPower < trap.MinPower {
		err = errors.New("max power should not be lower than min power")
		return trap, err
	}

	trap, err = s.repository.UpdateTrap(trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func (s *trapService) DeleteTrap(trap structs.Trap) (structs.Trap, error) {
	trap, err = s.repository.DeleteTrap(trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func (s *trapService) BuyTrap(trap structs.Trap, user structs.User) (structs.User, error) {
	if user.Gold < trap.Price {
		err := errors.New("not enough gold to buy trap")
		return user, err
	} else if user.TrapID == trap.ID {
		err := errors.New("you already have this trap")
		return user, err
	} else {
		user.Gold = user.Gold - trap.Price
		user.TrapID = trap.ID

		var ur repository.UserRepository
		user, err = ur.UpdateUser(user)
		return user, nil
	}
}
