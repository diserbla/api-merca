package handlers

import (
	"net/http"

	"api-merca/internal/database"
	models "api-merca/internal/models/maelote"

	"github.com/gin-gonic/gin"
)

// GET /maelote
func GetMaeloteList(c *gin.Context) {
	var maelotes []models.Maelote
	if err := database.DB.Find(&maelotes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, maelotes)
}

// GET /maelote/:id
func GetMaeloteByID(c *gin.Context) {
	id := c.Param("id")
	var maelote models.Maelote
	if err := database.DB.First(&maelote, "cod_lot = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maelote no encontrado"})
		return
	}
	c.JSON(http.StatusOK, maelote)
}

// POST /maelote
func CreateMaelote(c *gin.Context) {
	var maelote models.Maelote
	if err := c.ShouldBindJSON(&maelote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&maelote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, maelote)
}

// PUT /maelote/:id
func UpdateMaelote(c *gin.Context) {
	id := c.Param("id")
	var maelote models.Maelote
	if err := database.DB.First(&maelote, "cod_lot = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maelote no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&maelote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&maelote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, maelote)
}

// DELETE /maelote/:id
func DeleteMaelote(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Maelote{}, "cod_lot = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Maelote eliminado"})
}
