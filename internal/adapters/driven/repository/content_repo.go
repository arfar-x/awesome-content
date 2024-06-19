package repository

import (
	"context"

	"awesome-content/internal/core/entity"
	"awesome-content/internal/ports/driven/repository"

	"gorm.io/gorm"
)

type ContentRepository struct {
	BaseRepositoryAdapter
}

func NewContentRepository(db *gorm.DB) repository.ContentRepositoryInterface {
	return &ContentRepository{
		BaseRepositoryAdapter: BaseRepositoryAdapter{
			DB: db,
		},
	}
}

func (r *ContentRepository) Get(ctx context.Context, entity entity.Content) (entity.Content, error) {
	panic("implement me")
}

func (r *ContentRepository) Create(ctx context.Context, entity entity.Content) (entity.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContentRepository) Update(ctx context.Context, entity entity.Content) (entity.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContentRepository) Delete(ctx context.Context, entity entity.Content) (bool, error) {
	//TODO implement me
	panic("implement me")
}
