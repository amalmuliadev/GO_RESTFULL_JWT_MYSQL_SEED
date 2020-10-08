package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"

	"../security"
)

// User model
type User struct {
	ID        uint64    `gorm:"primary_id;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `json:"posts,omitempty"`
}

// BeforeSave hash the user password
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare cleans the inputs
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}
	default:
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
	}

	return nil
}
