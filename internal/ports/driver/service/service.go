package service

import (
	"context"
)

type ContentServicePort interface {
	Feed(ctx context.Context) any
	Top(ctx context.Context) any
	Catalog(ctx context.Context) any
	GetComments(ctx context.Context) any
	Preview(ctx context.Context) any
}
