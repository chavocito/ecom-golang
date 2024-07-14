package user

import (
	"database/sql"
	"fmt"

	"github.com/chavocito/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddUser(regUser types.RegisterUserPayload) (*types.User, error) {
	q := "INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)"
	rows, err := s.db.Query(q, regUser.FirstName, regUser.LastName, regUser.Email, regUser.Password)

	if err != nil {
		return nil, err
	}
	user := new(types.User)

	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	q := "select * from users where email = ?;"
	rows, err := s.db.Query(q, email)
	if err != nil {
		return nil, err
	}

	if rows.NextResultSet() {
		return nil, fmt.Errorf("more than 1 user detected with the provided email %s", email)
	}
	user := new(types.User)
	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
