package routes

import (
	"api-merca/internal/handlers"

	"github.com/gorilla/mux"
)

func MaeloteRoutes(r *mux.Router, h *handlers.MaeloteHandler) {
	r.HandleFunc("/maelote", h.CrearMaelote).Methods("POST")
	r.HandleFunc("/maelote", h.ObtenerMaelotes).Methods("GET")
	r.HandleFunc("/maelote/{cod_lot}", h.ObtenerMaelote).Methods("GET")
	r.HandleFunc("/maelote/{cod_lot}", h.ActualizarMaelote).Methods("PUT")
	r.HandleFunc("/maelote/{cod_lot}", h.EliminarMaelote).Methods("DELETE")
}
