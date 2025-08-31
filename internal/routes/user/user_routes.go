package user

import (
	conversationHandlers "user_service/internal/handlers/conversation"
	userHandlers "user_service/internal/handlers/user"
	"user_service/internal/middleware"
	"user_service/internal/repository"
	conversationServices "user_service/internal/service/conversation"
	userServices "user_service/internal/service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	conversationRepo := repository.NewConversationRepository(db)

	// Initialize services
	userService := userServices.NewUserService(userRepo)
	authService := userServices.NewAuthService(userRepo)
	conversationService := conversationServices.NewConversationService(conversationRepo)

	// Initialize handlers
	userHandler := userHandlers.NewUserHandler(userService)
	authHandler := userHandlers.NewAuthHandler(authService)
	conversationHandler := conversationHandlers.NewConversationHandler(conversationService)

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
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)

			// Get all conversations for a user
			users.GET("/:id/conversations", conversationHandler.GetAllConversations)
		}

		// Conversation routes (protected)
		conversations := v1.Group("/conversations")
		conversations.Use(middleware.Auth(authService)) // Apply JWT middleware
		{
			// Create new conversation
			conversations.POST("/", conversationHandler.CreateConversation)

			// Delete conversation
			conversations.DELETE("/:conversation_id", conversationHandler.DeleteConversation)

			// Pin/unpin conversation
			conversations.PATCH("/:conversation_id/pin", conversationHandler.ToggleConversationPin)

			// Add message to conversation
			conversations.POST("/:conversation_id/messages", conversationHandler.AddMessage)

			// Get conversation history
			conversations.GET("/:conversation_id/history", conversationHandler.GetConversationHistory)
		}
	}
}
