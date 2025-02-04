package repository

import (
	"Turon365/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	DB *sqlx.DB
}

func (r *ServiceRepository) Create(service *models.Service) error {
	query := `INSERT INTO services (id, title, description, photo, worker_id, category_id, location_id, price, created_at) 
			  VALUES (:id, :title, :description, :photo, :worker_id, :category_id, :location_id, :price, :created_at)`
	_, err := r.DB.NamedExec(query, service)
	return err
}

func (r *ServiceRepository) GetByID(id uuid.UUID) (*models.Service, error) {
	var service models.Service
	query := `SELECT * FROM services WHERE id = $1`
	err := r.DB.Get(&service, query, id)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) Update(service *models.Service) error {
	query := `UPDATE services SET title=:title, description=:description, photo=:photo, worker_id=:worker_id, 
			  category_id=:category_id, location_id=:location_id, price=:price WHERE id=:id`
	_, err := r.DB.NamedExec(query, service)
	return err
}

func (r *ServiceRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
