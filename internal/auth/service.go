package auth

import (
	"context"
	"errors"
)

var ErrInvalidToken = errors.New("invalid token")

var ErrTokenNotImplemented = errors.New("token service not implemented")

// Claims represents the authenticated user information extracted from a token.
type Claims struct {
	UID string
}

// TokenService owns token issuing and verification.
type TokenService struct{}

// NewTokenService creates the shared token service.
func NewTokenService() *TokenService {
	return &TokenService{}
}

// IssueAccessToken creates an access token for a user.
func (t *TokenService) IssueAccessToken(ctx context.Context, userID string) (string, error) {
	return "", ErrTokenNotImplemented
}

// VerifyToken validates an access token and returns the authenticated claims.
func (t *TokenService) VerifyToken(ctx context.Context, token string) (*Claims, error) {
	return nil, ErrInvalidToken
}
