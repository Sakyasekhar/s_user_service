package models

import "time"

// User represents a user in the system
type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey;column:user_id"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}
