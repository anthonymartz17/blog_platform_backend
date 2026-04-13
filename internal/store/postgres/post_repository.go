package postgres

import (
	"context"
	"fmt"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Validates PostStore implements PostRepository interface.
var _ post.PostRepository = (*PostStore)(nil)

type PostStore struct {
	pool *pgxpool.Pool
}

func NewPostStore(p *pgxpool.Pool) *PostStore {
	return &PostStore{
		pool: p,
	}
}

func (r *PostStore) Save(ctx context.Context, p *post.Post) (*post.Post, error) {
	row := r.pool.QueryRow(ctx, `
		INSERT INTO posts(user_id,content)
		VALUES($1,$2)
		RETURNING id, created_at, updated_at
	`, p.UserID, p.Content)

	if err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, fmt.Errorf("insert post: %w", err)
	}

	return p, nil
}

func (r *PostStore) GetPosts(ctx context.Context) ([]post.Post, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, content, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC, id DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("select posts: %w", err)
	}
	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan post: %w", err)
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate posts: %w", err)
	}

	return posts, nil
}
