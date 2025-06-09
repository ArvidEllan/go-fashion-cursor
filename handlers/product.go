package handlers

import (
	"net/http"

	"virtual-tryon/config"
	"virtual-tryon/models"

	"github.com/gin-gonic/gin"
)

// handleGetProducts retrieves all products with optional filtering
func handleGetProducts(c *gin.Context) {
	var products []models.Product
	query := config.DB.Model(&models.Product{})

	// Apply filters if provided
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if brand := c.Query("brand"); brand != "" {
		query = query.Where("brand = ?", brand)
	}

	// Execute query
	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// handleGetProduct retrieves a single product by ID
func handleGetProduct(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := config.DB.Preload("Sizes").First(&product, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

// handleAddToCart adds a product to the user's cart
func handleAddToCart(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		ProductID string `json:"product_id" binding:"required"`
		SizeID    string `json:"size_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify product exists
	var product models.Product
	if err := config.DB.First(&product, "id = ?", req.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Verify size exists
	var size models.Size
	if err := config.DB.First(&size, "id = ?", req.SizeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Size not found"})
		return
	}

	// TODO: Implement cart functionality
	// For now, we'll just return a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Product added to cart",
		"cart_item": gin.H{
			"product_id": req.ProductID,
			"size_id":    req.SizeID,
			"quantity":   req.Quantity,
		},
	})
}

// handleGetCart retrieves the user's cart
func handleGetCart(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// TODO: Implement cart retrieval
	// For now, we'll return an empty cart
	c.JSON(http.StatusOK, gin.H{
		"cart_items": []interface{}{},
	})
}
