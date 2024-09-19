package controllers

import (
	"log"
	"net/http"
	"ticket-share-backend/database"
	"ticket-share-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
    var event models.Event
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&event)
    c.JSON(http.StatusOK, event)
}

func GetAllEvents(c *gin.Context) {
    log.Println("GetAllEvents endpoint hit")
    var events []models.Event
    database.DB.Find(&events)
    c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
    id := c.Param("id")
    var event models.Event
    if err := database.DB.First(&event, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
        return
    }
    c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
    id := c.Param("id")
    var event models.Event
    if err := database.DB.First(&event, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
        return
    }

    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&event)
    c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Event{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

