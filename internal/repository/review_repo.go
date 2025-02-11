package repository

import (
    "Turon365/internal/models"
    "github.com/jmoiron/sqlx"
)

type ReviewRepository struct {
    DB *sqlx.DB
}

func (repo *ReviewRepository) Create(review *models.Review) error {
    query := `INSERT INTO reviews (id, job_id, rating, comment, created_at) 
              VALUES (:id, :job_id, :rating, :comment, :created_at)`
    _, err := repo.DB.NamedExec(query, review)
    return err
}

func (repo *ReviewRepository) GetByID(id string) (*models.Review, error) {
    var review models.Review
    query := `SELECT * FROM reviews WHERE id = $1`
    err := repo.DB.Get(&review, query, id)
    return &review, err
}

func (repo *ReviewRepository) Update(review *models.Review) error {
    query := `UPDATE reviews SET job_id = :job_id, rating = :rating, comment = :comment WHERE id = :id`
    _, err := repo.DB.NamedExec(query, review)
    return err
}

func (repo *ReviewRepository) Delete(id string) error {
    query := `DELETE FROM reviews WHERE id = $1`
    _, err := repo.DB.Exec(query, id)
    return err
}