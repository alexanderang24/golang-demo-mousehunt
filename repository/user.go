package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/dto"
	"strconv"
	"time"
)

func GetAllUsers(db *sql.DB) ([]dto.User, error) {
	var results []dto.User
	sqlStatement := `SELECT * FROM "user" ORDER BY id`

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
		var user = dto.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Gold, &user.LocationID, &user.TrapID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, user)
	}
	return results, nil
}

func GetUser(db *sql.DB, user dto.User) (dto.User, error) {
	var result dto.User
	sqlCheck := `SELECT * FROM "user" WHERE id = $1`

	rows := db.QueryRow(sqlCheck, user.ID)

	err = rows.Scan(&result.ID, &result.Username, &result.Password, &result.Role, &result.Gold, &result.LocationID, &result.TrapID, &result.CreatedAt, &result.UpdatedAt)
	if result == (dto.User{}) {
		err = errors.New("user with id " + strconv.Itoa(int(user.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func InsertUser(db *sql.DB, user dto.User) (dto.User, error) {
	now := time.Now()
	sqlStatement := `INSERT INTO "user"(username, password, role, gold, location_id, trap_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	rows := db.QueryRow(sqlStatement, user.Username, user.Password, user.Role, user.Gold, user.LocationID, user.TrapID, now, now)
	if rows.Err() != nil {
		err = rows.Err()
		return user, err
	} else {
		err = rows.Scan(&user.ID)
		if err != nil {
			return user, err
		}
		return user, nil
	}
}

func UpdateUser(db *sql.DB, user dto.User) (dto.User, error) {
	_, err := GetUser(db, user)

	if err != nil {
		return user, err
	}
	sqlStatement := `UPDATE "user" SET username = $1, password = $2, role = $3, gold = $4, location_id = $5, trap_id = $6, updated_at = $7 WHERE id = $8`
	rows := db.QueryRow(sqlStatement, user.Username, user.Password, user.Role, user.Gold, user.LocationID, user.TrapID, time.Now(), user.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return user, err
	} else {
		return user, nil
	}
}

func DeleteUser(db *sql.DB, user dto.User) (dto.User, error) {
	user, err = GetUser(db, user)
	if err != nil {
		return user, err
	}
	sqlStatement := `DELETE FROM "user" WHERE id = $1`
	rows := db.QueryRow(sqlStatement, user.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return user, err
	} else {
		return user, nil
	}
}

func GetUserByUsername(db *sql.DB, username string) (dto.User, error) {
	var result dto.User
	sqlCheck := `SELECT * FROM "user" WHERE username = $1`
	rows := db.QueryRow(sqlCheck, username)

	err = rows.Scan(&result.ID, &result.Username, &result.Password, &result.Role, &result.Gold, &result.LocationID, &result.TrapID, &result.CreatedAt, &result.UpdatedAt)
	return result, nil
}