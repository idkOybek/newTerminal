package repository

import (
	"context"
	"database/sql"

	"github.com/idkOybek/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := "SELECT id, inn, username, password, is_active, is_admin FROM users WHERE username=$1"
	row := r.db.QueryRowContext(ctx, query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.INN, &user.Username, &user.Password, &user.IsActive, &user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (inn, username, password, is_active, is_admin) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.INN, user.Username, user.Password, user.IsActive, user.IsAdmin).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := "SELECT id, inn, username, password, is_active, is_admin FROM users WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.INN, &user.Username, &user.Password, &user.IsActive, &user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET inn=$1, username=$2, password=$3, is_active=$4, is_admin=$5 WHERE id=$6"
	_, err := r.db.ExecContext(ctx, query, user.INN, user.Username, user.Password, user.IsActive, user.IsAdmin, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, inn, username, password, is_active, is_admin FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.INN, &user.Username, &user.Password, &user.IsActive, &user.IsAdmin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
