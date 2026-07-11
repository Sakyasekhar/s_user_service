package main

import (
	"context"
	"log"
	"os"
	"user_service/config"
	"user_service/internal/database"
	"user_service/internal/middleware"
	userRoutes "user_service/internal/routes/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Setup routes
	userRoutes.SetupRoutes(router, db)

	// Start server
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		ginLambda = ginadapter.NewV2(router)
		lambda.Start(Handler)
	} else {
		log.Printf("Server starting on port %s", cfg.Port)
		if err := router.Run(":" + cfg.Port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}
}
