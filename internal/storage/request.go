package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s *Storage) AddAsset(asset string) error {
	var err error
	var result pgconn.CommandTag
	add := `INSERT INTO watch(name) VALUES ($1)`
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		log.Println("Transaction begin error")
		return err
	}

	result, err = tx.Exec(context.Background(), add, asset)
	affectedRows := result.RowsAffected()
	if err != nil || affectedRows <= 0 {
		if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
			return fmt.Errorf("Transaction rollback error: %w\nFirst error: %w", rollbackErr, err)
		}
		return fmt.Errorf("Transaction rollbacked with error: %w", err)
	}

	if commitErr := tx.Commit(context.Background()); commitErr != nil {
		return fmt.Errorf("Transaction commit error: %w", commitErr)
	}
	log.Println("Transaction commited.")

	return nil
}

func (s *Storage) RemoveAsset(asset string) error {
	var err error
	var result pgconn.CommandTag
	add := `DELETE FROM watch WHERE name = $1`
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		log.Println("Transaction begin error")
		return err
	}

	result, err = tx.Exec(context.Background(), add, asset)
	affectedRows := result.RowsAffected()
	if err != nil || affectedRows <= 0 {
		if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
			return fmt.Errorf("Transaction rollback error: %w\nFirst error: %w", rollbackErr, err)
		}
		return fmt.Errorf("Transaction rollbacked with error: %w", err)
	}

	if commitErr := tx.Commit(context.Background()); commitErr != nil {
		return fmt.Errorf("Transaction commit error: %w", commitErr)
	}
	log.Println("Transaction commited.")

	return nil
}
