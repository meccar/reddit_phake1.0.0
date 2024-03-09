// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Userrole string

const (
	UserroleGuest Userrole = "guest"
	UserroleUser  Userrole = "user"
	UserroleAdmin Userrole = "admin"
)

func (e *Userrole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Userrole(s)
	case string:
		*e = Userrole(s)
	default:
		return fmt.Errorf("unsupported scan type for Userrole: %T", src)
	}
	return nil
}

type NullUserrole struct {
	Userrole Userrole `json:"userrole"`
	Valid    bool     `json:"valid"` // Valid is true if Userrole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserrole) Scan(value interface{}) error {
	if value == nil {
		ns.Userrole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Userrole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserrole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Userrole), nil
}

type Account struct {
	ID              uuid.UUID `json:"id"`
	Role            Userrole  `json:"role"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
}

type Form struct {
	ID         uuid.UUID `json:"id"`
	ViewerName string    `json:"viewer_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
}

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Article   string    `json:"article"`
	Picture   []byte    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID        uuid.UUID          `json:"id"`
	Username  string             `json:"username"`
	Role      string             `json:"role"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type VerifyEmail struct {
	Username   string    `json:"username"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}