package controller

import (
    "Turon365/internal/models"
    "Turon365/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

type AdminController struct {
    WorkerRepo    *repository.WorkerRepository
    CategoryRepo  *repository.CategoryRepository
    LocationRepo  *repository.LocationRepository
    JobRepo       *repository.JobRepository
    PaymentRepo   *repository.PaymentRepository
    ReviewRepo    *repository.ReviewRepository
}

func (ctrl *AdminController) ConfirmWorker(c *gin.Context) {
    id := c.Param("id")
    workerID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker ID"})
        return
    }

    if err := ctrl.WorkerRepo.Confirm(workerID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not confirm worker"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Worker confirmed"})
}

func (ctrl *AdminController) AddCategory(c *gin.Context) {
    var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if err := ctrl.CategoryRepo.Create(&category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
        return
    }
    c.JSON(http.StatusCreated, category)
}

func (ctrl *AdminController) UpdateCategory(c *gin.Context) {
    id := c.Param("id")
    categoryID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    category.ID = categoryID
    if err := ctrl.CategoryRepo.Update(&category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update category"})
        return
    }
    c.JSON(http.StatusOK, category)
}

func (ctrl *AdminController) DeleteCategory(c *gin.Context) {
    id := c.Param("id")
    categoryID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    if err := ctrl.CategoryRepo.Delete(categoryID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete category"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (ctrl *AdminController) AddLocation(c *gin.Context) {
    var location models.Location
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if err := ctrl.LocationRepo.Create(&location); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create location"})
        return
    }
    c.JSON(http.StatusCreated, location)
}

func (ctrl *AdminController) UpdateLocation(c *gin.Context) {
    id := c.Param("id")
    locationID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
        return
    }

    var location models.Location
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    location.ID = locationID
    if err := ctrl.LocationRepo.Update(&location); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update location"})
        return
    }
    c.JSON(http.StatusOK, location)
}

func (ctrl *AdminController) DeleteLocation(c *gin.Context) {
    id := c.Param("id")
    locationID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
        return
    }

    if err := ctrl.LocationRepo.Delete(locationID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete location"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (ctrl *AdminController) GetAllJobs(c *gin.Context) {
    jobs, err := ctrl.JobRepo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve jobs"})
        return
    }
    c.JSON(http.StatusOK, jobs)
}

func (ctrl *AdminController) UpdateJobStatus(c *gin.Context) {
    id := c.Param("id")
    jobID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
        return
    }

    var status struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&status); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if err := ctrl.JobRepo.UpdateStatus(jobID, status.Status); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update job status"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Job status updated"})
}
