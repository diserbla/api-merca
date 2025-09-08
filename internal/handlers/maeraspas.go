package handlers

import (
	"net/http"

			"api-merca/internal/database"
	"api-merca/internal/models/maeraspas"

	"github.com/gin-gonic/gin"
)

// GET /maeraspas
func GetMaeraspas(c *gin.Context) {
	var maeraspas []models.Maeraspas
	if err := database.DB.Find(&maeraspas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, maeraspas)
}

// GET /maeraspas/:id
func GetMaeraspasByID(c *gin.Context) {
	id := c.Param("id")
	var maeraspas models.Maeraspas
	if err := database.DB.First(&maeraspas, "codigo = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maeraspas no encontrado"})
		return
	}
	c.JSON(http.StatusOK, maeraspas)
}

// POST /maeraspas
func CreateMaeraspas(c *gin.Context) {
	var maeraspas models.Maeraspas
	if err := c.ShouldBindJSON(&maeraspas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&maeraspas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, maeraspas)
}

// PUT /maeraspas/:id
func UpdateMaeraspas(c *gin.Context) {
	id := c.Param("id")
	var maeraspas models.Maeraspas
	if err := database.DB.First(&maeraspas, "codigo = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maeraspas no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&maeraspas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&maeraspas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, maeraspas)
}

// DELETE /maeraspas/:id
func DeleteMaeraspas(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Maeraspas{}, "codigo = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Maeraspas eliminado"})
}
