package models

import (
	"time"

	"github.com/google/uuid"
)

// Conversation represents a chat conversation
type Conversation struct {
	ConversationID uuid.UUID `json:"conversation_id" gorm:"primaryKey;type:uuid;column:conversation_id"`
	UserID         uint      `json:"user_id" gorm:"not null;index;column:user_id"`
	Title          string    `json:"title" gorm:"not null;type:varchar(255);column:title"`
	ModelUsed      *string   `json:"model_used,omitempty" gorm:"type:varchar(100);column:model_used"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null;column:updated_at"`
	IsPinned       bool      `json:"is_pinned" gorm:"not null;default:false;column:is_pinned"`
}

// Message represents a single message in a conversation
type Message struct {
	MessageID       uuid.UUID  `json:"message_id" gorm:"primaryKey;type:uuid;column:message_id"`
	ConversationID  uuid.UUID  `json:"conversation_id" gorm:"not null;index;type:uuid;column:conversation_id"`
	ParentMessageID *uuid.UUID `json:"parent_message_id,omitempty" gorm:"index;type:uuid;column:parent_message_id"`
	Sender          string     `json:"sender" gorm:"not null;type:sender_role;column:sender"` // Supabase enum: 'user', 'ai', 'system'
	Content         string     `json:"content" gorm:"type:text;not null;column:content"`
	Metadata        *string    `json:"metadata,omitempty" gorm:"type:jsonb;column:metadata"`
	Timestamp       time.Time  `json:"timestamp" gorm:"not null;column:timestamp"`
}

// TableName specifies the table names
func (Conversation) TableName() string {
	return "conversations"
}

func (Message) TableName() string {
	return "messages"
}
