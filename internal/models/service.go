package models

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Photo       string    `db:"photo"`
	WorkerID    uuid.UUID `db:"worker_id"`
	CategoryID  uuid.UUID `db:"category_id"`
	LocationID  uuid.UUID `db:"location_id"`
	Price       float64   `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
}
