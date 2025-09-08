package services

import "api-merca/internal/models"

type MaeloteService struct {
	Repo MaeloteRepositoryInterface
}

type MaeloteRepositoryInterface interface {
	GetAll() ([]models.Maelote, error)
	GetByID(cod string) (*models.Maelote, error)
	Create(l *models.Maelote) error
	Update(l *models.Maelote) error
	Delete(cod string) error
}

func (s *MaeloteService) GetAllMaelotes() ([]models.Maelote, error) {
	return s.Repo.GetAll()
}

func (s *MaeloteService) GetMaeloteByID(cod string) (*models.Maelote, error) {
	return s.Repo.GetByID(cod)
}

func (s *MaeloteService) CreateMaelote(l *models.Maelote) error {
	return s.Repo.Create(l)
}

func (s *MaeloteService) UpdateMaelote(l *models.Maelote) error {
	return s.Repo.Update(l)
}

func (s *MaeloteService) DeleteMaelote(cod string) error {
	return s.Repo.Delete(cod)
}
