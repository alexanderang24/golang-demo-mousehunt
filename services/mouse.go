package services

import (
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

func GetAllMice() ([]structs.Mouse, error) {
	var mouses, err = repository.GetAllMice(database.DbConnection)
	if err != nil {
		return mouses, err
	} else {
		return mouses, nil
	}
}

func GetMouse(mouse structs.Mouse) (structs.Mouse, error) {
	mouse, err = repository.GetMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func InsertMouse(mouse structs.Mouse) (structs.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	mouse, err = repository.InsertMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func UpdateMouse(mouse structs.Mouse) (structs.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	mouse, err = repository.UpdateMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func DeleteMouse(mouse structs.Mouse) (structs.Mouse, error) {
	mouse, err = repository.DeleteMouse(database.DbConnection, mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}