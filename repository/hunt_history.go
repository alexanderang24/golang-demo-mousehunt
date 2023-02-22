package repository

import (
	"database/sql"
	"golang-demo-mousehunt/dto"
	"time"
)

func GetHuntHistoriesByUserId(db *sql.DB, id int64) ([]dto.HuntHistory, error) {
	var results []dto.HuntHistory
	sqlStatement := "SELECT * FROM hunt_history WHERE user_id = $1 ORDER BY id DESC"

	var rows, err = db.Query(sqlStatement, id)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var history = dto.HuntHistory{}
		err = rows.Scan(&history.ID, &history.UserID, &history.MouseID, &history.LocationID, &history.TrapID, &history.Success, &history.CreatedAt)
		if err != nil {
			return results, err
		}
		results = append(results, history)
	}
	return results, nil
}

func InsertHuntHistory(db *sql.DB, history dto.HuntHistory) (dto.HuntHistory, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO hunt_history(user_id, mouse_id, location_id, trap_id, success, created_at) VALUES($1, $2, $3, $4, $5, $6)"
	rows := db.QueryRow(sqlStatement, history.UserID, history.MouseID, history.LocationID, history.TrapID, history.Success, now)
	if rows.Err() != nil {
		err = rows.Err()
		return history, err
	} else {
		return history, nil
	}
}