package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserModel struct {
	DB *sql.DB //DB field, which is a pointer to a sql.DB instance.
}

type User struct {
	Id       int    `json:"id"` // json here for ensuring proper data serialization and deserialization.
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

var ErrNotFound = errors.New("user not found")

func (m UserModel) Insert(user *User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `INSERT INTO users (email,password,name) VALUES ($1, $2, $3) RETURNING id`

	err := m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil

}

func (m *UserModel) Get(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Explicitly select columns to avoid scan mismatch errors
	query := `
        SELECT id, email, name, password 
        FROM users 
        WHERE id = $1
    `

	var user User

	// Execute the query
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err // unexpected DB error
	}

	return &user, nil
}
