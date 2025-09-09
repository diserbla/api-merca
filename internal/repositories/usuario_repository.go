package repositories

import (
	"api-merca/internal/models/usuarios"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func (r *UsuarioRepository) GetAll() ([]usuarios.Usuario, error) {
	var usuarios []usuarios.Usuario
	if err := r.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UsuarioRepository) GetByID(id int) (*usuarios.Usuario, error) {
	var usuario usuarios.Usuario
	if err := r.DB.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) Create(usuario *usuarios.Usuario) error {
	return r.DB.Create(usuario).Error
}

func (r *UsuarioRepository) Update(usuario *usuarios.Usuario) error {
	return r.DB.Save(usuario).Error
}

func (r *UsuarioRepository) Delete(id int) error {
	return r.DB.Delete(&usuarios.Usuario{}, id).Error
}