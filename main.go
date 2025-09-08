package main

import (
	"api-merca/internal/config"
	"api-merca/internal/database"
	"api-merca/internal/routes"
	"fmt"
	"log"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Conectar DB
	database.ConnectDB(cfg)
	log.Println("✅ Conexión a la base de datos establecida")

	// Iniciar servidor
	r := routes.SetupRouter()

	// Puerto desde .env (ej: SERVER_PORT=8080)
	port := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Servidor escuchando en http://localhost%s\n", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("❌ Error al iniciar servidor: %v", err)
	}
}
