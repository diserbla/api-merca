package maelote

import "time"

type Maelote struct {
	CodLot           string    `gorm:"column:cod_lot;primaryKey;size:3" json:"cod_lot"`
	NomLot           string    `gorm:"column:nom_lot;size:50;not null" json:"nom_lot"`
	CodEstLot        string    `gorm:"column:cod_est_lot;size:2;default:A;not null" json:"cod_est_lot"`
	UsrUltModif      string    `gorm:"column:usr_ult_modif;size:15;not null" json:"usr_ult_modif"`
	FecUltModif      time.Time `gorm:"column:fec_ult_modif;not null" json:"fec_ult_modif"`
	NombreCorto      *string   `gorm:"column:nombre_corto;size:30" json:"nombre_corto,omitempty"`
	CntCifrasPremio  int       `gorm:"column:cnt_cifras_premio;default:4;not null" json:"cnt_cifras_premio"`
	LinkResultado    *string   `gorm:"column:link_resultado;size:200" json:"link_resultado,omitempty"`
	NumSor           *string   `gorm:"column:num_sor;size:4;default:1" json:"num_sor,omitempty"`
	VlrFra           float64   `gorm:"column:vlr_fra;default:0" json:"vlr_fra"`
	NroFracciones    *int      `gorm:"column:nro_fracciones" json:"nro_fracciones,omitempty"`
	Linea            int       `gorm:"column:linea;default:0" json:"linea"`
	DiaJuegaSorteo   int       `gorm:"column:dia_juega_sorteo;default:1" json:"dia_juega_sorteo"`
	TipoDistribucion *string   `gorm:"column:tipo_distribucion;size:20;default:COOPERATIVA" json:"tipo_distribucion,omitempty"`
	CodProv          *int      `gorm:"column:cod_prov" json:"cod_prov,omitempty"`
	CodLotMae        *string   `gorm:"column:cod_lot_mae;size:3" json:"cod_lot_mae,omitempty"`
}

func (Maelote) TableName() string {
	return "maelote"
}