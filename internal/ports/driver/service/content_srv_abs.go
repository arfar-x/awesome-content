package service

import (
	"context"
)

type ContentServicePort interface {
	Feed(ctx context.Context) (any, error)
	Top(ctx context.Context) (any, error)
	Catalog(ctx context.Context) (any, error)
	GetComments(ctx context.Context) (any, error)
	Preview(ctx context.Context) (any, error)
}
