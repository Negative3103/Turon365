package models

import (
	"github.com/google/uuid"
	"time"
)

type Worker struct {
	ID         uuid.UUID `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	Experience string    `db:"experience"`
	IsApproved bool      `db:"is_approved"`
	Confirmed  bool      `db:"confirmed"`
	CreatedAt  time.Time `db:"created_at"`
}
