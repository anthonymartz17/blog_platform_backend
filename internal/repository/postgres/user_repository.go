package postgres

import (
	"context"
	"fmt"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Validates UserStore implements user.Repository.
var _ user.Repository = (*UserStore)(nil)

type UserStore struct {
	pool *pgxpool.Pool
}

func NewUserStore(pool *pgxpool.Pool) *UserStore {
	return &UserStore{pool: pool}
}

func (s *UserStore) Create(ctx context.Context, u *user.User) (*user.User, error) {
	row := s.pool.QueryRow(ctx, `
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, COALESCE(NULLIF($3, ''), 'user'))
		RETURNING id, email, password_hash, role, created_at, updated_at
	`, u.Email, u.PasswordHash, u.Role)

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("insert user: %w", err)
	}

	return u, nil
}

func (s *UserStore) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	u := &user.User{}
	row := s.pool.QueryRow(ctx, `
		SELECT id, email, password_hash, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`, email)

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("find user by email: %w", err)
	}

	return u, nil
}

func (s *UserStore) FindByID(ctx context.Context, id string) (*user.User, error) {
	u := &user.User{}
	row := s.pool.QueryRow(ctx, `
		SELECT id, email, password_hash, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`, id)

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}

	return u, nil
}


