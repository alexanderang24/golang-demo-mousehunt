package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/dto"
	"strconv"
	"time"
)

var (
	err error
)

func GetAllTraps(db *sql.DB) ([]dto.Trap, error) {
	var results []dto.Trap
	sqlStatement := "SELECT * FROM trap ORDER BY id"

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
		var trap = dto.Trap{}

		err = rows.Scan(&trap.ID, &trap.Name, &trap.Description, &trap.MinPower, &trap.MaxPower, &trap.Price, &trap.CreatedAt, &trap.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, trap)
	}
	return results, nil
}

func GetTrap(db *sql.DB, trap dto.Trap) (dto.Trap, error) {
	var result dto.Trap
	sqlCheck := "SELECT * FROM trap WHERE id = $1"

	rows := db.QueryRow(sqlCheck, trap.ID)

	err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.MinPower, &result.MaxPower, &result.Price, &result.CreatedAt, &result.UpdatedAt)
	if result == (dto.Trap{}) {
		err = errors.New("trap with id " + strconv.Itoa(int(trap.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func InsertTrap(db *sql.DB, trap dto.Trap) (dto.Trap, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO trap(name, description, min_power, max_power, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	rows := db.QueryRow(sqlStatement, trap.Name, trap.Description, trap.MinPower, trap.MaxPower, trap.Price, now, now)
	if rows.Err() != nil {
		err = rows.Err()
		return trap, err
	} else {
		err = rows.Scan(&trap.ID)
		if err != nil {
			return trap, err
		}
		return trap, nil
	}
}

func UpdateTrap(db *sql.DB, trap dto.Trap) (dto.Trap, error) {
	_, err = GetTrap(db, trap)
	if err != nil {
		return trap, err
	}
	sqlStatement := "UPDATE trap SET name = $1, description = $2, min_power = $3, max_power = $4, price = $5, updated_at = $6 WHERE id = $7"
	rows := db.QueryRow(sqlStatement, trap.Name, trap.Description, trap.MinPower, trap.MaxPower, trap.Price, time.Now(), trap.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return trap, err
	} else {
		return trap, nil
	}
}

func DeleteTrap(db *sql.DB, trap dto.Trap) (dto.Trap, error) {
	sqlStatement := "DELETE FROM trap WHERE id = $1"
	rows := db.QueryRow(sqlStatement, trap.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return trap, err
	} else {
		return trap, nil
	}
}