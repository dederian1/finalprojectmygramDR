// controllers/commentController.go
package controllers

import (
	"mygram/config"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateComment - Create a new comment
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&comment)
	c.JSON(http.StatusCreated, gin.H{"data": comment})
}

// GetComments - Get all comments
func GetComments(c *gin.Context) {
	var comments []models.Comment
	config.DB.Find(&comments)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateComment - Update a comment
func UpdateComment(c *gin.Context) {
	var comment models.Comment
	if err := config.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updateInput models.Comment
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&comment).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteComment - Delete a comment
func DeleteComment(c *gin.Context) {
	var comment models.Comment
	if err := config.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
