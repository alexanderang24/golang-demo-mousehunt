package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/dto"
	"strconv"
	"time"
)

func GetAllLocations(db *sql.DB) ([]dto.Location, error) {
	var results []dto.Location
	sqlStatement := "SELECT * FROM location ORDER BY id"

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
		var location = dto.Location{}

		err = rows.Scan(&location.ID, &location.Name, &location.Description, &location.TravelCost, &location.CreatedAt, &location.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, location)
	}
	return results, nil
}

func GetLocation(db *sql.DB, location dto.Location) (dto.Location, error) {
	var result dto.Location
	sqlCheck := "SELECT * FROM location WHERE id = $1"

	rows := db.QueryRow(sqlCheck, location.ID)

	err := rows.Scan(&result.ID, &result.Name, &result.Description, &result.TravelCost, &result.CreatedAt, &result.UpdatedAt)
	if result == (dto.Location{}) {
		err = errors.New("location with id " + strconv.Itoa(int(location.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func InsertLocation(db *sql.DB, location dto.Location) (dto.Location, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO location(name, description, travel_cost, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	rows := db.QueryRow(sqlStatement, location.Name, location.Description, location.TravelCost, now, now)
	if rows.Err() != nil {
		err := rows.Err()
		return location, err
	} else {
		err := rows.Scan(&location.ID)
		if err != nil {
			return location, err
		}
		return location, nil
	}
}

func UpdateLocation(db *sql.DB, location dto.Location) (dto.Location, error) {
	_, err := GetLocation(db, location)
	if err != nil {
		return location, err
	}
	sqlStatement := "UPDATE location SET name = $1, description = $2, travel_cost = $3, updated_at = $4 WHERE id = $5"
	rows := db.QueryRow(sqlStatement, location.Name, location.Description, location.TravelCost, time.Now(), location.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return location, err
	} else {
		return location, nil
	}
}

func DeleteLocation(db *sql.DB, location dto.Location) (dto.Location, error) {
	location, err := GetLocation(db, location)
	if err != nil {
		return location, err
	}
	sqlStatement := "DELETE FROM location WHERE id = $1"
	rows := db.QueryRow(sqlStatement, location.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return location, err
	} else {
		return location, nil
	}
}