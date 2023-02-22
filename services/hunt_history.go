package services

import (
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

func GetAllHistories() ([]structs.HuntHistory, error) {
	var histories, err = repository.GetAllHuntHistories(database.DbConnection)
	if err != nil {
		return histories, err
	} else {
		return histories, nil
	}
}

func DoHunt(history structs.HuntHistory) (structs.HuntHistory, error) {
	// hunt logics here



	history, err = repository.InsertHuntHistory(database.DbConnection, history)
	if err != nil {
		return history, err
	} else {
		return history, nil
	}
}