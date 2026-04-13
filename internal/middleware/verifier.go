package middleware

import (
	"context"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/auth"
)

// AuthVerifier defines token validation method
type AuthVerifier interface {
	VerifyToken(ctx context.Context, token string) (*auth.Claims, error)
}
