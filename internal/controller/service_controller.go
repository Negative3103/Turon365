package controller

import (
	"Turon365/internal/models"
	"Turon365/internal/repository"
	"Turon365/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ServiceController struct {
	Repo *repository.ServiceRepository
}

func (ctrl *ServiceController) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	service.ID = uuid.New()
	service.CreatedAt = time.Now()
	if err := ctrl.Repo.Create(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create service"})
		fmt.Printf(err.Error(), "123")
		return
	}
	c.JSON(http.StatusCreated, service)
}

func (ctrl *ServiceController) GetService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	service, err := ctrl.Repo.GetByID(serviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}
func (ctrl *ServiceController) UpdateService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	service.ID = serviceID
	if err := ctrl.Repo.Update(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update service"})
		fmt.Printf(err.Error(), "123")
		return
	}
	c.JSON(http.StatusOK, service)
}

func (ctrl *ServiceController) DeleteService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	if err := ctrl.Repo.Delete(serviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete service"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (ctrl *ServiceController) UploadServicePhoto(c *gin.Context) {
	serviceID := c.Param("id")
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	filePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	objectName := "services/" + serviceID + "/" + file.Filename
	if err := storage.UploadFile("photos", objectName, filePath, file.Header.Get("Content-Type")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
