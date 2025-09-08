package services

import (
	models "api-merca/internal/models/usuarios"
	"api-merca/internal/repositories"
	"fmt"
	"time"
)

type UsuarioService struct {
	Repo *repositories.UsuarioRepository
}

func (s *UsuarioService) CrearUsuario(u *models.Usuario) error {
	if u.NomUsu == "" || u.ClaveUsu == "" || u.IdUsu == "" {
		return fmt.Errorf("nombre, clave e id_usu son obligatorios")
	}

	// Actualizar fecha de última modificación
	u.UsrUltModif = 1 // ejemplo: usuario que hace la modificación
	u.FecUltModif = time.Now()

	return s.Repo.Crear(u)
}

func (s *UsuarioService) ObtenerUsuarioPorID(codUsu int) (*models.Usuario, error) {
	return s.Repo.ObtenerPorID(codUsu)
}
