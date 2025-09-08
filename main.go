package main

import (
	"api-merca/internal/init"
	"api-merca/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Inicializar handlers
	usuarioHandler := init.InitUsuarioHandler(db)
	maeloteHandler := init.InitMaeloteHandler(db)
	maeraspasHandler := init.InitMaeraspasHandler(db)

	// Rutas
	routes.UsuariosRoutes(r, usuarioHandler)
	routes.MaeloteRoutes(r, maeloteHandler)
	routes.MaeraspasRoutes(r, maeraspasHandler)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
