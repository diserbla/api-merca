package routes

import (
	"api-merca/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUsuarioRoutes → rutas de Usuarios
func RegisterUsuarioRoutes(r *gin.Engine) {
	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("/", handlers.GetUsuarios)
		usuarios.GET("/:id", handlers.GetUsuarioByID)
		usuarios.POST("/", handlers.CreateUsuario)
		usuarios.PUT("/:id", handlers.UpdateUsuario)
		usuarios.DELETE("/:id", handlers.DeleteUsuario)
	}
}
