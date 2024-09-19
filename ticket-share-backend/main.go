package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"ticket-share-backend/controllers"
	"ticket-share-backend/database"
	"ticket-share-backend/middlewares"
	"ticket-share-backend/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Invalid port number:", err)
	}

	// DB credentials from environment variables
	dbCredential := database.Credential{
		Host:         os.Getenv("DB_HOST"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         dbPort,
	}

    // Connect to the database
    database.Connect(dbCredential)

    // Drop tables if needed (during development)
    if err := database.DB.Migrator().DropTable("users", "events", "carts"); err != nil {
        log.Fatal("Failed to drop tables:", err)
    }

    // Migrate the models (User, Event, Ticket)
    if err := database.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Cart{}); err != nil {
        log.Fatal("Failed to migrate tables:", err)
    }

    log.Println("Database migration completed successfully")

	// Insert dummy event data
    dummyEvents := []models.Event{
        {
            Name:        "Music Concert",
            Description: "A live music concert featuring top artists.",
            Date:        time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
            Price:       50,
        },
        {
            Name:        "Tech Conference",
            Description: "A conference about the latest in tech.",
            Date:        time.Now().AddDate(0, 0, 14).Format("2006-01-02"),
            Price:       100,
        },
        {
            Name:        "Art Exhibition",
            Description: "An exhibition showcasing modern art.",
            Date:        time.Now().AddDate(0, 0, 21).Format("2006-01-02"),
            Price:       30,
        },
        {
            Name:        "Coding Bootcamp",
            Description: "A laravel workshop",
            Date:        time.Now().AddDate(0, 0, 21).Format("2006-01-02"),
            Price:       30,
        },
    }

	for _, event := range dummyEvents {
		database.DB.Create(&event)
	}

	log.Println("Dummy data inserted successfully")

    // Set up Gin router
    router := gin.Default()

    // Enable CORS for the frontend
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Health check endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK"})
    })

    // User authentication routes
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    // Event routes
    router.GET("/events", controllers.GetAllEvents)         // Get all events
    router.GET("/events/:id", controllers.GetEventByID)      // Get event by ID

    protected := router.Group("/")
    protected.Use(middlewares.AuthMiddleware()) // Apply AuthMiddleware here

    // Cart routes
    protected.GET("/cart/:user_id", controllers.GetAllCartItems)
    protected.POST("/cart", controllers.AddToCart)               // Add item to cart
    protected.PUT("/cart/:id", controllers.UpdateCartItemQuantity) // Update cart item quantity
    protected.DELETE("/cart/:id", controllers.DeleteCartItem)      // Delete cart item


    // Start server on port 8080
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
