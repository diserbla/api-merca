package repositories

import (
	"api-merca/internal/models/maeraspas"

	"gorm.io/gorm"
)

type MaeraspasRepository struct {
	DB *gorm.DB
}

func (r *MaeraspasRepository) GetAll() ([]maeraspas.Maeraspas, error) {
	var items []maeraspas.Maeraspas
	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MaeraspasRepository) GetByID(id int) (*maeraspas.Maeraspas, error) {
	var item maeraspas.Maeraspas
	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MaeraspasRepository) Create(item *maeraspas.Maeraspas) error {
	return r.DB.Create(item).Error
}

func (r *MaeraspasRepository) Update(item *maeraspas.Maeraspas) error {
	return r.DB.Save(item).Error
}

func (r *MaeraspasRepository) Delete(id int) error {
	return r.DB.Delete(&maeraspas.Maeraspas{}, id).Error
}