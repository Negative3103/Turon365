package controller

import (
	"Turon365/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type LocationController struct {
	Repo *repository.LocationRepository
}

func (ctrl *LocationController) GetLocation(c *gin.Context) {
	id := c.Param("id")
	locationID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	location, err := ctrl.Repo.GetByID(locationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}
	c.JSON(http.StatusOK, location)
}
