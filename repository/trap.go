package repository

import (
	"database/sql"
	"errors"
	"golang-demo-mousehunt/structs"
	"strconv"
	"time"
)

var (
	err error
)

type TrapRepository interface {
	GetAllTraps() ([]structs.Trap, error)
	GetTrap(trap structs.Trap) (structs.Trap, error)
	InsertTrap(trap structs.Trap) (structs.Trap, error)
	UpdateTrap(trap structs.Trap) (structs.Trap, error)
	DeleteTrap(trap structs.Trap) (structs.Trap, error)
}

type trapRepository struct {
	db *sql.DB
}

func NewTrapRepository(db *sql.DB) *trapRepository {
	return &trapRepository{db}
}

func (r *trapRepository) GetAllTraps() ([]structs.Trap, error) {
	var results []structs.Trap
	sqlStatement := "SELECT * FROM trap ORDER BY id"

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
		var trap = structs.Trap{}

		err = rows.Scan(&trap.ID, &trap.Name, &trap.Description, &trap.Power, &trap.Price, &trap.CreatedAt, &trap.UpdatedAt)
		if err != nil {
			return results, err
		}

		results = append(results, trap)
	}
	return results, nil
}

func (r *trapRepository) GetTrap(trap structs.Trap) (structs.Trap, error) {
	var result structs.Trap
	sqlCheck := "SELECT * FROM trap WHERE id = $1"

	rows := r.db.QueryRow(sqlCheck, trap.ID)

	err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.Power, &result.Price, &result.CreatedAt, &result.UpdatedAt)
	if result == (structs.Trap{}) {
		err = errors.New("trap with id " + strconv.Itoa(int(trap.ID)) + " not found")
		return result, err
	}
	return result, nil
}

func (r *trapRepository) InsertTrap(trap structs.Trap) (structs.Trap, error) {
	now := time.Now()
	sqlStatement := "INSERT INTO trap(name, description, power, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	rows := r.db.QueryRow(sqlStatement, trap.Name, trap.Description, trap.Power, trap.Price, now, now)
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

func (r *trapRepository) UpdateTrap(trap structs.Trap) (structs.Trap, error) {
	_, err = r.GetTrap(trap)
	if err != nil {
		return trap, err
	}
	sqlStatement := "UPDATE trap SET name = $1, description = $2, power = $3, price = $4, updated_at = $5 WHERE id = $6"
	rows := r.db.QueryRow(sqlStatement, trap.Name, trap.Description, trap.Power, trap.Price, time.Now(), trap.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return trap, err
	} else {
		return trap, nil
	}
}

func (r *trapRepository) DeleteTrap(trap structs.Trap) (structs.Trap, error) {
	trap, err = r.GetTrap(trap)
	if err != nil {
		return trap, err
	}
	sqlStatement := "DELETE FROM trap WHERE id = $1"
	rows := r.db.QueryRow(sqlStatement, trap.ID)
	if rows.Err() != nil {
		err = rows.Err()
		return trap, err
	} else {
		return trap, nil
	}
}

//func GetBooksByCategory(db *sql.DB, cat structs.Trap) (results []structs.Book, err error) {
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