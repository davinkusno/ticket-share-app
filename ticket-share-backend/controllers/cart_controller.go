package controllers

import (
	"net/http"
	"ticket-share-backend/database"
	"ticket-share-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllCartItems(c *gin.Context) {
    userID := c.Param("user_id")
    var cartItems []models.Cart

    // Load cart items for the user, along with associated events
    if err := database.DB.Preload("Event").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
        return
    }

    // Calculate total price for each cart item (quantity * event price)
	totalPrice := 0

    for _, cartItem := range cartItems {
        cartItemPrice := float64(cartItem.Quantity) * cartItem.Event.Price
		totalPrice += int(cartItemPrice)
        // You can return this total price or use it for further processing
    }

    c.JSON(http.StatusOK, gin.H{"cartItems": cartItems, "totalPrice": totalPrice})
}


// AddToCart adds an item to the user's cart
func AddToCart(c *gin.Context) {
	var cart models.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to cart successfully", "cart": cart})
}

// UpdateCartItemQuantity updates the quantity of a cart item
func UpdateCartItemQuantity(c *gin.Context) {
	cartID := c.Param("id")
	var cart models.Cart

	if err := database.DB.First(&cart, cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	var input struct {
		Quantity int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart.Quantity = input.Quantity
	if err := database.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item updated successfully", "cart": cart})
}

// DeleteCartItem removes an item from the user's cart
func DeleteCartItem(c *gin.Context) {
	cartID := c.Param("id")

	if err := database.DB.Delete(&models.Cart{}, cartID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}
