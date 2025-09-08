package routes

import (
	"api-merca/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterMaeloteRoutes → rutas de Maelote
func RegisterMaeloteRoutes(r *gin.Engine) {
	maelotes := r.Group("/maelotes")
	{
		maelotes.GET("/", handlers.GetMaelotes)
		maelotes.GET("/:id", handlers.GetMaeloteByID)
		maelotes.POST("/", handlers.CreateMaelote)
		maelotes.PUT("/:id", handlers.UpdateMaelote)
		maelotes.DELETE("/:id", handlers.DeleteMaelote)
	}
}
