package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at"`
}
