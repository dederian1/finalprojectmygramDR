// controllers/photoController.go
package controllers

import (
	"mygram/config"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePhoto - Create a new photo
func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&photo)
	c.JSON(http.StatusCreated, gin.H{"data": photo})
}

// GetPhotos - Get all photos
func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	config.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

// UpdatePhoto - Update a photo
func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updateInput models.Photo
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&photo).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

// DeletePhoto - Delete a photo
func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
