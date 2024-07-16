package repository

import (
	"context"
	"errors"

	"awesome-content/internal/adapters/driven/repository/model/mysql"
	"awesome-content/internal/core/entity"
	"awesome-content/internal/ports/driven/repository"

	"gorm.io/gorm"
)

type ContentRepository struct {
	BaseRepositoryAdapter
	Model        mysql.ContentModel
	DomainEntity entity.Content
}

func NewContentRepository(db *gorm.DB) repository.ContentRepositoryInterface {
	return &ContentRepository{
		BaseRepositoryAdapter: BaseRepositoryAdapter{
			DB: db,
		},
		Model:        mysql.ContentModel{},
		DomainEntity: entity.Content{},
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

func (r *ContentRepository) Query(ctx context.Context, collection entity.Content) (entity.Content, error) {
	r.Model = mysql.ContentModel{
		Title: collection.Title,
		Rate:  collection.Rate,
		Text:  collection.Text,
	}

	result := r.DB.WithContext(ctx).
		Where(collection).
		Find(&r.Model)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Content{}, err
		} else {
			panic(err)
		}
	}

	return r.toDomainEntity(), nil

	// return toCollection(content), nil
}

func (r *ContentRepository) toDomainEntity() entity.Content {
	r.DomainEntity.Title = r.Model.Title
	r.DomainEntity.Rate = r.Model.Rate
	r.DomainEntity.Text = r.Model.Text
	return r.DomainEntity
}

// func toCollection(models []mysql.ContentModel) entity.Collection {
// 	collection := entity.Collection{}

// 	for _, model := range models {
// 		entity := entity.Entity{
// 			ID:        model.ID,
// 			CreatedAt: model.CreatedAt,
// 			UpdatedAt: model.UpdatedAt,
// 		}
// 		collection.Entities = append(collection.Entities, entity)
// 	}

// 	return collection
// }
