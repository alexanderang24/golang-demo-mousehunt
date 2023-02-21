package services

import (
	"errors"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

type MouseService interface {
	GetAllMice() ([]structs.Mouse, error)
	GetMouse(mouse structs.Mouse) (structs.Mouse, error)
	InsertMouse(mouse structs.Mouse) (structs.Mouse, error)
	UpdateMouse(mouse structs.Mouse) (structs.Mouse, error)
	DeleteMouse(mouse structs.Mouse) (structs.Mouse, error)
}

type mouseService struct {
	repository repository.MouseRepository
}

func NewMouseService(repo repository.MouseRepository) *mouseService {
	return &mouseService{repo}
}

func (s *mouseService) GetAllMice() ([]structs.Mouse, error) {
	var mouses, err = s.repository.GetAllMice()
	if err != nil {
		return mouses, err
	} else {
		return mouses, nil
	}
}

func (s *mouseService) GetMouse(mouse structs.Mouse) (structs.Mouse, error) {
	mouse, err = s.repository.GetMouse(mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func (s *mouseService) InsertMouse(mouse structs.Mouse) (structs.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	mouse, err = s.repository.InsertMouse(mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func (s *mouseService) UpdateMouse(mouse structs.Mouse) (structs.Mouse, error) {
	if mouse.MaxPower < mouse.MinPower {
		err = errors.New("max power should not be lower than min power")
		return mouse, err
	}

	mouse, err = s.repository.UpdateMouse(mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}

func (s *mouseService) DeleteMouse(mouse structs.Mouse) (structs.Mouse, error) {
	mouse, err = s.repository.DeleteMouse(mouse)
	if err != nil {
		return mouse, err
	} else {
		return mouse, nil
	}
}