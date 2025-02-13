package controller

import (
	"Turon365/internal/models"
	"Turon365/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type JobController struct {
	Repo *repository.JobRepository
}

func (ctrl *JobController) CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	job.ID = uuid.New()
	job.CreatedAt = time.Now()
	if err := ctrl.Repo.Create(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
		fmt.Printf(err.Error())
		return
	}
	c.JSON(http.StatusCreated, job)
}

func (ctrl *JobController) GetJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := ctrl.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (ctrl *JobController) UpdateJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	job.ID = id
	if err := ctrl.Repo.Update(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update job"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (ctrl *JobController) DeleteJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	if err := ctrl.Repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete job"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
