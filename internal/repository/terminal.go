package repository

import (
	"context"
	"database/sql"

	"github.com/idkOybek/internal/models"
)

type TerminalRepository struct {
	db *sql.DB
}

func NewTerminalRepository(db *sql.DB) *TerminalRepository {
	return &TerminalRepository{db: db}
}

func (r *TerminalRepository) GetAll(ctx context.Context) ([]models.Terminal, error) {
	query := "SELECT id, inn, company_name, address, cash_register_number, module_number, assembly_number, last_request_date, database_update_date, status, user_id, created_at, updated_at, free_record_balance FROM terminals"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		err := rows.Scan(&terminal.ID, &terminal.INN, &terminal.CompanyName, &terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber, &terminal.AssemblyNumber, &terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.Status, &terminal.PartnerID)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}
	return terminals, nil
}

func (r *TerminalRepository) GetByID(ctx context.Context, id int) (*models.Terminal, error) {
	query := "SELECT id, inn, company_name, address, cash_register_number, module_number, assembly_number, last_request_date, database_update_date, status, user_id, free_record_balance FROM terminals WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)

	var terminal models.Terminal
	err := row.Scan(&terminal.ID, &terminal.INN, &terminal.CompanyName, &terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber, &terminal.AssemblyNumber, &terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.Status, &terminal.PartnerID)
	if err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (r *TerminalRepository) Create(ctx context.Context, terminal *models.Terminal) error {
	query := "INSERT INTO terminals (inn, company_name, address, cash_register_number, module_number, assembly_number, last_request_date, database_update_date, status, user_id, free_record_balance) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, terminal.INN, terminal.CompanyName, terminal.Address, terminal.CashRegisterNumber, terminal.ModuleNumber, terminal.AssemblyNumber, terminal.LastRequestDate, terminal.DatabaseUpdateDate, terminal.Status, terminal.PartnerID).Scan(&terminal.ID)
	return err
}

func (r *TerminalRepository) Update(ctx context.Context, terminal *models.Terminal) error {
	query := "UPDATE terminals SET inn=$1, company_name=$2, address=$3, cash_register_number=$4, module_number=$5, assembly_number=$6, last_request_date=$7, database_update_date=$8, status=$9, user_id=$10, free_record_balance=$11 WHERE id=$12"
	_, err := r.db.ExecContext(ctx, query, terminal.INN, terminal.CompanyName, terminal.Address, terminal.CashRegisterNumber, terminal.ModuleNumber, terminal.AssemblyNumber, terminal.LastRequestDate, terminal.DatabaseUpdateDate, terminal.Status, terminal.PartnerID, terminal.ID)
	return err
}

func (r *TerminalRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM terminals WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
