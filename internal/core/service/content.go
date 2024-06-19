package service

import (
	"context"
	"errors"

	"awesome-content/internal/core/entity"
	"awesome-content/internal/ports/driven/repository"
	"awesome-content/internal/ports/driver/service"
	"awesome-content/pkg/errconv"
)

type ContentService struct {
	Repo repository.ContentRepositoryInterface
}

func NewContentService(repo repository.ContentRepositoryInterface) service.ContentServicePort {
	return &ContentService{Repo: repo}
}

func (srv *ContentService) Feed(ctx context.Context) (any, error) {
	result, err := srv.Repo.Get(ctx, entity.Content{})

	if err != nil {
		if errors.Is(err, errconv.ErrNotFound) {
			return nil, errconv.ErrNotFound
		}
		return nil, err
	}

	return result, nil
}

func (srv *ContentService) Top(ctx context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (srv *ContentService) Catalog(ctx context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (srv *ContentService) GetComments(ctx context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (srv *ContentService) Preview(ctx context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}
