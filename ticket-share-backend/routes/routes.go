package routes

import (
	"ticket-share-backend/controllers"
	"ticket-share-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/login", controllers.Login)
    router.POST("/register", controllers.Register)

    authorized := router.Group("/api")
    authorized.Use(middlewares.JwtAuthMiddleware())
    {
        authorized.POST("/receipts", controllers.CreateReceipt)
        authorized.GET("/receipts", controllers.GetReceipts)
        authorized.GET("/receipts/:id", controllers.GetReceiptByID)
        authorized.DELETE("/receipts/:id", controllers.DeleteReceipt)
    }
}
