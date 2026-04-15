package user

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("user service not implemented")

// validates Service implements AuthService interface
var _ AuthService = (*Service)(nil)

type TokenIssuer interface {
	IssueAccessToken(ctx context.Context, userID string) (string, error)
}

// Service implements user authentication business logic.
type Service struct {
	repo   Repository
	tokens TokenIssuer
}

// NewService creates a new user service.
func NewService(repo Repository, tokens TokenIssuer) *Service {
	return &Service{
		repo:   repo,
		tokens: tokens,
	}
}

// Signup creates a user account and returns authentication data.
func (s *Service) Signup(ctx context.Context, email, password string) (*AuthResponse, error) {
	return nil, ErrNotImplemented
}

// Login authenticates a user and returns authentication data.
func (s *Service) Login(ctx context.Context, email, password string) (*AuthResponse, error) {
	return nil, ErrNotImplemented
}
