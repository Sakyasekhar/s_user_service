package repository

import (
	"user_service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

// CreateConversation creates a new conversation
func (r *ConversationRepository) CreateConversation(conversation *models.Conversation) error {
	return r.db.Create(conversation).Error
}

// CreateMessage creates a new message in a conversation
func (r *ConversationRepository) CreateMessage(message *models.Message) error {
	return r.db.Create(message).Error
}

// GetConversationHistory retrieves all messages for a conversation
func (r *ConversationRepository) GetConversationHistory(conversationID uuid.UUID) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("conversation_id = ?", conversationID).Order("timestamp ASC").Find(&messages).Error
	return messages, err
}

// GetAllConversationsByUserID retrieves all conversations for a user
func (r *ConversationRepository) GetAllConversationsByUserID(userID uint) ([]models.Conversation, error) {
	var conversations []models.Conversation
	err := r.db.Where("user_id = ?", userID).Order("updated_at DESC").Find(&conversations).Error
	return conversations, err
}

// GetConversationByID retrieves a conversation by ID
func (r *ConversationRepository) GetConversationByID(conversationID uuid.UUID) (*models.Conversation, error) {
	var conversation models.Conversation
	err := r.db.Where("conversation_id = ?", conversationID).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// UpdateConversationTimestamp updates the updated_at field of a conversation
func (r *ConversationRepository) UpdateConversationTimestamp(conversationID uuid.UUID) error {
	return r.db.Model(&models.Conversation{}).Where("conversation_id = ?", conversationID).Update("updated_at", "NOW()").Error
}

// DeleteConversation deletes a conversation and all its messages
func (r *ConversationRepository) DeleteConversation(conversationID uuid.UUID) error {
	return r.db.Where("conversation_id = ?", conversationID).Delete(&models.Conversation{}).Error
}

// UpdateConversationPin updates the is_pinned status of a conversation
func (r *ConversationRepository) UpdateConversationPin(conversationID uuid.UUID, isPinned bool) error {
	return r.db.Model(&models.Conversation{}).Where("conversation_id = ?", conversationID).Update("is_pinned", isPinned).Error
}
