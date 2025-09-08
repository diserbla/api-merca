package repositories

import (
	"api-merca/internal/models"

	"gorm.io/gorm"
)

type MaeloteRepository struct {
	DB *gorm.DB
}

func (r *MaeloteRepository) GetAll() ([]models.Maelote, error) {
	var lotes []models.Maelote
	if err := r.DB.Find(&lotes).Error; err != nil {
		return nil, err
	}
	return lotes, nil
}

func (r *MaeloteRepository) GetByID(cod string) (*models.Maelote, error) {
	var lote models.Maelote
	if err := r.DB.First(&lote, "cod_lot = ?", cod).Error; err != nil {
		return nil, err
	}
	return &lote, nil
}

func (r *MaeloteRepository) Create(lote *models.Maelote) error {
	return r.DB.Create(lote).Error
}

func (r *MaeloteRepository) Update(lote *models.Maelote) error {
	return r.DB.Save(lote).Error
}

func (r *MaeloteRepository) Delete(cod string) error {
	return r.DB.Delete(&models.Maelote{}, "cod_lot = ?", cod).Error
}
