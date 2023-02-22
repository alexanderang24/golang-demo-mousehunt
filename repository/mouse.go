package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/dto"
	"strconv"
	"time"
)

func GetAllMice(db *sql.DB) ([]dto.Mouse, error) {
	var results []dto.Mouse
	sqlStatement := "SELECT * FROM mouse ORDER BY id"

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
		var mouse = dto.Mouse{}

		err = rows.Scan(&mouse.ID, &mouse.Name, &mouse.Description, &mouse.MinPower, &mouse.MaxPower, &mouse.Gold, &mouse.LocationID, &mouse.CreatedAt, &mouse.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, mouse)
	}
	return results, nil
}

func GetAllMiceInLocation(db *sql.DB, locationID int64) ([]dto.Mouse, error) {
	var results []dto.Mouse
	sqlStatement := "SELECT * FROM mouse WHERE location_id = $1 ORDER BY id"

	var rows, err = db.Query(sqlStatement, locationID)
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
		var mouse = dto.Mouse{}

		err = rows.Scan(&mouse.ID, &mouse.Name, &mouse.Description, &mouse.MinPower, &mouse.MaxPower, &mouse.Gold, &mouse.LocationID, &mouse.CreatedAt, &mouse.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, mouse)
	}
	return results, nil
}

func GetMouse(db *sql.DB, mouse dto.Mouse) (dto.Mouse, error) {
	var result dto.Mouse
	sqlCheck := "SELECT * FROM mouse WHERE id = $1"
	rows := db.QueryRow(sqlCheck, mouse.ID)

	err := rows.Scan(&result.ID, &result.Name, &result.Description, &result.MinPower, &result.MaxPower, &result.Gold, &result.LocationID, &result.CreatedAt, &result.UpdatedAt)
	if result == (dto.Mouse{}) {
		err = errors.New("mouse with id " + strconv.Itoa(int(mouse.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func InsertMouse(db *sql.DB, mouse dto.Mouse) (dto.Mouse, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO mouse(name, description, min_power, max_power, gold, location_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	rows := db.QueryRow(sqlStatement, mouse.Name, mouse.Description, mouse.MinPower, mouse.MaxPower, mouse.Gold, mouse.LocationID, now, now)
	if rows.Err() != nil {
		err := rows.Err()
		return mouse, err
	} else {
		err := rows.Scan(&mouse.ID)
		if err != nil {
			return mouse, err
		}
		return mouse, nil
	}
}

func UpdateMouse(db *sql.DB, mouse dto.Mouse) (dto.Mouse, error) {
	sqlStatement := "UPDATE mouse SET name = $1, description = $2, min_power = $3, max_power = $4, gold = $5, location_id = $6, updated_at = $7 WHERE id = $8"
	rows := db.QueryRow(sqlStatement, mouse.Name, mouse.Description, mouse.MinPower, mouse.MaxPower, mouse.Gold, mouse.LocationID, time.Now(), mouse.ID)
	if rows.Err() != nil {
		err := rows.Err()
		return mouse, err
	} else {
		return mouse, nil
	}
}

func DeleteMouse(db *sql.DB, mouse dto.Mouse) (dto.Mouse, error) {
	sqlStatement := "DELETE FROM mouse WHERE id = $1"
	rows := db.QueryRow(sqlStatement, mouse.ID)
	if rows.Err() != nil {
		err := rows.Err()
		return mouse, err
	} else {
		return mouse, nil
	}
}
