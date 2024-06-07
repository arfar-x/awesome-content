package repository

import (
	"context"

	"awesome-content/internal/core/entity"
)

type CrudRepositoryInterface interface {
	// TODO: the second parameter must a more general struct rather than Content
	Get(ctx context.Context, entity entity.Content) (entity.Content, error)
	Create(ctx context.Context, entity entity.Content) (entity.Content, error)
	Update(ctx context.Context, entity entity.Content) (entity.Content, error)
	Delete(ctx context.Context, entity entity.Content) (bool, error)
}
