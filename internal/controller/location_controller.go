package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type LocationController struct {
    Repo *repository.LocationRepository
}

func (ctrl *LocationController) CreateLocation(c *gin.Context) {
    var location models.Location
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    location.ID = uuid.New()
    location.CreatedAt = time.Now()
    if err := ctrl.Repo.Create(&location); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create location"})
        return
    }
    c.JSON(http.StatusCreated, location)
}

func (ctrl *LocationController) GetLocation(c *gin.Context) {
    id := c.Param("id")
    location, err := ctrl.Repo.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
        return
    }
    c.JSON(http.StatusOK, location)
}

func (ctrl *LocationController) UpdateLocation(c *gin.Context) {
    id := c.Param("id")
    var location models.Location
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    location.ID = uuid.MustParse(id)
    if err := ctrl.Repo.Update(&location); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update location"})
        return
    }
    c.JSON(http.StatusOK, location)
}

func (ctrl *LocationController) DeleteLocation(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.Repo.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete location"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Location deleted"})
}