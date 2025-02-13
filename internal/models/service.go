package models

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Photo       string    `db:"photo" json:"photo"`
	WorkerID    uuid.UUID `db:"worker_id" json:"worker_id"`
	CategoryID  uuid.UUID `db:"category_id" json:"category_id"`
	LocationID  uuid.UUID `db:"location_id" json:"location_id"`
	Price       float64   `db:"price" json:"price"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}
