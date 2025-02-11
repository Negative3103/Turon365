package main

import (
    "Turon365/internal/controller"
    "Turon365/internal/middleware"
    "Turon365/internal/repository"
    "Turon365/internal/storage"
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
    storage.InitMinio()
}

func main() {
    r := gin.Default()

    // Middleware
    r.Use(middleware.LoggingMiddleware(logger))
    r.Use(middleware.AuthMiddleware())

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

    // Routes
    r.POST("/users", userCtrl.RegisterUser)
    r.GET("/users/:id", userCtrl.GetUser)
    r.PUT("/users/:id", userCtrl.UpdateUser)
    r.DELETE("/users/:id", userCtrl.DeleteUser)

    r.POST("/workers", workerCtrl.RegisterWorker)
    r.GET("/workers/:id", workerCtrl.GetWorker)
    r.PUT("/workers/:id", workerCtrl.UpdateWorker)
    r.DELETE("/workers/:id", workerCtrl.DeleteWorker)

    r.POST("/categories", categoryCtrl.CreateCategory)
    r.GET("/categories/:id", categoryCtrl.GetCategory)
    r.PUT("/categories/:id", categoryCtrl.UpdateCategory)
    r.DELETE("/categories/:id", categoryCtrl.DeleteCategory)

    r.POST("/locations", locationCtrl.CreateLocation)
    r.GET("/locations/:id", locationCtrl.GetLocation)
    r.PUT("/locations/:id", locationCtrl.UpdateLocation)
    r.DELETE("/locations/:id", locationCtrl.DeleteLocation)

    r.POST("/services", serviceCtrl.CreateService)
    r.GET("/services/:id", serviceCtrl.GetService)
    r.PUT("/services/:id", serviceCtrl.UpdateService)
    r.DELETE("/services/:id", serviceCtrl.DeleteService)

    r.POST("/jobs", jobCtrl.CreateJob)
    r.GET("/jobs/:id", jobCtrl.GetJob)
    r.PUT("/jobs/:id", jobCtrl.UpdateJob)
    r.DELETE("/jobs/:id", jobCtrl.DeleteJob)

    r.POST("/payments", paymentCtrl.CreatePayment)
    r.GET("/payments/:id", paymentCtrl.GetPayment)
    r.PUT("/payments/:id", paymentCtrl.UpdatePayment)
    r.DELETE("/payments/:id", paymentCtrl.DeletePayment)

    r.POST("/reviews", reviewCtrl.CreateReview)
    r.GET("/reviews/:id", reviewCtrl.GetReview)
    r.PUT("/reviews/:id", reviewCtrl.UpdateReview)
    r.DELETE("/reviews/:id", reviewCtrl.DeleteReview)

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