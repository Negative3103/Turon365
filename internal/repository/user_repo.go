package repository

import (
	"Turon365/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (id, name, phone, email, password_hash, rating, role, created_at)
			  VALUES (:id, :name, :phone, :email, :password_hash, :rating, :role, :created_at)`
	_, err := r.DB.NamedExec(query, user)
	return err
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET name=:name, phone=:phone, email=:email, password_hash=:password_hash,
			  rating=:rating, role=:role WHERE id=:id`
	_, err := r.DB.NamedExec(query, user)
	return err
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
