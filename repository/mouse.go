package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/structs"
	"strconv"
	"time"
)

func GetAllMice(db *sql.DB) ([]structs.Mouse, error) {
	var results []structs.Mouse
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
		var mouse = structs.Mouse{}

		err = rows.Scan(&mouse.ID, &mouse.Name, &mouse.Description, &mouse.MinPower, &mouse.MaxPower, &mouse.Gold, &mouse.LocationID, &mouse.CreatedAt, &mouse.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, mouse)
	}
	return results, nil
}

func GetMouse(db *sql.DB, mouse structs.Mouse) (structs.Mouse, error) {
	var result structs.Mouse
	sqlCheck := "SELECT * FROM mouse WHERE id = $1"

	rows := db.QueryRow(sqlCheck, mouse.ID)

	err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.MinPower, &result.MaxPower, &result.Gold, &result.LocationID, &result.CreatedAt, &result.UpdatedAt)
	if result == (structs.Mouse{}) {
		err = errors.New("mouse with id " + strconv.Itoa(int(mouse.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func InsertMouse(db *sql.DB, mouse structs.Mouse) (structs.Mouse, error) {
	var location = structs.Location{
		ID: mouse.LocationID,
	}
	_, err = GetLocation(database.DbConnection, location)
	if err != nil {
		return mouse, err
	}

	now := time.Now()
	sqlStatement := "INSERT INTO mouse(name, description, min_power, max_power, gold, location_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	rows := db.QueryRow(sqlStatement, mouse.Name, mouse.Description, mouse.MinPower, mouse.MaxPower, mouse.Gold, mouse.LocationID, now, now)
	if rows.Err() != nil {
		err = rows.Err()
		return mouse, err
	} else {
		err = rows.Scan(&mouse.ID)
		if err != nil {
			return mouse, err
		}
		return mouse, nil
	}
}

func UpdateMouse(db *sql.DB, mouse structs.Mouse) (structs.Mouse, error) {
	var location = structs.Location{
		ID: mouse.LocationID,
	}
	_, err = GetLocation(database.DbConnection, location)
	if err != nil {
		return mouse, err
	}

	_, err = GetMouse(db, mouse)
	if err != nil {
		return mouse, err
	}

	sqlStatement := "UPDATE mouse SET name = $1, description = $2, min_power = $3, max_power = $4, gold = $5, location_id = $6, updated_at = $7 WHERE id = $8"
	rows := db.QueryRow(sqlStatement, mouse.Name, mouse.Description, mouse.MinPower, mouse.MaxPower, mouse.Gold, mouse.LocationID, time.Now(), mouse.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return mouse, err
	} else {
		return mouse, nil
	}
}

func DeleteMouse(db *sql.DB, mouse structs.Mouse) (structs.Mouse, error) {
	mouse, err = GetMouse(db, mouse)
	if err != nil {
		return mouse, err
	}
	sqlStatement := "DELETE FROM mouse WHERE id = $1"
	rows := db.QueryRow(sqlStatement, mouse.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return mouse, err
	} else {
		return mouse, nil
	}
}

//func GetMouseByLocation(db *sql.DB, cat structs.Mouse) (results []structs.Book, err error) {
//	sqlStatement := "SELECT * FROM book WHERE category_id = $1 ORDER BY id"
//
//	rows, err := db.Query(sqlStatement, cat.ID)
//	if err != nil {
//		return
//	}
//	defer func(rows *sql.Rows) {
//		err := rows.Close()
//		if err != nil {
//			return
//		}
//	}(rows)
//
//	for rows.Next() {
//		var book = structs.Book{}
//
//		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price,
//			&book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID)
//		if err != nil {
//			return
//		}
//		results = append(results, book)
//	}
//	return
//}