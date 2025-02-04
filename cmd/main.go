package main

import (
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
	dsn := "host=localhost user=postgres password=postgres dbname=job_platform sslmode=disable"
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Run migrations
	if err := runMigrations(); err != nil {
		logger.Fatal("Failed to run migrations", zap.Error(err))
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

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
