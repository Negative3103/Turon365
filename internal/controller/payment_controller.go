package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type PaymentController struct {
    Repo *repository.PaymentRepository
}

func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    payment.ID = uuid.New()
    payment.CreatedAt = time.Now()
    if err := ctrl.Repo.Create(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create payment"})
        return
    }
    c.JSON(http.StatusCreated, payment)
}

func (ctrl *PaymentController) GetPayment(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
        return
    }

    payment, err := ctrl.Repo.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (ctrl *PaymentController) UpdatePayment(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
        return
    }

    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    payment.ID = id
    if err := ctrl.Repo.Update(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update payment"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (ctrl *PaymentController) DeletePayment(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
        return
    }

    if err := ctrl.Repo.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete payment"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}