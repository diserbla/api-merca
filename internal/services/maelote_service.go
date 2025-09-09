package services

import "api-merca/internal/models/maelote"

type MaeloteService struct {
	Repo MaeloteRepositoryInterface
}

type MaeloteRepositoryInterface interface {
	GetAll() ([]maelote.Maelote, error)
	GetByID(cod string) (*maelote.Maelote, error)
	Create(l *maelote.Maelote) error
	Update(l *maelote.Maelote) error
	Delete(cod string) error
}

func (s *MaeloteService) GetAllMaelotes() ([]maelote.Maelote, error) {
	return s.Repo.GetAll()
}

func (s *MaeloteService) GetMaeloteByID(cod string) (*maelote.Maelote, error) {
	return s.Repo.GetByID(cod)
}

func (s *MaeloteService) CreateMaelote(l *maelote.Maelote) error {
	return s.Repo.Create(l)
}

func (s *MaeloteService) UpdateMaelote(l *maelote.Maelote) error {
	return s.Repo.Update(l)
}

func (s *MaeloteService) DeleteMaelote(cod string) error {
	return s.Repo.Delete(cod)
}