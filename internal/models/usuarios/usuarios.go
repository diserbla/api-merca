package models

import "time"

type Usuario struct {
	CodUsu      int        `gorm:"column:cod_usu;primaryKey" json:"cod_usu"`
	NomUsu      string     `gorm:"column:nom_usu;size:50;not null" json:"nom_usu"`
	ClaveUsu    string     `gorm:"column:clave_usu;size:50;not null" json:"clave_usu"`
	CedUsu      *int       `gorm:"column:ced_usu" json:"ced_usu,omitempty"`
	Correo      *string    `gorm:"column:correo;size:30" json:"correo,omitempty"`
	TelFijUsu   *string    `gorm:"column:tel_fij_usu;size:15" json:"tel_fij_usu,omitempty"`
	TelCelUsu   *string    `gorm:"column:tel_cel_usu;size:15" json:"tel_cel_usu,omitempty"`
	UsuCodCar   *int       `gorm:"column:usu_cod_car" json:"usu_cod_car,omitempty"`
	UsuPtoVta   *int       `gorm:"column:usu_pto_vta" json:"usu_pto_vta,omitempty"`
	UsrUltModif int        `gorm:"column:usr_ult_modif;not null" json:"usr_ult_modif"`
	FecUltModif time.Time  `gorm:"column:fec_ult_modif;not null" json:"fec_ult_modif"`
	IdUsu       string     `gorm:"column:id_usu;size:15;unique;not null" json:"id_usu"`
	EstadoUsu   *string    `gorm:"column:estado_usu;size:1;default:A" json:"estado_usu,omitempty"`
	CodVendedor *int       `gorm:"column:cod_vendedor" json:"cod_vendedor,omitempty"`
	FecIngreso  *time.Time `gorm:"column:fec_ingreso" json:"fec_ingreso,omitempty"`
	FecRetiro   *time.Time `gorm:"column:fec_retiro" json:"fec_retiro,omitempty"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
