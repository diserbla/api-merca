package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes → función central para registrar todas las rutas
func RegisterRoutes(r *gin.Engine) {
	RegisterUsuarioRoutes(r)
	RegisterMaeloteRoutes(r)
	RegisterMaeraspasRoutes(r)
}

// SetupRouter → crea el router y registra todas las rutas
func SetupRouter() *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r)
	return r
}
