package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/structs"
	"strconv"
	"time"
)

type UserRepository interface {
	GetAllUsers() ([]structs.User, error)
	GetUser(user structs.User) (structs.User, error)
	InsertUser(user structs.User) (structs.User, error)
	UpdateUser(user structs.User) (structs.User, error)
	DeleteUser(user structs.User) (structs.User, error)
	GetUserByUsername(username string) (structs.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers() ([]structs.User, error) {
	var results []structs.User
	sqlStatement := `SELECT * FROM "user" ORDER BY id`

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
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Gold, &user.LocationID, &user.TrapID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, user)
	}
	return results, nil
}

func (r *userRepository) GetUser(user structs.User) (structs.User, error) {
	var result structs.User
	sqlCheck := `SELECT * FROM "user" WHERE id = $1`

	rows := r.db.QueryRow(sqlCheck, user.ID)

	err = rows.Scan(&result.ID, &result.Username, &result.Password, &result.Role, &result.Gold, &result.LocationID, &result.TrapID, &result.CreatedAt, &result.UpdatedAt)
	if result == (structs.User{}) {
		err = errors.New("user with id " + strconv.Itoa(int(user.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func (r *userRepository) GetUserByUsername(username string) (structs.User, error) {
	var result structs.User
	sqlCheck := `SELECT * FROM "user" WHERE username = $1`

	rows := r.db.QueryRow(sqlCheck, username)

	err = rows.Scan(&result.ID, &result.Username, &result.Password, &result.Role, &result.Gold, &result.LocationID, &result.TrapID, &result.CreatedAt, &result.UpdatedAt)
	if result == (structs.User{}) {
		err = errors.New("username " + username + " not found")
		return result, err
	}
	return result, nil
}

func (r *userRepository) InsertUser(user structs.User) (structs.User, error) {
	now := time.Now()
	sqlStatement := `INSERT INTO "user"(username, password, role, gold, location_id, trap_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	rows := r.db.QueryRow(sqlStatement, user.Username, user.Password, user.Role, user.Gold, user.LocationID, user.TrapID, now, now)
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

func (r *userRepository) UpdateUser(user structs.User) (structs.User, error) {
	_, err := r.GetUser(user)

	if err != nil {
		return user, err
	}
	sqlStatement := `UPDATE "user" SET username = $1, password = $2, role = $3, gold = $4, location_id = $5, trap_id = $6, updated_at = $7 WHERE id = $8`
	rows := r.db.QueryRow(sqlStatement, user.Username, user.Password, user.Role, user.Gold, user.LocationID, user.TrapID, time.Now(), user.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return user, err
	} else {
		return user, nil
	}
}

func (r *userRepository) DeleteUser(user structs.User) (structs.User, error) {
	user, err = r.GetUser(user)
	if err != nil {
		return user, err
	}
	sqlStatement := `DELETE FROM "user" WHERE id = $1`
	rows := r.db.QueryRow(sqlStatement, user.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return user, err
	} else {
		return user, nil
	}
}