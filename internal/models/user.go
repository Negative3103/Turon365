package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Phone        string    `db:"phone"`
	Email        *string   `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Rating       float64   `db:"rating"`
	Role         string    `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) Create(user *User) error {
	query := `INSERT INTO users (id, name, phone, email, password_hash, rating, role, created_at) 
			  VALUES (:id, :name, :phone, :email, :password_hash, :rating, :role, :created_at)`
	_, err := r.DB.NamedExec(query, user)
	return err
}

func (r *UserRepository) GetByID(id uuid.UUID) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *User) error {
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
