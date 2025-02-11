package repository

import (
    "Turon365/internal/models"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
    DB *sqlx.DB
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
    query := `INSERT INTO payments (id, job_id, amount, status, created_at) 
              VALUES (:id, :job_id, :amount, :status, :created_at)`
    _, err := r.DB.NamedExec(query, payment)
    return err
}

func (r *PaymentRepository) GetByID(id uuid.UUID) (*models.Payment, error) {
    var payment models.Payment
    query := `SELECT * FROM payments WHERE id = $1`
    err := r.DB.Get(&payment, query, id)
    if err != nil {
        return nil, err
    }
    return &payment, nil
}

func (r *PaymentRepository) Update(payment *models.Payment) error {
    query := `UPDATE payments SET job_id=:job_id, amount=:amount, status=:status WHERE id=:id`
    _, err := r.DB.NamedExec(query, payment)
    return err
}

func (r *PaymentRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM payments WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}