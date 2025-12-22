package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dyleme/timecache"
	"github.com/dyleme/notifier/internal/domain/apperr"
)

type Generic[T any] struct {
	cache *timecache.Cache[string, T]
}

func NewGeneric[T any]() *Generic[T] {
	return &Generic[T]{
		cache: timecache.New[string, T](),
	}
}

func (gc *Generic[T]) Get(_ context.Context, key string) (T, error) {
	obj, err := gc.cache.Get(key)
	if err != nil {
		var zero T
		if errors.Is(err, timecache.ErrNotExists) {
			return zero, apperr.ErrNotFound
		}

		return zero, err
	}

	return obj, nil
}

func (gc *Generic[I]) Put(ctx context.Context, key string, obj I) error {
	gc.cache.StoreDefDur(key, obj)

	return nil
}

func (gc *Generic[I]) Wrap(ctx context.Context, key string, f func() (I, error)) (I, error) {
	obj, err := gc.Get(ctx, key)
	if err == nil {
		return obj, nil
	}
	var zero I
	if !errors.Is(err, apperr.ErrNotFound) {
		return zero, fmt.Errorf("get: %w", err)
	}

	obj, err = f()
	if err != nil {
		return zero, fmt.Errorf("get func: %w", err)
	}

	err = gc.Put(ctx, key, obj)
	if err != nil {
		return zero, fmt.Errorf("put by key: %w", err)
	}

	return obj, nil
}
