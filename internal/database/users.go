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

func (m *UserModel) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	// Explicitly select columns to avoid scan mismatch errors
	query := `
        SELECT id, email, name, password 
        FROM users 
        WHERE id = $1
    `
	return m.getUser(query, id)
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	// Explicitly select columns to avoid scan mismatch errors
	query := `
        SELECT id, email, name, password 
        FROM users 
        WHERE id = $1
    `
	return m.getUser(query, email)
}
