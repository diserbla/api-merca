package main

import (
	"api-merca/internal/config"
	"api-merca/internal/database"
	"api-merca/internal/routes"
)

func main() {
	// cargar configuración
	cfg := config.LoadConfig()

	// conectar DB
	database.ConnectDB(cfg)

	// iniciar servidor
	r := routes.SetupRouter()
	r.Run(":8080")
}
