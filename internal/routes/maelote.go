package routes

import (
	"api-merca/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterMaeloteRoutes → rutas de Maelote
func RegisterMaeloteRoutes(r *gin.Engine) {
	maelote := r.Group("/maelote")
	{
		maelote.GET("/", handlers.GetMaeloteList)      // lista todos
		maelote.GET("/:id", handlers.GetMaeloteByID)   // obtiene uno por id
		maelote.POST("/", handlers.CreateMaelote)      // crea uno
		maelote.PUT("/:id", handlers.UpdateMaelote)    // actualiza uno
		maelote.DELETE("/:id", handlers.DeleteMaelote) // elimina uno
	}
}
