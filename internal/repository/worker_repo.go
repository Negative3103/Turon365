package repository

import (
	"github.com/jmoiron/sqlx"

	"Turon365/internal/models"
)

type WorkerRepository struct {
	DB *sqlx.DB
}

func (r *WorkerRepository) Create(worker *models.Worker) error {
	query := `INSERT INTO workers (id, user_id, experience, is_approved, created_at) 
			  VALUES (:id, :user_id, :experience, :is_approved, :created_at)`
	_, err := r.DB.NamedExec(query, worker)
	return err
}

func (r *WorkerRepository) GetByID(id string) (*models.Worker, error) {
	var worker models.Worker
	query := `SELECT * FROM workers WHERE id = $1`
	err := r.DB.Get(&worker, query, id)
	if err != nil {
		return nil, err
	}
	return &worker, nil
}

func (r *WorkerRepository) Update(worker *models.Worker) error {
	query := `UPDATE workers SET experience=:experience, is_approved=:is_approved WHERE id=:id`
	_, err := r.DB.NamedExec(query, worker)
	return err
}

func (r *WorkerRepository) Delete(id string) error {
	query := `DELETE FROM workers WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
