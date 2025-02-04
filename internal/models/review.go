package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Review struct {
	ID        uuid.UUID `db:"id"`
	JobID     uuid.UUID `db:"job_id"`
	Rating    int       `db:"rating"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

type ReviewRepository struct {
	DB *sqlx.DB
}

func (r *ReviewRepository) Create(review *Review) error {
	query := `INSERT INTO reviews (id, job_id, rating, comment, created_at) 
			  VALUES (:id, :job_id, :rating, :comment, :created_at)`
	_, err := r.DB.NamedExec(query, review)
	return err
}

func (r *ReviewRepository) GetByID(id uuid.UUID) (*Review, error) {
	var review Review
	query := `SELECT * FROM reviews WHERE id = $1`
	err := r.DB.Get(&review, query, id)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(review *Review) error {
	query := `UPDATE reviews SET job_id=:job_id, rating=:rating, comment=:comment WHERE id=:id`
	_, err := r.DB.NamedExec(query, review)
	return err
}

func (r *ReviewRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM reviews WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
