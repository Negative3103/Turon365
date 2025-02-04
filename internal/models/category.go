package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryRepository struct {
	DB *sqlx.DB
}

func (r *CategoryRepository) Create(category *Category) error {
	query := `INSERT INTO categories (id, name, created_at) VALUES (:id, :name, :created_at)`
	_, err := r.DB.NamedExec(query, category)
	return err
}

func (r *CategoryRepository) GetByID(id uuid.UUID) (*Category, error) {
	var category Category
	query := `SELECT * FROM categories WHERE id = $1`
	err := r.DB.Get(&category, query, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *Category) error {
	query := `UPDATE categories SET name=:name WHERE id=:id`
	_, err := r.DB.NamedExec(query, category)
	return err
}

func (r *CategoryRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
