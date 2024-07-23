package services

import (
	"context"

	"github.com/idkOybek/internal/logger"
	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/repository"
)

type TerminalService struct {
	repo          *repository.TerminalRepository
	fiscalService *FiscalService
}

func NewTerminalService(repo *repository.TerminalRepository, fiscalService *FiscalService) *TerminalService {
	return &TerminalService{
		repo:          repo,
		fiscalService: fiscalService,
	}
}

func (s *TerminalService) GetAllTerminals(ctx context.Context) ([]models.Terminal, error) {
	terminals, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorLogger.Printf("Error retrieving terminals from repository: %v", err)
		return nil, err
	}
	return terminals, nil
}

func (s *TerminalService) GetTerminalByID(ctx context.Context, id int) (*models.Terminal, error) {
	terminal, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error retrieving terminal by ID from repository: %v", err)
		return nil, err
	}
	return terminal, nil
}

func (s *TerminalService) CreateTerminal(ctx context.Context, terminal *models.Terminal) error {
	if err := s.repo.Create(ctx, terminal); err != nil {
		logger.ErrorLogger.Printf("Error creating terminal in repository: %v", err)
		return err
	}
	return nil
}

func (s *TerminalService) UpdateTerminal(ctx context.Context, terminal *models.Terminal) error {
	if err := s.repo.Update(ctx, terminal); err != nil {
		logger.ErrorLogger.Printf("Error updating terminal in repository: %v", err)
		return err
	}
	return nil
}

func (s *TerminalService) DeleteTerminal(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorLogger.Printf("Error deleting terminal from repository: %v", err)
		return err
	}
	return nil
}
