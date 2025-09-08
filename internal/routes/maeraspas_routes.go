package routes

import (
	"api-merca/internal/handlers"

	"github.com/gorilla/mux"
)

func MaeraspasRoutes(r *mux.Router, h *handlers.MaeraspasHandler) {
	r.HandleFunc("/maeraspas", h.CrearMaeraspas).Methods("POST")
	r.HandleFunc("/maeraspas", h.ObtenerMaeraspas).Methods("GET")
	r.HandleFunc("/maeraspas/{codigo}", h.ObtenerMaeraspasByID).Methods("GET")
	r.HandleFunc("/maeraspas/{codigo}", h.ActualizarMaeraspas).Methods("PUT")
	r.HandleFunc("/maeraspas/{codigo}", h.EliminarMaeraspas).Methods("DELETE")
}
