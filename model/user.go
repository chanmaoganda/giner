package model

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func QueryByName(db *sql.DB, name string) *User {
	var user User
	err := db.QueryRow("SELECT username, email FROM users WHERE username = $1", name).Scan(&user)
	if err != nil {
		logrus.Debug("Scan error")
	}

	return &user
}
