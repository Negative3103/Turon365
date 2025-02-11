package repository

import (
    "Turon365/internal/models"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
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

func (r *WorkerRepository) GetByID(id uuid.UUID) (*models.Worker, error) {
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

func (r *WorkerRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM workers WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}

func (r *WorkerRepository) Confirm(workerID uuid.UUID) error {
    query := "UPDATE workers SET confirmed = TRUE WHERE id = $1"
    _, err := r.DB.Exec(query, workerID)
    return err
}