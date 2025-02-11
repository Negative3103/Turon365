package repository

import (
    "Turon365/internal/models"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
    DB *sqlx.DB
}

func (r *CategoryRepository) Create(category *models.Category) error {
    query := `INSERT INTO categories (id, name, created_at) VALUES (:id, :name, :created_at)`
    _, err := r.DB.NamedExec(query, category)
    return err
}

func (r *CategoryRepository) GetByID(id uuid.UUID) (*models.Category, error) {
    var category models.Category
    query := `SELECT * FROM categories WHERE id = $1`
    err := r.DB.Get(&category, query, id)
    if err != nil {
        return nil, err
    }
    return &category, nil
}

func (r *CategoryRepository) Update(category *models.Category) error {
    query := `UPDATE categories SET name=:name WHERE id=:id`
    _, err := r.DB.NamedExec(query, category)
    return err
}

func (r *CategoryRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM categories WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}