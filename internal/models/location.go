package models

import (
	"github.com/google/uuid"
	"time"
)

type Location struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
