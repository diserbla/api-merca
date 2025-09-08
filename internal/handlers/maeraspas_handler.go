package handlers

import (
	"api-merca/internal/models"
	"api-merca/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MaeraspasHandler struct {
	Service *services.MaeraspasService
}

func (h *MaeraspasHandler) CrearMaeraspas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Maeraspas
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateMaeraspas(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *MaeraspasHandler) ObtenerMaeraspas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := h.Service.GetAllMaeraspas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *MaeraspasHandler) ObtenerMaeraspasByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["codigo"])
	item, err := h.Service.GetMaeraspasByID(id)
	if err != nil {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *MaeraspasHandler) ActualizarMaeraspas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["codigo"])
	var item models.Maeraspas
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.Codigo = id
	if err := h.Service.UpdateMaeraspas(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *MaeraspasHandler) EliminarMaeraspas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["codigo"])
	if err := h.Service.DeleteMaeraspas(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Registro eliminado"})
}
