package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID        uuid.UUID `db:"id"`
	JobID     uuid.UUID `db:"job_id"`
	Amount    float64   `db:"amount"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
