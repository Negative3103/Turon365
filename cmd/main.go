package main

import (
	"Turon365/internal/controller"
	"Turon365/internal/middleware"
	"Turon365/internal/repository"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sqlx.DB
var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	dsn := "host=localhost port=54321 user=note-user password=note-password dbname=note sslmode=disable"
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Run migrations
	if err := runMigrations(); err != nil {
		logger.Fatal("Failed to run migrations", zap.Error(err))
	}

	// Initialize Minio
	//storage.InitMinio()
}

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.LoggingMiddleware(logger))

	// Repositories
	userRepo := &repository.UserRepository{DB: db}
	workerRepo := &repository.WorkerRepository{DB: db}
	categoryRepo := &repository.CategoryRepository{DB: db}
	locationRepo := &repository.LocationRepository{DB: db}
	serviceRepo := &repository.ServiceRepository{DB: db}
	jobRepo := &repository.JobRepository{DB: db}
	paymentRepo := &repository.PaymentRepository{DB: db}
	reviewRepo := &repository.ReviewRepository{DB: db}

	// Controllers
	userCtrl := &controller.UserController{Repo: userRepo}
	workerCtrl := &controller.WorkerController{Repo: workerRepo}
	categoryCtrl := &controller.CategoryController{Repo: categoryRepo}
	locationCtrl := &controller.LocationController{Repo: locationRepo}
	serviceCtrl := &controller.ServiceController{Repo: serviceRepo}
	jobCtrl := &controller.JobController{Repo: jobRepo}
	paymentCtrl := &controller.PaymentController{Repo: paymentRepo}
	reviewCtrl := &controller.ReviewController{Repo: reviewRepo}
	adminCtrl := &controller.AdminController{
		WorkerRepo:   workerRepo,
		CategoryRepo: categoryRepo,
		LocationRepo: locationRepo,
		JobRepo:      jobRepo,
		PaymentRepo:  paymentRepo,
		ReviewRepo:   reviewRepo,
	}

	// Routes
	r.POST("/users", userCtrl.RegisterUser)
	r.GET("/users/:id", middleware.AuthMiddleware(), userCtrl.GetUser)
	r.PUT("/users/:id", middleware.AuthMiddleware(), userCtrl.UpdateUser)
	r.DELETE("/users/:id", middleware.AuthMiddleware(), userCtrl.DeleteUser)

	r.POST("/workers", workerCtrl.RegisterWorker)
	r.GET("/workers/:id", middleware.AuthMiddleware(), workerCtrl.GetWorker)
	r.PUT("/workers/:id", middleware.AuthMiddleware(), workerCtrl.UpdateWorker)
	r.DELETE("/workers/:id", middleware.AuthMiddleware(), workerCtrl.DeleteWorker)

	r.POST("/categories", middleware.AuthMiddleware(), categoryCtrl.CreateCategory)
	r.GET("/categories/:id", middleware.AuthMiddleware(), categoryCtrl.GetCategory)
	r.PUT("/categories/:id", middleware.AuthMiddleware(), categoryCtrl.UpdateCategory)
	r.DELETE("/categories/:id", middleware.AuthMiddleware(), categoryCtrl.DeleteCategory)

	r.POST("/locations", middleware.AuthMiddleware(), locationCtrl.CreateLocation)
	r.GET("/locations/:id", middleware.AuthMiddleware(), locationCtrl.GetLocation)
	r.PUT("/locations/:id", middleware.AuthMiddleware(), locationCtrl.UpdateLocation)
	r.DELETE("/locations/:id", middleware.AuthMiddleware(), locationCtrl.DeleteLocation)

	r.POST("/services", middleware.AuthMiddleware(), serviceCtrl.CreateService)
	r.GET("/services/:id", middleware.AuthMiddleware(), serviceCtrl.GetService)
	r.PUT("/services/:id", middleware.AuthMiddleware(), serviceCtrl.UpdateService)
	r.DELETE("/services/:id", middleware.AuthMiddleware(), serviceCtrl.DeleteService)
	r.POST("/services/:id/photo", middleware.AuthMiddleware(), serviceCtrl.UploadServicePhoto)

	r.POST("/jobs", middleware.AuthMiddleware(), jobCtrl.CreateJob)
	r.GET("/jobs/:id", middleware.AuthMiddleware(), jobCtrl.GetJob)
	r.PUT("/jobs/:id", middleware.AuthMiddleware(), jobCtrl.UpdateJob)
	r.DELETE("/jobs/:id", middleware.AuthMiddleware(), jobCtrl.DeleteJob)

	r.POST("/payments", middleware.AuthMiddleware(), paymentCtrl.CreatePayment)
	r.GET("/payments/:id", middleware.AuthMiddleware(), paymentCtrl.GetPayment)
	r.PUT("/payments/:id", middleware.AuthMiddleware(), paymentCtrl.UpdatePayment)
	r.DELETE("/payments/:id", middleware.AuthMiddleware(), paymentCtrl.DeletePayment)

	r.POST("/reviews", middleware.AuthMiddleware(), reviewCtrl.CreateReview)
	r.GET("/reviews/:id", middleware.AuthMiddleware(), reviewCtrl.GetReview)
	r.PUT("/reviews/:id", middleware.AuthMiddleware(), reviewCtrl.UpdateReview)
	r.DELETE("/reviews/:id", middleware.AuthMiddleware(), reviewCtrl.DeleteReview)

	// Admin Routes
	r.POST("/admin/workers/:id/confirm", middleware.AuthMiddleware(), adminCtrl.ConfirmWorker)
	r.POST("/admin/categories", middleware.AuthMiddleware(), adminCtrl.AddCategory)
	r.PUT("/admin/categories/:id", middleware.AuthMiddleware(), adminCtrl.UpdateCategory)
	r.DELETE("/admin/categories/:id", middleware.AuthMiddleware(), adminCtrl.DeleteCategory)
	r.POST("/admin/locations", middleware.AuthMiddleware(), adminCtrl.AddLocation)
	r.PUT("/admin/locations/:id", middleware.AuthMiddleware(), adminCtrl.UpdateLocation)
	r.DELETE("/admin/locations/:id", middleware.AuthMiddleware(), adminCtrl.DeleteLocation)
	r.GET("/admin/jobs", middleware.AuthMiddleware(), adminCtrl.GetAllJobs)
	r.PUT("/admin/jobs/:id/status", middleware.AuthMiddleware(), adminCtrl.UpdateJobStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func runMigrations() error {
	files := []string{
		"migrations/create_users.sql",
		"migrations/create_workers.sql",
		"migrations/create_categories.sql",
		"migrations/create_locations.sql",
		"migrations/create_services.sql",
		"migrations/create_jobs.sql",
		"migrations/create_payments.sql",
		"migrations/create_reviews.sql",
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		_, err = db.Exec(string(content))
		if err != nil {
			return err
		}
	}
	return nil
}
