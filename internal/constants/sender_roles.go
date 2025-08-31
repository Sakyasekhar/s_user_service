package constants

// SenderRole constants for message senders
const (
	SenderRoleUser   = "user"
	SenderRoleAI     = "ai"
	SenderRoleSystem = "system"
)

// ValidSenderRoles returns a slice of all valid sender roles
func ValidSenderRoles() []string {
	return []string{SenderRoleUser, SenderRoleAI, SenderRoleSystem}
}

// IsValidSenderRole checks if a sender role is valid
func IsValidSenderRole(role string) bool {
	for _, validRole := range ValidSenderRoles() {
		if role == validRole {
			return true
		}
	}
	return false
}
