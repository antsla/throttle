package storage

import (
	"context"
)

type Storage interface {
	Increment(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (int, error)
}
