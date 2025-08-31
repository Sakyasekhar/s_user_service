package service

import (
	"errors"
	dto "user_service/internal/dto/user"
	"user_service/internal/models"
	"user_service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser creates a new user with business logic validation
func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if email already exists
	if s.userRepo.EmailExists(req.Email) {
		return nil, errors.New("email already exists")
	}

	// Check if username already exists
	if s.userRepo.UsernameExists(req.Username) {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.toUserResponse(user), nil
}

// UpdateUser updates a user with business logic validation
func (s *UserService) UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check email uniqueness if updating
	if req.Email != nil && *req.Email != user.Email {
		if s.userRepo.EmailExists(*req.Email) {
			return nil, errors.New("email already exists")
		}
		user.Email = *req.Email
	}

	// Check username uniqueness if updating
	if req.Username != nil && *req.Username != user.Username {
		if s.userRepo.UsernameExists(*req.Username) {
			return nil, errors.New("username already exists")
		}
		user.Username = *req.Username
	}

	// Update other fields
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uint) error {
	// Check if user exists
	if _, err := s.userRepo.GetByID(id); err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}

// toUserResponse converts a User model to UserResponse
func (s *UserService) toUserResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: int64(user.CreatedAt.Unix()),
		UpdatedAt: int64(user.UpdatedAt.Unix()),
	}
}
