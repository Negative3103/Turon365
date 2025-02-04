package controller

import (
	"Turon365/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"

	"Turon365/internal/repository"
)

type UserController struct {
	Repo *repository.UserRepository
}

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	if err := ctrl.Repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := ctrl.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
