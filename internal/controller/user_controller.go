package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
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
    id := c.Param("id")
    user, err := ctrl.Repo.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    user.ID = uuid.MustParse(id)
    if err := ctrl.Repo.Update(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.Repo.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}