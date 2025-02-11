package repository

import (
    "Turon365/internal/models"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

type JobRepository struct {
    DB *sqlx.DB
}

func (r *JobRepository) Create(job *models.Job) error {
    query := `INSERT INTO jobs (id, title, description, photo, client_id, worker_id, service_id, status, created_at) 
              VALUES (:id, :title, :description, :photo, :client_id, :worker_id, :service_id, :status, :created_at)`
    _, err := r.DB.NamedExec(query, job)
    return err
}

func (r *JobRepository) GetByID(id uuid.UUID) (*models.Job, error) {
    var job models.Job
    query := `SELECT * FROM jobs WHERE id = $1`
    err := r.DB.Get(&job, query, id)
    if err != nil {
        return nil, err
    }
    return &job, nil
}

func (r *JobRepository) GetAll() ([]models.Job, error) {
    var jobs []models.Job
    query := `SELECT * FROM jobs`
    err := r.DB.Select(&jobs, query)
    if err != nil {
        return nil, err
    }
    return jobs, nil
}

func (r *JobRepository) Update(job *models.Job) error {
    query := `UPDATE jobs SET title=:title, description=:description, photo=:photo, client_id=:client_id, 
              worker_id=:worker_id, service_id=:service_id, status=:status WHERE id=:id`
    _, err := r.DB.NamedExec(query, job)
    return err
}

func (r *JobRepository) UpdateStatus(id uuid.UUID, status string) error {
    query := `UPDATE jobs SET status=$1 WHERE id=$2`
    _, err := r.DB.Exec(query, status, id)
    return err
}

func (r *JobRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM jobs WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}