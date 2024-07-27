package services

import (
	"context"
	"log"

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
	log.Println("Service: Fetching all fiscal modules")
	modules, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Service: Error fetching all fiscal modules: %v", err)
		return nil, err
	}
	log.Println("Service: Successfully fetched all fiscal modules")
	return modules, nil
}

func (s *FiscalService) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	log.Printf("Service: Fetching fiscal module by ID: %d", id)
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("Service: Error fetching fiscal module by ID %d: %v", id, err)
		return nil, err
	}
	log.Printf("Service: Successfully fetched fiscal module by ID: %d", id)
	return module, nil
}

func (s *FiscalService) Create(ctx context.Context, module models.FiscalModule) error {
	log.Println("Service: Creating new fiscal module")
	err := s.repo.Create(ctx, &module)
	if err != nil {
		log.Printf("Service: Error creating fiscal module: %v", err)
		return err
	}
	log.Println("Service: Successfully created new fiscal module")
	return nil
}

func (s *FiscalService) Update(ctx context.Context, module models.FiscalModule) error {
	log.Printf("Service: Updating fiscal module with ID: %d", module.ID)
	err := s.repo.Update(ctx, &module)
	if err != nil {
		log.Printf("Service: Error updating fiscal module with ID %d: %v", module.ID, err)
		return err
	}
	log.Printf("Service: Successfully updated fiscal module with ID: %d", module.ID)
	return nil
}

func (s *FiscalService) Delete(ctx context.Context, id int) error {
	log.Printf("Service: Deleting fiscal module with ID: %d", id)
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Printf("Service: Error deleting fiscal module with ID %d: %v", id, err)
		return err
	}
	log.Printf("Service: Successfully deleted fiscal module with ID: %d", id)
	return nil
}
