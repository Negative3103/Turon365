package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Phone        string    `db:"phone"`
	Email        *string   `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Rating       float64   `db:"rating"`
	Role         string    `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}
