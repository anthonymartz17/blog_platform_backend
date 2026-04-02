package auth

import (
	"context"
	"errors"
)

var ErrInvalidToken = errors.New("invalid token")

// Claims represents the authenticated user information extracted from a token.
type Claims struct {
	UID string
}

type Service struct{}

// New creates an auth service.
func New() *Service {
	return &Service{}
}

// VerifyToken validates an access token and returns the authenticated claims.
func (a *Service) VerifyToken(ctx context.Context, token string) (*Claims, error) {
	return nil, ErrInvalidToken
}
