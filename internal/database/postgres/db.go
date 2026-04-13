package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultPort           = 5432
	defaultMaxConns       = 10
	defaultMinConns       = 2
	defaultSSLMode        = "disable"
	defaultPoolMaxIdleSec = 300
)

// Config holds the database connection settings used to build the pool.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	MaxConns int32
	MinConns int32
}

// ConfigFromEnv loads the Postgres settings from environment variables.
func ConfigFromEnv() (Config, error) {
	port, err := intFromEnv("DB_PORT", defaultPort)
	if err != nil {
		return Config{}, err
	}

	maxConns, err := intFromEnv("DB_MAX_CONNS", defaultMaxConns)
	if err != nil {
		return Config{}, err
	}

	minConns, err := intFromEnv("DB_MIN_CONNS", defaultMinConns)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  stringFromEnv("DB_SSLMODE", defaultSSLMode),
		MaxConns: int32(maxConns),
		MinConns: int32(minConns),
	}

	if cfg.Host == "" {
		return Config{}, fmt.Errorf("DB_HOST environment variable not set")
	}

	if cfg.User == "" {
		return Config{}, fmt.Errorf("DB_USER environment variable not set")
	}

	if cfg.Name == "" {
		return Config{}, fmt.Errorf("DB_NAME environment variable not set")
	}
	
	if cfg.Password == "" {
    return Config{}, fmt.Errorf("DB_PASSWORD environment variable not set")
}

	if cfg.MinConns > cfg.MaxConns {
    return Config{}, fmt.Errorf("DB_MIN_CONNS cannot be greater than DB_MAX_CONNS")
}

	return cfg, nil
}

// NewPool creates a Postgres connection pool and validates it with Ping.
func NewPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(connectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnIdleTime = time.Duration(defaultPoolMaxIdleSec) * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return pool, nil
}

func connectionString(cfg Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
}

func intFromEnv(name string, fallback int) (int, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer: %w", name, err)
	}

	return parsed, nil
}

func stringFromEnv(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return fallback
}
