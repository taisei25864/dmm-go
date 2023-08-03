package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch account which has specified username
	Create(ctx context.Context, a *object.Status) error
	// TODO: Add Other APIs

}
