package repositories

import (
	"api-merca/internal/models"
	"database/sql"
	"fmt"
)

type UsuarioRepository struct {
	DB *sql.DB
}

func (r *UsuarioRepository) CrearUsuario(u *models.Usuario) error {
	query := `
	INSERT INTO usuarios (
		cod_usu, nom_usu, clave_usu, ced_usu, correo, tel_fij_usu, tel_cel_usu, 
		usu_cod_car, usu_pto_vta, usr_ult_modif, fec_ult_modif, id_usu, estado_usu, 
		cod_vendedor, fec_ingreso, fec_retiro
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
	`
	_, err := r.DB.Exec(query,
		u.CodUsu, u.NomUsu, u.ClaveUsu, u.CedUsu, u.Correo, u.TelFijUsu, u.TelCelUsu,
		u.UsuCodCar, u.UsuPtoVta, u.UsrUltModif, u.FecUltModif, u.IDUsu, u.EstadoUsu,
		u.CodVendedor, u.FecIngreso, u.FecRetiro,
	)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %w", err)
	}
	return nil
}

func (r *UsuarioRepository) ObtenerUsuarioPorID(codUsu int) (*models.Usuario, error) {
	query := `SELECT cod_usu, nom_usu, clave_usu, ced_usu, correo, tel_fij_usu, tel_cel_usu,
	          usu_cod_car, usu_pto_vta, usr_ult_modif, fec_ult_modif, id_usu, estado_usu,
	          cod_vendedor, fec_ingreso, fec_retiro
	          FROM usuarios WHERE cod_usu=$1`
	row := r.DB.QueryRow(query, codUsu)

	var u models.Usuario
	err := row.Scan(
		&u.CodUsu, &u.NomUsu, &u.ClaveUsu, &u.CedUsu, &u.Correo, &u.TelFijUsu, &u.TelCelUsu,
		&u.UsuCodCar, &u.UsuPtoVta, &u.UsrUltModif, &u.FecUltModif, &u.IDUsu, &u.EstadoUsu,
		&u.CodVendedor, &u.FecIngreso, &u.FecRetiro,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
