package main

import (
	"api-merca/internal/handlers"
	"api-merca/internal/repositories"
	"api-merca/internal/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables de entorno del sistema")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	repo := &repositories.UsuarioRepository{DB: db}
	service := &services.UsuarioService{Repo: repo}
	handler := &handlers.UsuarioHandler{Service: service}

	http.HandleFunc("/usuarios", handler.CrearUsuario)
	http.HandleFunc("/usuarios/get", handler.ObtenerUsuario)

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
