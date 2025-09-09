package handlers

import (
	"api-merca/internal/models/maelote"
	"api-merca/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MaeloteHandler struct {
	Service *services.MaeloteService
}

func (h *MaeloteHandler) CrearMaelote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lote maelote.Maelote
	if err := json.NewDecoder(r.Body).Decode(&lote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateMaelote(&lote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(lote)
}

func (h *MaeloteHandler) ObtenerMaelotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lotes, err := h.Service.GetAllMaelotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(lotes)
}

func (h *MaeloteHandler) ObtenerMaelote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	cod := params["cod_lot"]
	lote, err := h.Service.GetMaeloteByID(cod)
	if err != nil {
		http.Error(w, "Maelote no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(lote)
}

func (h *MaeloteHandler) ActualizarMaelote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	cod := params["cod_lot"]
	var lote maelote.Maelote
	if err := json.NewDecoder(r.Body).Decode(&lote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lote.CodLot = cod
	if err := h.Service.UpdateMaelote(&lote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(lote)
}

func (h *MaeloteHandler) EliminarMaelote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	cod := params["cod_lot"]
	if err := h.Service.DeleteMaelote(cod); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Maelote eliminado"})
}