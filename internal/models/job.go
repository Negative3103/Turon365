package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Job struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Photo       string    `db:"photo"`
	ClientID    uuid.UUID `db:"client_id"`
	WorkerID    uuid.UUID `db:"worker_id"`
	ServiceID   uuid.UUID `db:"service_id"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

type JobRepository struct {
	DB *sqlx.DB
}

func (r *JobRepository) Create(job *Job) error {
	query := `INSERT INTO jobs (id, title, description, photo, client_id, worker_id, service_id, status, created_at) 
			  VALUES (:id, :title, :description, :photo, :client_id, :worker_id, :service_id, :status, :created_at)`
	_, err := r.DB.NamedExec(query, job)
	return err
}

func (r *JobRepository) GetByID(id uuid.UUID) (*Job, error) {
	var job Job
	query := `SELECT * FROM jobs WHERE id = $1`
	err := r.DB.Get(&job, query, id)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepository) Update(job *Job) error {
	query := `UPDATE jobs SET title=:title, description=:description, photo=:photo, client_id=:client_id, 
			  worker_id=:worker_id, service_id=:service_id, status=:status WHERE id=:id`
	_, err := r.DB.NamedExec(query, job)
	return err
}

func (r *JobRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM jobs WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
