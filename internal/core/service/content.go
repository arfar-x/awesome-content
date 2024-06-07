package service

import (
	"context"

	"awesome-content/internal/ports/driven/repository"
	"awesome-content/internal/ports/driver/service"
)

type ContentService struct {
	Repo repository.ContentRepositoryInterface
}

func NewContentService(repo repository.ContentRepositoryInterface) service.ContentServicePort {
	return &ContentService{Repo: repo}
}

func (c *ContentService) Feed(ctx context.Context) any {
	//TODO implement me
	panic("implement me")
}

func (c *ContentService) Top(ctx context.Context) any {
	//TODO implement me
	panic("implement me")
}

func (c *ContentService) Catalog(ctx context.Context) any {
	//TODO implement me
	panic("implement me")
}

func (c *ContentService) GetComments(ctx context.Context) any {
	//TODO implement me
	panic("implement me")
}

func (c *ContentService) Preview(ctx context.Context) any {
	//TODO implement me
	panic("implement me")
}
