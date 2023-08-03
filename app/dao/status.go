package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) Create(ctx context.Context, s *object.Status) error {
	query := "INSERT INTO statuses (account_id, content) VALUES (?, ?)"
	_, err := r.db.Exec(query, s.AccountID, s.Content)
	if err != nil {
		return fmt.Errorf("failed to register new content to db: %w", err)
	}
	return nil
}
