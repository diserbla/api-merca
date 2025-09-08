package routes

import (
	"api-merca/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterMaeraspasRoutes → rutas de Maeraspas
func RegisterMaeraspasRoutes(r *gin.Engine) {
	maeraspas := r.Group("/maeraspas")
	{
		maeraspas.GET("/", handlers.GetMaeraspas)
		maeraspas.GET("/:id", handlers.GetMaeraspasByID)
		maeraspas.POST("/", handlers.CreateMaeraspas)
		maeraspas.PUT("/:id", handlers.UpdateMaeraspas)
		maeraspas.DELETE("/:id", handlers.DeleteMaeraspas)
	}
}
