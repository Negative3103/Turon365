package controller

import (
	"Turon365/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CategoryController struct {
	Repo *repository.CategoryRepository
}

func (ctrl *CategoryController) GetCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := ctrl.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}
