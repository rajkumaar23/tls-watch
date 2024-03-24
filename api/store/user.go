package store

import (
	"time"
)

type User struct {
	ID          uint64     `json:"id"`
	OIDCSubject string    `json:"oidc_subject"`
	Name        string    `json:"name"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateUser(user *User) error {
	_, err := DB.Exec("INSERT INTO users (oidc_subject, name, picture) VALUES (?, ?, ?)", user.OIDCSubject, user.Name, user.Picture)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByOIDCSubject(oidc_subject string) (*User, error) {
	var user User
	row := DB.QueryRow("SELECT * FROM users WHERE oidc_subject = ?", oidc_subject)
	if err := row.Scan(&user.ID, &user.OIDCSubject, &user.Name, &user.Picture, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
