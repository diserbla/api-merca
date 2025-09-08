package services

import "api-merca/internal/models"

type MaeraspasService struct {
	Repo MaeraspasRepositoryInterface
}

type MaeraspasRepositoryInterface interface {
	GetAll() ([]models.Maeraspas, error)
	GetByID(id int) (*models.Maeraspas, error)
	Create(m *models.Maeraspas) error
	Update(m *models.Maeraspas) error
	Delete(id int) error
}

func (s *MaeraspasService) GetAllMaeraspas() ([]models.Maeraspas, error) {
	return s.Repo.GetAll()
}

func (s *MaeraspasService) GetMaeraspasByID(id int) (*models.Maeraspas, error) {
	return s.Repo.GetByID(id)
}

func (s *MaeraspasService) CreateMaeraspas(m *models.Maeraspas) error {
	return s.Repo.Create(m)
}

func (s *MaeraspasService) UpdateMaeraspas(m *models.Maeraspas) error {
	return s.Repo.Update(m)
}

func (s *MaeraspasService) DeleteMaeraspas(id int) error {
	return s.Repo.Delete(id)
}
