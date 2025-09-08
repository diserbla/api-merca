package repositories

import (
	"api-merca/internal/models"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func (r *UsuarioRepository) Crear(u *models.Usuario) error {
	return r.DB.Create(u).Error
}

func (r *UsuarioRepository) ObtenerPorID(codUsu int) (*models.Usuario, error) {
	var u models.Usuario
	if err := r.DB.First(&u, codUsu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
