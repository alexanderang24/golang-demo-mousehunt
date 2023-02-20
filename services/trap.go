package services

import (
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
	trap, err = s.repository.InsertTrap(trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func (s *trapService) UpdateTrap(trap structs.Trap) (structs.Trap, error) {
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