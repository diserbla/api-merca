package models

import "time"

type Maeraspas struct {
	Codigo      int       `gorm:"column:codigo;primaryKey;autoIncrement" json:"codigo"`
	Nombre      string    `gorm:"column:nombre;size:50;not null" json:"nombre"`
	NombreCorto string    `gorm:"column:nombre_corto;size:30;not null" json:"nombre_corto"`
	Estado      string    `gorm:"column:estado;size:2;default:A;not null" json:"estado"`
	Valor       float64   `gorm:"column:valor;default:0;not null" json:"valor"`
	UsrUltModif string    `gorm:"column:usr_ult_modif;size:15;not null" json:"usr_ult_modif"`
	FecUltModif time.Time `gorm:"column:fec_ult_modif;not null" json:"fec_ult_modif"`
}

func (Maeraspas) TableName() string {
	return "maeraspas"
}
