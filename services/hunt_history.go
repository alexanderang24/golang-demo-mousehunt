package services

import (
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

type HuntHistoryService interface {
	GetAllHistories() ([]structs.HuntHistory, error)
	DoHunt(history structs.HuntHistory) (structs.HuntHistory, error)
}

type huntHistoryService struct {
	repository repository.HuntHistoryRepository
}

func NewHistoryService(repo repository.HuntHistoryRepository) *huntHistoryService {
	return &huntHistoryService{repo}
}

func (s *huntHistoryService) GetAllHistories() ([]structs.HuntHistory, error) {
	var historys, err = s.repository.GetAllHuntHistories()
	if err != nil {
		return historys, err
	} else {
		return historys, nil
	}
}

func (s *huntHistoryService) DoHunt(history structs.HuntHistory) (structs.HuntHistory, error) {
	// hunt logics here



	history, err = s.repository.InsertHuntHistory(history)
	if err != nil {
		return history, err
	} else {
		return history, nil
	}
}