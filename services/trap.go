package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/dto"
	"golang-demo-mousehunt/repository"
)

func GetAllTraps() ([]dto.Trap, error) {
	var traps, err = repository.GetAllTraps(database.DbConnection)
	if err != nil {
		return traps, err
	} else {
		return traps, nil
	}
}

func GetTrap(trap dto.Trap) (dto.Trap, error) {
	trap, err := repository.GetTrap(database.DbConnection, trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func InsertTrap(trap dto.Trap) (dto.Trap, error) {
	if trap.MaxPower < trap.MinPower {
		err := errors.New("max power should not be lower than min power")
		return trap, err
	}

	trap, err := repository.InsertTrap(database.DbConnection, trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func UpdateTrap(trap dto.Trap) (dto.Trap, error) {
	if trap.MaxPower < trap.MinPower {
		err := errors.New("max power should not be lower than min power")
		return trap, err
	}
	trap, err := repository.UpdateTrap(database.DbConnection, trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func DeleteTrap(trap dto.Trap) (dto.Trap, error) {
	trap, err := GetTrap(trap)
	if err != nil {
		return trap, err
	}

	trap, err = repository.DeleteTrap(database.DbConnection, trap)
	if err != nil {
		return trap, err
	} else {
		return trap, nil
	}
}

func BuyTrap(trap dto.Trap, user dto.User) (dto.User, error) {
	if user.TrapID == trap.ID {
		err := errors.New("you already have this trap")
		return user, err
	} else if user.Gold < trap.Price {
		err := errors.New("not enough gold to buy trap")
		return user, err
	} else {
		user.Gold = user.Gold - trap.Price
		user.TrapID = trap.ID
		user, _ = UpdateUser(user)
		return user, nil
	}
}
