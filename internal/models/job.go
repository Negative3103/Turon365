package models

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Photo       string    `db:"photo"`
	ClientID    uuid.UUID `db:"client_id"`
	WorkerID    uuid.UUID `db:"worker_id"`
	ServiceID   uuid.UUID `db:"service_id"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}
