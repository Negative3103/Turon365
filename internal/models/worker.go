package models

import (
	"github.com/google/uuid"
	"time"
)

type Worker struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	PhoneNumber string    `db:"phone_number"`
	Experience  string    `db:"experience"`
	IsApproved  bool      `db:"is_approved"`
	Confirmed   bool      `db:"confirmed"`
	CreatedAt   time.Time `db:"created_at"`
}
