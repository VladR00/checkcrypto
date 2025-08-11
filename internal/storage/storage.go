package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Db *pgxpool.Pool
}

type Config struct {
	User     string
	Password string
	Host     string
	DbName   string
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{Db: db}
}

// migrate -path ./internal/domain/migrations/ -database "postgres://user:pass@localhost:5042/database?sslmode=disable" up
func migrations(url string) error {
	m, err := migrate.New("file://internal/migrations/postgresql/migrations", url)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Println("PostgreSql: Migrations applied.")
	return nil
}

func ConnectPostgreSQL() (*pgxpool.Pool, error) {
	cfg := Config{User: "user", Password: "asd", Host: "postgres", DbName: "test"}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.DbName)

	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("Error parseConfig: %w", err)
	}

	poolConfig.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Error apply poolConfig: %w", err)
	}

	if err := pingDB(pool); err != nil {
		return nil, fmt.Errorf("Error ping DB: %w", err)
	}
	if err := migrations(dbURL); err != nil {
		return nil, fmt.Errorf("Error apply migrations: %w", err)
	}
	return pool, nil
}

func pingDB(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return err
	}
	log.Println("Posgtresql connected")
	return nil
}
