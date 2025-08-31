package conversation

import (
	"time"

	"github.com/google/uuid"
)

// ================================ Create a new conversation ================================
type CreateConversationRequest struct {
	Title     string  `json:"title" binding:"required"`
	ModelUsed *string `json:"model_used,omitempty"`
}

type CreateConversationResponse struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	Title          string    `json:"title"`
	ModelUsed      *string   `json:"model_used,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	IsPinned       bool      `json:"is_pinned"`
}

// ================================ Add a message to a conversation ================================
type AddMessageRequest struct {
	ConversationID string `json:"conversation_id,omitempty"` // Optional in body, set from URL param
	Message        string `json:"message" binding:"required"`
	Sender         string `json:"sender" binding:"required,oneof=user ai system"`
}

// ================================ List of conversations (all conversations) ================================
type ConversationsListItem struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	Title          string    `json:"title"`
	IsPinned       bool      `json:"is_pinned"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetAllConversationsResponse struct {
	Conversations []ConversationsListItem `json:"conversations"`
}

// ================================ Delete a conversation ================================
type DeleteConversationResponse struct {
	Message string `json:"message"`
}

// ================================ Pin/Unpin a conversation ================================
type PinConversationRequest struct {
	IsPinned bool `json:"is_pinned"`
}

type PinConversationResponse struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	IsPinned       bool      `json:"is_pinned"`
	Message        string    `json:"message"`
}

// ================================ Conversation History ================================
type MessageHistoryItem struct {
	Message   string    `json:"message"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
}

type GetConversationResponse struct {
	Messages []MessageHistoryItem `json:"messages"`
}
