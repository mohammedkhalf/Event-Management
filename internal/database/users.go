package database

import "database/sql"

type UserModel struct {
	DB *sql.DB //DB field, which is a pointer to a sql.DB instance.
}

type User struct {
	Id       int    `json:"id"` // json here for ensuring proper data serialization and deserialization.
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
