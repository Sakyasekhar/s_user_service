package conversation

import (
	"errors"
	"time"
	dto "user_service/internal/dto/conversation"
	"user_service/internal/models"
	"user_service/internal/repository"

	"github.com/google/uuid"
)

type ConversationService struct {
	conversationRepo *repository.ConversationRepository
}

func NewConversationService(conversationRepo *repository.ConversationRepository) *ConversationService {
	return &ConversationService{
		conversationRepo: conversationRepo,
	}
}

// CreateConversation creates a new conversation
func (s *ConversationService) CreateConversation(userID uint, req *dto.CreateConversationRequest) (*dto.CreateConversationResponse, error) {
	// Generate UUID
	conversationID := uuid.New()

	// Create conversation model
	conversation := &models.Conversation{
		ConversationID: conversationID,
		UserID:         userID,
		Title:          req.Title,
		ModelUsed:      req.ModelUsed,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsPinned:       false,
	}

	// Save conversation
	err := s.conversationRepo.CreateConversation(conversation)
	if err != nil {
		return nil, err
	}

	// Return response
	return &dto.CreateConversationResponse{
		ConversationID: conversation.ConversationID,
		Title:          conversation.Title,
		ModelUsed:      conversation.ModelUsed,
		CreatedAt:      conversation.CreatedAt,
		IsPinned:       conversation.IsPinned,
	}, nil
}

// AddMessage adds a new message to a conversation
func (s *ConversationService) AddMessage(req *dto.AddMessageRequest) error {
	// Parse conversation ID
	conversationID, err := uuid.Parse(req.ConversationID)
	if err != nil {
		return errors.New("invalid conversation ID")
	}

	// Verify conversation exists
	_, err = s.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return errors.New("conversation not found")
	}

	// Create message
	message := &models.Message{
		MessageID:      uuid.New(), // Generate UUID in Go
		ConversationID: conversationID,
		Sender:         req.Sender,
		Content:        req.Message,
		Timestamp:      time.Now(),
	}

	// Save message
	err = s.conversationRepo.CreateMessage(message)
	if err != nil {
		return err
	}

	// Update conversation timestamp
	err = s.conversationRepo.UpdateConversationTimestamp(conversationID)
	if err != nil {
		return err
	}

	return nil
}

// GetConversationHistory retrieves all messages for a conversation
func (s *ConversationService) GetConversationHistory(conversationID uuid.UUID) (*dto.GetConversationResponse, error) {
	// Verify conversation exists
	_, err := s.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, errors.New("conversation not found")
	}

	// Get messages
	messages, err := s.conversationRepo.GetConversationHistory(conversationID)
	if err != nil {
		return nil, err
	}

	// Convert to DTO
	var messageItems []dto.MessageHistoryItem
	for _, msg := range messages {
		messageItems = append(messageItems, dto.MessageHistoryItem{
			Message:   msg.Content,
			Role:      msg.Sender,
			Timestamp: msg.Timestamp,
		})
	}

	return &dto.GetConversationResponse{
		Messages: messageItems,
	}, nil
}

// GetAllConversations retrieves all conversations for a user
func (s *ConversationService) GetAllConversations(userID uint) (*dto.GetAllConversationsResponse, error) {
	// Get conversations
	conversations, err := s.conversationRepo.GetAllConversationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert to DTO
	var conversationItems []dto.ConversationsListItem
	for _, conv := range conversations {
		conversationItems = append(conversationItems, dto.ConversationsListItem{
			ConversationID: conv.ConversationID,
			Title:          conv.Title,
			IsPinned:       conv.IsPinned,
			UpdatedAt:      conv.UpdatedAt,
		})
	}

	return &dto.GetAllConversationsResponse{
		Conversations: conversationItems,
	}, nil
}

// DeleteConversation deletes a conversation and all its messages
func (s *ConversationService) DeleteConversation(conversationID uuid.UUID, userID uint) (*dto.DeleteConversationResponse, error) {
	// Verify conversation exists and belongs to user
	conversation, err := s.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, errors.New("conversation not found")
	}

	// Check if user owns the conversation
	if conversation.UserID != userID {
		return nil, errors.New("access denied: you can only delete your own conversations")
	}

	// Delete conversation (messages will be deleted due to CASCADE)
	err = s.conversationRepo.DeleteConversation(conversationID)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteConversationResponse{
		Message: "Conversation deleted successfully",
	}, nil
}

// ToggleConversationPin toggles the pin status of a conversation
func (s *ConversationService) ToggleConversationPin(conversationID uuid.UUID, userID uint, req *dto.PinConversationRequest) (*dto.PinConversationResponse, error) {
	// Verify conversation exists and belongs to user
	conversation, err := s.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, errors.New("conversation not found")
	}

	// Check if user owns the conversation
	if conversation.UserID != userID {
		return nil, errors.New("access denied: you can only pin your own conversations")
	}

	// Update pin status
	err = s.conversationRepo.UpdateConversationPin(conversationID, req.IsPinned)
	if err != nil {
		return nil, err
	}

	message := "Conversation unpinned successfully"
	if req.IsPinned {
		message = "Conversation pinned successfully"
	}

	return &dto.PinConversationResponse{
		ConversationID: conversationID,
		IsPinned:       req.IsPinned,
		Message:        message,
	}, nil
}
