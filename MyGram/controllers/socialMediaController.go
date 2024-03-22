// controllers/socialMediaController.go
package controllers

import (
	"mygram/config"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSocialMedia - Create a new social media link
func CreateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&socialMedia)
	c.JSON(http.StatusCreated, gin.H{"data": socialMedia})
}

// GetSocialMedias - Get all social media links
func GetSocialMedias(c *gin.Context) {
	var socialMedias []models.SocialMedia
	config.DB.Find(&socialMedias)
	c.JSON(http.StatusOK, gin.H{"data": socialMedias})
}

// UpdateSocialMedia - Update a social media link
func UpdateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	if err := config.DB.Where("id = ?", c.Param("id")).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updateInput models.SocialMedia
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&socialMedia).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": socialMedia})
}

// DeleteSocialMedia - Delete a social media link
func DeleteSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	if err := config.DB.Where("id = ?", c.Param("id")).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&socialMedia)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
