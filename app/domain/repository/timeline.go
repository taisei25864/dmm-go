package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetPublicTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64) (*object.Timeline, error)
}
