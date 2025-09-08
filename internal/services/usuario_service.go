package services

import "api-merca/internal/models"

type UsuarioService struct {
	Repo UsuarioRepositoryInterface
}

type UsuarioRepositoryInterface interface {
	GetAll() ([]models.Usuario, error)
	GetByID(id int) (*models.Usuario, error)
	Create(u *models.Usuario) error
	Update(u *models.Usuario) error
	Delete(id int) error
}

func (s *UsuarioService) GetAllUsuarios() ([]models.Usuario, error) {
	return s.Repo.GetAll()
}

func (s *UsuarioService) GetUsuarioByID(id int) (*models.Usuario, error) {
	return s.Repo.GetByID(id)
}

func (s *UsuarioService) CreateUsuario(u *models.Usuario) error {
	return s.Repo.Create(u)
}

func (s *UsuarioService) UpdateUsuario(u *models.Usuario) error {
	return s.Repo.Update(u)
}

func (s *UsuarioService) DeleteUsuario(id int) error {
	return s.Repo.Delete(id)
}
