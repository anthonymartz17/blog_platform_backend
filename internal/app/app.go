package app

import (
	"context"
	"errors"
	"os"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/auth"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/database/postgres"
	httpServer "github.com/anthonymartz17/blog_platform_backend.git/internal/http"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/jackc/pgx/v5/pgxpool"

	postRepository "github.com/anthonymartz17/blog_platform_backend.git/internal/store/postgres"
)

type App struct {
	server   *httpServer.Server
	connPool *pgxpool.Pool
}

// New initializes the application dependencies and returns an App.
func New() (*App, error) {

	ctx := context.Background()

	authService := auth.New()

	cfg, err := postgres.ConfigFromEnv()
	if err != nil {
		return nil, err
	}

	pool, err := postgres.NewPool(ctx, cfg)
	if err != nil {
		return nil, err
	}

	postRepo := postRepository.NewPostStore(pool)
	postSvc := post.NewPostService(postRepo)
	postHandler := post.NewHandler(postSvc)

	httpRouter := httpServer.NewRouter()
	postHandler.RegisterRoutes(httpRouter, authService)

	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT environment variable not set")
	}

	svr := httpServer.NewServer(port, httpRouter, func(context.Context) error {
		pool.Close()
		return nil
	})

	newApp := &App{
		server:   svr,
		connPool: pool,
	}

	return newApp, nil

}

// Start starts the application HTTP server.
func (a *App) Start() error {
	return a.server.Start()
}

// Shutdown gracefully stops the application.
func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

// Address returns the HTTP server bind address.
func (a *App) Address() string {
	return a.server.Address()
}
