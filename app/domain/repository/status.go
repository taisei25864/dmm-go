package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindStatusByID(ctx context.Context, id int64) (*object.Status, error)
	// Fetch account which has specified username
	Create(ctx context.Context, a *object.Status) error
	// TODO: Add Other APIs

}
