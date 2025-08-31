package conversation

import (
	"net/http"
	"strconv"
	dto "user_service/internal/dto/conversation"
	conversationService "user_service/internal/service/conversation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationHandler struct {
	conversationService *conversationService.ConversationService
}

func NewConversationHandler(conversationService *conversationService.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		conversationService: conversationService,
	}
}

// CreateConversation handles creating a new conversation
// POST /conversations
func (h *ConversationHandler) CreateConversation(c *gin.Context) {
	var req dto.CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user info from JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Create conversation
	response, err := h.conversationService.CreateConversation(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// AddMessage handles adding a new message to a conversation
// POST /conversations/:conversation_id/messages
func (h *ConversationHandler) AddMessage(c *gin.Context) {
	conversationID := c.Param("conversation_id")

	var req dto.AddMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set conversation ID from URL parameter
	req.ConversationID = conversationID

	// Get user info from JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	_ = userID // We have user ID for future authorization checks

	err := h.conversationService.AddMessage(&req)
	if err != nil {
		if err.Error() == "conversation not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message added successfully"})
}

// GetConversationHistory handles retrieving conversation history
// GET /conversations/:conversation_id/history
func (h *ConversationHandler) GetConversationHistory(c *gin.Context) {
	conversationIDStr := c.Param("conversation_id")

	conversationID, err := uuid.Parse(conversationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	// Get user info from JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	_ = userID // We have user ID for future authorization checks

	response, err := h.conversationService.GetConversationHistory(conversationID)
	if err != nil {
		if err.Error() == "conversation not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllConversations handles retrieving all conversations for a user
// GET /users/:id/conversations
func (h *ConversationHandler) GetAllConversations(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get authenticated user info from JWT middleware
	authUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if user is requesting their own conversations
	if uint(userID) != authUserID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	response, err := h.conversationService.GetAllConversations(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteConversation handles deleting a conversation
// DELETE /conversations/:conversation_id
func (h *ConversationHandler) DeleteConversation(c *gin.Context) {
	conversationIDStr := c.Param("conversation_id")
	conversationID, err := uuid.Parse(conversationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	// Get authenticated user info from JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	response, err := h.conversationService.DeleteConversation(conversationID, userID.(uint))
	if err != nil {
		if err.Error() == "conversation not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		} else if err.Error() == "access denied: you can only delete your own conversations" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// ToggleConversationPin handles pinning/unpinning a conversation
// PATCH /conversations/:conversation_id/pin
func (h *ConversationHandler) ToggleConversationPin(c *gin.Context) {
	conversationIDStr := c.Param("conversation_id")
	conversationID, err := uuid.Parse(conversationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	var req dto.PinConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get authenticated user info from JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	response, err := h.conversationService.ToggleConversationPin(conversationID, userID.(uint), &req)
	if err != nil {
		if err.Error() == "conversation not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		} else if err.Error() == "access denied: you can only pin your own conversations" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
