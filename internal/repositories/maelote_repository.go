package repositories

import (
	"api-merca/internal/models/maelote"

	"gorm.io/gorm"
)

type MaeloteRepository struct {
	DB *gorm.DB
}

func (r *MaeloteRepository) GetAll() ([]maelote.Maelote, error) {
	var lotes []maelote.Maelote
	if err := r.DB.Find(&lotes).Error; err != nil {
		return nil, err
	}
	return lotes, nil
}

func (r *MaeloteRepository) GetByID(cod string) (*maelote.Maelote, error) {
	var lote maelote.Maelote
	if err := r.DB.First(&lote, "cod_lot = ?", cod).Error; err != nil {
		return nil, err
	}
	return &lote, nil
}

func (r *MaeloteRepository) Create(lote *maelote.Maelote) error {
	return r.DB.Create(lote).Error
}

func (r *MaeloteRepository) Update(lote *maelote.Maelote) error {
	return r.DB.Save(lote).Error
}

func (r *MaeloteRepository) Delete(cod string) error {
	return r.DB.Delete(&maelote.Maelote{}, "cod_lot = ?", cod).Error
}