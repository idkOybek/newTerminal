package services

import (
	"context"

	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/repository"
)

type FiscalService struct {
	repo *repository.FiscalRepository
}

func NewFiscalService(repo *repository.FiscalRepository) *FiscalService {
	return &FiscalService{repo: repo}
}

func (s *FiscalService) GetAll(ctx context.Context) ([]models.FiscalModule, error) {
	return s.repo.GetAll(ctx)
}

func (s *FiscalService) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FiscalService) Create(ctx context.Context, module models.FiscalModule) error {
	return s.repo.Create(ctx, &module)
}

func (s *FiscalService) Update(ctx context.Context, module models.FiscalModule) error {
	return s.repo.Update(ctx, &module)
}

func (s *FiscalService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
