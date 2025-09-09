package setup

import (
	"api-merca/internal/handlers"
	"api-merca/internal/repositories"
	"api-merca/internal/services"

	"gorm.io/gorm"
)

func InitUsuarioHandler(db *gorm.DB) *handlers.UsuarioHandler {
	repo := &repositories.UsuarioRepository{DB: db}
	service := &services.UsuarioService{Repo: repo}
	return &handlers.UsuarioHandler{Service: service}
}

func InitMaeloteHandler(db *gorm.DB) *handlers.MaeloteHandler {
	repo := &repositories.MaeloteRepository{DB: db}
	service := &services.MaeloteService{Repo: repo}
	return &handlers.MaeloteHandler{Service: service}
}

func InitMaeraspasHandler(db *gorm.DB) *handlers.MaeraspasHandler {
	repo := &repositories.MaeraspasRepository{DB: db}
	service := &services.MaeraspasService{Repo: repo}
	return &handlers.MaeraspasHandler{Service: service}
}