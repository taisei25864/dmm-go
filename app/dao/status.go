package dao

import (
	"context"
	"database/sql"
	"errors"
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

// FindByUsername : ユーザ名からユーザを取得
func (r *status) FindStatusByID(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from statuses where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
