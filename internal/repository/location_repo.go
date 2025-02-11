package repository

import (
    "Turon365/internal/models"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

type LocationRepository struct {
    DB *sqlx.DB
}

func (repo *LocationRepository) Create(location *models.Location) error {
    query := `INSERT INTO locations (id, name, created_at) VALUES (:id, :name, :created_at)`
    _, err := repo.DB.NamedExec(query, location)
    return err
}

func (repo *LocationRepository) GetByID(id uuid.UUID) (*models.Location, error) {
    var location models.Location
    query := `SELECT * FROM locations WHERE id = $1`
    err := repo.DB.Get(&location, query, id)
    return &location, err
}

func (repo *LocationRepository) Update(location *models.Location) error {
    query := `UPDATE locations SET name = :name WHERE id = :id`
    _, err := repo.DB.NamedExec(query, location)
    return err
}

func (repo *LocationRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM locations WHERE id = $1`
    _, err := repo.DB.Exec(query, id)
    return err
}