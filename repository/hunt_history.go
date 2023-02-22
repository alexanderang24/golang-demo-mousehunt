package repository

import (
	"database/sql"
	"golang-demo-mousehunt/structs"
	"time"
)

func GetAllHuntHistories(db *sql.DB) ([]structs.HuntHistory, error) {
	var results []structs.HuntHistory
	sqlStatement := "SELECT * FROM hunt_history ORDER BY id"

	var rows, err = db.Query(sqlStatement)
	if err != nil {
		return results, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var history = structs.HuntHistory{}

		err = rows.Scan(&history.ID, &history.UserID, &history.MouseID, &history.LocationID, &history.TrapID, &history.CreatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, history)
	}
	return results, nil
}

func InsertHuntHistory(db *sql.DB, history structs.HuntHistory) (structs.HuntHistory, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO hunt_history(user_id, mouse_id, location_id, trap_id, created_at) VALUES($1, $2, $3, $4, $5)"
	rows := db.QueryRow(sqlStatement, history.UserID, history.MouseID, history.LocationID, history.TrapID, now)
	if rows.Err() != nil {
		err = rows.Err()
		return history, err
	} else {
		err = rows.Scan(&history.ID)
		if err != nil {
			return history, err
		}
		return history, nil
	}
}