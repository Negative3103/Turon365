package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type WorkerController struct {
    Repo *repository.WorkerRepository
}

func (ctrl *WorkerController) RegisterWorker(c *gin.Context) {
    var worker models.Worker
    if err := c.ShouldBindJSON(&worker); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    worker.ID = uuid.New()
    worker.CreatedAt = time.Now()
    if err := ctrl.Repo.Create(&worker); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create worker"})
        return
    }
    c.JSON(http.StatusCreated, worker)
}

func (ctrl *WorkerController) GetWorker(c *gin.Context) {
    id := c.Param("id")
    workerID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker ID"})
        return
    }

    worker, err := ctrl.Repo.GetByID(workerID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
        return
    }
    c.JSON(http.StatusOK, worker)
}

func (ctrl *WorkerController) UpdateWorker(c *gin.Context) {
    id := c.Param("id")
    workerID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker ID"})
        return
    }

    var worker models.Worker
    if err := c.ShouldBindJSON(&worker); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    worker.ID = workerID
    if err := ctrl.Repo.Update(&worker); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update worker"})
        return
    }
    c.JSON(http.StatusOK, worker)
}

func (ctrl *WorkerController) DeleteWorker(c *gin.Context) {
    id := c.Param("id")
    workerID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker ID"})
        return
    }

    if err := ctrl.Repo.Delete(workerID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete worker"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Worker deleted"})
}

func (ctrl *WorkerController) ConfirmWorker(c *gin.Context) {
    id := c.Param("id")
    workerID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker ID"})
        return
    }

    if err := ctrl.Repo.Confirm(workerID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not confirm worker"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Worker confirmed"})
}