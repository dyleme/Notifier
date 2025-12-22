package repository

import "context"

type GenericCache[A any] interface {
	Wrap(ctx context.Context, key string, f func() (A, error)) (A, error)
	Put(ctx context.Context, key string, val A) error
	Get(ctx context.Context, id string) (A, error)
}
