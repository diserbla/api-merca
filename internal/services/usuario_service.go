package services

import (
	"api-merca/internal/models"
	"api-merca/internal/repositories"
	"errors"
	"time"
)

type UsuarioService struct {
	Repo *repositories.UsuarioRepository
}

// CrearUsuario valida campos y luego lo crea
func (s *UsuarioService) CrearUsuario(u *models.Usuario) error {
	if u.NomUsu == "" || u.ClaveUsu == "" || u.IDUsu == "" {
		return errors.New("nombre, clave e id_usu son obligatorios")
	}

	// Asignar fecha de última modificación si no está
	if u.FecUltModif.IsZero() {
		u.FecUltModif = time.Now()
	}

	return s.Repo.CrearUsuario(u)
}

// ObtenerUsuarioPorID devuelve un usuario por cod_usu
func (s *UsuarioService) ObtenerUsuarioPorID(codUsu int) (*models.Usuario, error) {
	return s.Repo.ObtenerUsuarioPorID(codUsu)
}
