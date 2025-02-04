package repository

import (
	"Turon365/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LocationRepository struct {
	DB *sqlx.DB
}

func (r *LocationRepository) Create(location *models.Location) error {
	query := `INSERT INTO locations (id, name, created_at) VALUES (:id, :name, :created_at)`
	_, err := r.DB.NamedExec(query, location)
	return err
}

func (r *LocationRepository) GetByID(id uuid.UUID) (*models.Location, error) {
	var location models.Location
	query := `SELECT * FROM locations WHERE id = $1`
	err := r.DB.Get(&location, query, id)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepository) Update(location *models.Location) error {
	query := `UPDATE locations SET name=:name WHERE id=:id`
	_, err := r.DB.NamedExec(query, location)
	return err
}

func (r *LocationRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM locations WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
