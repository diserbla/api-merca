package services

import "api-merca/internal/models/usuarios"

type UsuarioService struct {
	Repo UsuarioRepositoryInterface
}

type UsuarioRepositoryInterface interface {
	GetAll() ([]usuarios.Usuario, error)
	GetByID(id int) (*usuarios.Usuario, error)
	Create(u *usuarios.Usuario) error
	Update(u *usuarios.Usuario) error
	Delete(id int) error
}

func (s *UsuarioService) GetAllUsuarios() ([]usuarios.Usuario, error) {
	return s.Repo.GetAll()
}

func (s *UsuarioService) GetUsuarioByID(id int) (*usuarios.Usuario, error) {
	return s.Repo.GetByID(id)
}

func (s *UsuarioService) CreateUsuario(u *usuarios.Usuario) error {
	return s.Repo.Create(u)
}

func (s *UsuarioService) UpdateUsuario(u *usuarios.Usuario) error {
	return s.Repo.Update(u)
}

func (s *UsuarioService) DeleteUsuario(id int) error {
	return s.Repo.Delete(id)
}