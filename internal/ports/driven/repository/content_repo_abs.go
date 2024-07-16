package repository

import (
	"awesome-content/internal/core/entity"
	"context"
)

type ContentRepositoryInterface interface {
	Get(ctx context.Context, entity entity.Content) (entity.Content, error)
	Create(ctx context.Context, entity entity.Content) (entity.Content, error)
	Update(ctx context.Context, entity entity.Content) (entity.Content, error)
	Delete(ctx context.Context, entity entity.Content) (bool, error)
	Query(ctx context.Context, collection entity.Content) (entity.Content, error)
}
