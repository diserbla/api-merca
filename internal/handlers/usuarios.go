package handlers

import (
	"net/http"

			"api-merca/internal/database"
	"api-merca/internal/models/usuarios"

	"github.com/gin-gonic/gin"
)

// GET /usuarios
func GetUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	if err := database.DB.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// GET /usuarios/:id
func GetUsuarioByID(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	if err := database.DB.First(&usuario, "cod_usu = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

// POST /usuarios
func CreateUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

// PUT /usuarios/:id
func UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	if err := database.DB.First(&usuario, "cod_usu = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// DELETE /usuarios/:id
func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Usuario{}, "cod_usu = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
