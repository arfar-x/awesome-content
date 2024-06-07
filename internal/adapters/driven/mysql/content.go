package mysql

import (
	"context"

	"awesome-content/internal/core/entity"
	"awesome-content/internal/ports/driven/repository"

	"gorm.io/gorm"
)

type ContentModel struct {
	BaseModel
	Name  string
	Title string
	Rate  int
	Text  string
}

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

func (m *ContentRepository) Get(ctx context.Context, entity entity.Content) (entity.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ContentRepository) Create(ctx context.Context, entity entity.Content) (entity.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ContentRepository) Update(ctx context.Context, entity entity.Content) (entity.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ContentRepository) Delete(ctx context.Context, entity entity.Content) (bool, error) {
	//TODO implement me
	panic("implement me")
}
