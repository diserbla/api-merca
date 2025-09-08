package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-merca/internal/config"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	log.Println("✅ Conectado a PostgreSQL con éxito")
}
