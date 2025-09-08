package repositories

import (
	"api-merca/internal/models"

	"gorm.io/gorm"
)

type MaeraspasRepository struct {
	DB *gorm.DB
}

func (r *MaeraspasRepository) GetAll() ([]models.Maeraspas, error) {
	var items []models.Maeraspas
	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MaeraspasRepository) GetByID(id int) (*models.Maeraspas, error) {
	var item models.Maeraspas
	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MaeraspasRepository) Create(item *models.Maeraspas) error {
	return r.DB.Create(item).Error
}

func (r *MaeraspasRepository) Update(item *models.Maeraspas) error {
	return r.DB.Save(item).Error
}

func (r *MaeraspasRepository) Delete(id int) error {
	return r.DB.Delete(&models.Maeraspas{}, id).Error
}
