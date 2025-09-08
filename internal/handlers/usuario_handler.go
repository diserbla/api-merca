package handlers

import (
	models "api-merca/internal/models/usuarios"
	"api-merca/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UsuarioHandler struct {
	Service *services.UsuarioService
}

func (h *UsuarioHandler) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error al leer cuerpo de solicitud", http.StatusBadRequest)
		return
	}

	if err := h.Service.CrearUsuario(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (h *UsuarioHandler) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("cod_usu")
	codUsu, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "cod_usu inválido", http.StatusBadRequest)
		return
	}

	u, err := h.Service.ObtenerUsuarioPorID(codUsu)
	if err != nil {
		http.Error(w, "Error al obtener usuario", http.StatusInternalServerError)
		return
	}
	if u == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(u)
}
