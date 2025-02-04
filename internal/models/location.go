package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Location struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type LocationRepository struct {
	DB *sqlx.DB
}

func (r *LocationRepository) Create(location *Location) error {
	query := `INSERT INTO locations (id, name, created_at) VALUES (:id, :name, :created_at)`
	_, err := r.DB.NamedExec(query, location)
	return err
}

func (r *LocationRepository) GetByID(id uuid.UUID) (*Location, error) {
	var location Location
	query := `SELECT * FROM locations WHERE id = $1`
	err := r.DB.Get(&location, query, id)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepository) Update(location *Location) error {
	query := `UPDATE locations SET name=:name WHERE id=:id`
	_, err := r.DB.NamedExec(query, location)
	return err
}

func (r *LocationRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM locations WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
