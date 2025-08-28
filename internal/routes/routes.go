package routes

import (
	userHandlers "user_service/internal/handlers/user"
	"user_service/internal/middleware"
	"user_service/internal/repository"
	userServices "user_service/internal/service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := userServices.NewUserService(userRepo)
	authService := userServices.NewAuthService(userRepo)

	// Initialize handlers
	userHandler := userHandlers.NewUserHandler(userService)
	authHandler := userHandlers.NewAuthHandler(authService)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Service is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/user_service/v1")
	{
		// Authentication routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// User routes (protected)
		users := v1.Group("/users")
		users.Use(middleware.Auth(authService)) // Apply JWT middleware
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// TODO: Add more route groups as needed
		// products := v1.Group("/products")
		// orders := v1.Group("/orders")
	}
}
