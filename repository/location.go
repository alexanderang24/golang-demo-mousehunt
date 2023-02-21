package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/structs"
	"strconv"
	"time"
)

type LocationRepository interface {
	GetAllLocations() ([]structs.Location, error)
	GetLocation(location structs.Location) (structs.Location, error)
	InsertLocation(location structs.Location) (structs.Location, error)
	UpdateLocation(location structs.Location) (structs.Location, error)
	DeleteLocation(location structs.Location) (structs.Location, error)
}

type locationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *locationRepository {
	return &locationRepository{db}
}

func (r *locationRepository) GetAllLocations() ([]structs.Location, error) {
	var results []structs.Location
	sqlStatement := "SELECT * FROM location ORDER BY id"

	var rows, err = r.db.Query(sqlStatement)
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
		var location = structs.Location{}

		err = rows.Scan(&location.ID, &location.Name, &location.Description, &location.TravelCost, &location.CreatedAt, &location.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, location)
	}
	return results, nil
}

func (r *locationRepository) GetLocation(location structs.Location) (structs.Location, error) {
	var result structs.Location
	sqlCheck := "SELECT * FROM location WHERE id = $1"

	rows := r.db.QueryRow(sqlCheck, location.ID)

	err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.TravelCost, &result.CreatedAt, &result.UpdatedAt)
	if result == (structs.Location{}) {
		err = errors.New("location with id " + strconv.Itoa(int(location.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func (r *locationRepository) InsertLocation(location structs.Location) (structs.Location, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO location(name, description, travel_cost, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	rows := r.db.QueryRow(sqlStatement, location.Name, location.Description, location.TravelCost, now, now)
	if rows.Err() != nil {
		err = rows.Err()
		return location, err
	} else {
		err = rows.Scan(&location.ID)
		if err != nil {
			return location, err
		}
		return location, nil
	}
}

func (r *locationRepository) UpdateLocation(location structs.Location) (structs.Location, error) {
	_, err = r.GetLocation(location)
	if err != nil {
		return location, err
	}
	sqlStatement := "UPDATE location SET name = $1, description = $2, travel_cost = $3, updated_at = $4 WHERE id = $5"
	rows := r.db.QueryRow(sqlStatement, location.Name, location.Description, location.TravelCost, time.Now(), location.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return location, err
	} else {
		return location, nil
	}
}

func (r *locationRepository) DeleteLocation(location structs.Location) (structs.Location, error) {
	location, err = r.GetLocation(location)
	if err != nil {
		return location, err
	}
	sqlStatement := "DELETE FROM location WHERE id = $1"
	rows := r.db.QueryRow(sqlStatement, location.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return location, err
	} else {
		return location, nil
	}
}