package routes

import (
	"api-merca/internal/handlers"

	"github.com/gorilla/mux"
)

func UsuariosRoutes(r *mux.Router, h *handlers.UsuarioHandler) {
	r.HandleFunc("/usuarios", h.CrearUsuario).Methods("POST")
	r.HandleFunc("/usuarios", h.ObtenerUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{cod_usu}", h.ObtenerUsuario).Methods("GET")
	r.HandleFunc("/usuarios/{cod_usu}", h.ActualizarUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{cod_usu}", h.EliminarUsuario).Methods("DELETE")
}
