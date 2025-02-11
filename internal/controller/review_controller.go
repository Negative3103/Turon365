package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type ReviewController struct {
    Repo *repository.ReviewRepository
}

func (ctrl *ReviewController) CreateReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    review.ID = uuid.New()
    review.CreatedAt = time.Now()
    if err := ctrl.Repo.Create(&review); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create review"})
        return
    }
    c.JSON(http.StatusCreated, review)
}

func (ctrl *ReviewController) GetReview(c *gin.Context) {
    id := c.Param("id")
    reviewID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
        return
    }

    review, err := ctrl.Repo.GetByID(reviewID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
        return
    }
    c.JSON(http.StatusOK, review)
}

func (ctrl *ReviewController) UpdateReview(c *gin.Context) {
    id := c.Param("id")
    reviewID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
        return
    }

    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    review.ID = reviewID
    if err := ctrl.Repo.Update(&review); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update review"})
        return
    }
    c.JSON(http.StatusOK, review)
}

func (ctrl *ReviewController) DeleteReview(c *gin.Context) {
    id := c.Param("id")
    reviewID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
        return
    }

    if err := ctrl.Repo.Delete(reviewID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete review"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}