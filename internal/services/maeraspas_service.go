package services

import "api-merca/internal/models/maeraspas"

type MaeraspasService struct {
	Repo MaeraspasRepositoryInterface
}

type MaeraspasRepositoryInterface interface {
	GetAll() ([]maeraspas.Maeraspas, error)
	GetByID(id int) (*maeraspas.Maeraspas, error)
	Create(m *maeraspas.Maeraspas) error
	Update(m *maeraspas.Maeraspas) error
	Delete(id int) error
}

func (s *MaeraspasService) GetAllMaeraspas() ([]maeraspas.Maeraspas, error) {
	return s.Repo.GetAll()
}

func (s *MaeraspasService) GetMaeraspasByID(id int) (*maeraspas.Maeraspas, error) {
	return s.Repo.GetByID(id)
}

func (s *MaeraspasService) CreateMaeraspas(m *maeraspas.Maeraspas) error {
	return s.Repo.Create(m)
}

func (s *MaeraspasService) UpdateMaeraspas(m *maeraspas.Maeraspas) error {
	return s.Repo.Update(m)
}

func (s *MaeraspasService) DeleteMaeraspas(id int) error {
	return s.Repo.Delete(id)
}