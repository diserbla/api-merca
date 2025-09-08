package repositories

import (
	models "api-merca/internal/models/usuarios"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func (r *UsuarioRepository) Crear(u *models.Usuario) error {
	result := r.DB.Create(u)
	return result.Error
}

func (r *UsuarioRepository) ObtenerPorID(codUsu int) (*models.Usuario, error) {
	var usuario models.Usuario
	result := r.DB.First(&usuario, codUsu)
	return &usuario, result.Error
}