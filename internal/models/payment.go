package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Payment struct {
	ID        uuid.UUID `db:"id"`
	JobID     uuid.UUID `db:"job_id"`
	Amount    float64   `db:"amount"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type PaymentRepository struct {
	DB *sqlx.DB
}

func (r *PaymentRepository) Create(payment *Payment) error {
	query := `INSERT INTO payments (id, job_id, amount, status, created_at) 
			  VALUES (:id, :job_id, :amount, :status, :created_at)`
	_, err := r.DB.NamedExec(query, payment)
	return err
}

func (r *PaymentRepository) GetByID(id uuid.UUID) (*Payment, error) {
	var payment Payment
	query := `SELECT * FROM payments WHERE id = $1`
	err := r.DB.Get(&payment, query, id)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) Update(payment *Payment) error {
	query := `UPDATE payments SET job_id=:job_id, amount=:amount, status=:status WHERE id=:id`
	_, err := r.DB.NamedExec(query, payment)
	return err
}

func (r *PaymentRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
