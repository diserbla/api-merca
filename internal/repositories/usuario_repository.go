package repositories

import (
	"api-merca/internal/models"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func (r *UsuarioRepository) GetAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := r.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UsuarioRepository) GetByID(id int) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := r.DB.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
	return r.DB.Create(usuario).Error
}

func (r *UsuarioRepository) Update(usuario *models.Usuario) error {
	return r.DB.Save(usuario).Error
}

func (r *UsuarioRepository) Delete(id int) error {
	return r.DB.Delete(&models.Usuario{}, id).Error
}
