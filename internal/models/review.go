package models

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID `db:"id"`
	JobID     uuid.UUID `db:"job_id" json:"job_id"`
	Rating    int       `db:"rating"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}
