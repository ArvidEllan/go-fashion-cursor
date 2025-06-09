package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"virtual-tryon/config"
	"virtual-tryon/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// handleUploadPhoto processes the user's photo upload
func handleUploadPhoto(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get the file from the request
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Create uploads directory if it doesn't exist
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%s_%s%s", userID, uuid.New().String(), filepath.Ext(file.Filename))
	filepath := filepath.Join(uploadDir, filename)

	// Save the file
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create try-on record
	tryOn := models.TryOn{
		UserID:        uuid.MustParse(userID.(string)),
		OriginalImage: filepath,
		Status:        "pending",
	}

	if err := config.DB.Create(&tryOn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create try-on record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Photo uploaded successfully",
		"try_on_id": tryOn.ID,
	})
}

// handleProcessTryOn processes the virtual try-on request
func handleProcessTryOn(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		TryOnID   string `json:"try_on_id" binding:"required"`
		ProductID string `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get try-on record
	var tryOn models.TryOn
	if err := config.DB.Where("id = ? AND user_id = ?", req.TryOnID, userID).First(&tryOn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Try-on record not found"})
		return
	}

	// Update try-on record with product ID
	tryOn.ProductID = uuid.MustParse(req.ProductID)
	tryOn.Status = "processing"

	if err := config.DB.Save(&tryOn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update try-on record"})
		return
	}

	// TODO: Implement AI/ML processing here
	// For now, we'll simulate processing with a delay
	go func() {
		time.Sleep(5 * time.Second)
		tryOn.Status = "completed"
		tryOn.ResultImage = tryOn.OriginalImage // In reality, this would be the processed image
		config.DB.Save(&tryOn)
	}()

	c.JSON(http.StatusOK, gin.H{
		"message":   "Try-on processing started",
		"try_on_id": tryOn.ID,
	})
}

// handleGetTryOnHistory retrieves the user's try-on history
func handleGetTryOnHistory(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var tryOns []models.TryOn
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&tryOns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch try-on history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"try_ons": tryOns,
	})
}

// handleDeleteTryOnHistory deletes a try-on record
func handleDeleteTryOnHistory(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	tryOnID := c.Param("id")

	// Delete try-on record
	if err := config.DB.Where("id = ? AND user_id = ?", tryOnID, userID).Delete(&models.TryOn{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete try-on record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Try-on record deleted successfully",
	})
}
