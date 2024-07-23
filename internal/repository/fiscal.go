package repository

import (
	"context"
	"database/sql"

	"github.com/idkOybek/internal/models"
)

type FiscalRepository struct {
	db *sql.DB
}

func NewFiscalRepository(db *sql.DB) *FiscalRepository {
	return &FiscalRepository{db: db}
}

func (r *FiscalRepository) GetAll(ctx context.Context) ([]models.FiscalModule, error) {
	query := "SELECT id, factory_number, fiscal_number, user_id FROM fiscal_modules"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []models.FiscalModule
	for rows.Next() {
		var module models.FiscalModule
		err := rows.Scan(&module.ID, &module.FactoryNumber, &module.FiscalNumber, &module.UserID)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (r *FiscalRepository) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	query := "SELECT id, factory_number, fiscal_number, user_id FROM fiscal_modules WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)

	var module models.FiscalModule
	err := row.Scan(&module.ID, &module.FactoryNumber, &module.FiscalNumber, &module.UserID)
	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *FiscalRepository) Create(ctx context.Context, module *models.FiscalModule) error {
	query := "INSERT INTO fiscal_modules (factory_number, fiscal_number, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, module.FactoryNumber, module.FiscalNumber, module.UserID).Scan(&module.ID)
	return err
}

func (r *FiscalRepository) Update(ctx context.Context, module *models.FiscalModule) error {
	query := "UPDATE fiscal_modules SET factory_number=$1, fiscal_number=$2, user_id=$3 WHERE id=$4"
	_, err := r.db.ExecContext(ctx, query, module.FactoryNumber, module.FiscalNumber, module.UserID, module.ID)
	return err
}

func (r *FiscalRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM fiscal_modules WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
