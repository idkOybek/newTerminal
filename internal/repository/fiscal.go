package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/idkOybek/internal/models"
)

type FiscalRepository struct {
	db *sql.DB
}

func NewFiscalRepository(db *sql.DB) *FiscalRepository {
	return &FiscalRepository{db: db}
}

func (r *FiscalRepository) GetAll(ctx context.Context) ([]models.FiscalModule, error) {
	log.Println("Repository: Fetching all fiscal modules")
	query := "SELECT id, factory_number, fiscal_number, user_id FROM fiscal_modules"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Repository: Error querying all fiscal modules: %v", err)
		return nil, err
	}
	defer rows.Close()

	var modules []models.FiscalModule
	for rows.Next() {
		var module models.FiscalModule
		err := rows.Scan(&module.ID, &module.FactoryNumber, &module.FiscalNumber, &module.UserID)
		if err != nil {
			log.Printf("Repository: Error scanning fiscal module row: %v", err)
			return nil, err
		}
		modules = append(modules, module)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Repository: Error iterating fiscal module rows: %v", err)
		return nil, err
	}
	log.Println("Repository: Successfully fetched all fiscal modules")
	return modules, nil
}

func (r *FiscalRepository) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	log.Printf("Repository: Fetching fiscal module by ID: %d", id)
	query := "SELECT id, factory_number, fiscal_number, user_id FROM fiscal_modules WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)

	var module models.FiscalModule
	err := row.Scan(&module.ID, &module.FactoryNumber, &module.FiscalNumber, &module.UserID)
	if err != nil {
		log.Printf("Repository: Error scanning fiscal module by ID %d: %v", id, err)
		return nil, err
	}

	log.Printf("Repository: Successfully fetched fiscal module by ID: %d", id)
	return &module, nil
}

func (r *FiscalRepository) Create(ctx context.Context, module *models.FiscalModule) error {
	log.Println("Repository: Creating new fiscal module")
	query := "INSERT INTO fiscal_modules (factory_number, fiscal_number, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, module.FactoryNumber, module.FiscalNumber, module.UserID).Scan(&module.ID)
	if err != nil {
		log.Printf("Repository: Error creating fiscal module: %v", err)
		return err
	}
	log.Println("Repository: Successfully created new fiscal module")
	return nil
}

func (r *FiscalRepository) Update(ctx context.Context, module *models.FiscalModule) error {
	log.Printf("Repository: Updating fiscal module with ID: %d", module.ID)
	query := "UPDATE fiscal_modules SET factory_number=$1, fiscal_number=$2, user_id=$3 WHERE id=$4"
	_, err := r.db.ExecContext(ctx, query, module.FactoryNumber, module.FiscalNumber, module.UserID, module.ID)
	if err != nil {
		log.Printf("Repository: Error updating fiscal module with ID %d: %v", module.ID, err)
		return err
	}
	log.Printf("Repository: Successfully updated fiscal module with ID: %d", module.ID)
	return nil
}

func (r *FiscalRepository) Delete(ctx context.Context, id int) error {
	log.Printf("Repository: Deleting fiscal module with ID: %d", id)
	query := "DELETE FROM fiscal_modules WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Repository: Error deleting fiscal module with ID %d: %v", id, err)
		return err
	}
	log.Printf("Repository: Successfully deleted fiscal module with ID: %d", id)
	return nil
}
