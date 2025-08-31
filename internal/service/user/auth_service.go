package service

import (
	"errors"
	"os"
	"time"
	dto "user_service/internal/dto/user"
	"user_service/internal/models"
	"user_service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication logic
type AuthService struct {
	userRepo *repository.UserRepository
}

// NewAuthService creates a new authentication service
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register creates a new user account
func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
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

	// Generate JWT token
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, err
	}

	// Convert to response
	userResponse := &dto.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: int64(user.CreatedAt.Unix()),
		UpdatedAt: int64(user.UpdatedAt.Unix()),
	}

	return &dto.AuthResponse{
		Token: token,
		User:  *userResponse,
	}, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, err
	}

	// Convert to response
	userResponse := &dto.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: int64(user.CreatedAt.Unix()),
		UpdatedAt: int64(user.UpdatedAt.Unix()),
	}

	return &dto.AuthResponse{
		Token: token,
		User:  *userResponse,
	}, nil
}

// generateJWT creates a JWT token for a user
func (s *AuthService) generateJWT(user *models.User) (string, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET not configured")
	}

	// Get JWT expiry from environment (default: 24h)
	jwtExpiry := os.Getenv("JWT_EXPIRY")
	if jwtExpiry == "" {
		jwtExpiry = "24h"
	}

	// Parse expiry duration
	expiryDuration, err := time.ParseDuration(jwtExpiry)
	if err != nil {
		expiryDuration = 24 * time.Hour
	}

	// Create claims
	claims := &dto.Claims{
		UserID:   user.UserID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * expiryDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "user_service",
			Subject:   user.Email,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT token and returns the claims
func (s *AuthService) ValidateJWT(tokenString string) (*dto.Claims, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET not configured")
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
