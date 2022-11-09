package catalog

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, catalog *Catalog) error
	FindAll(ctx context.Context) (u []Catalog, err error)
	FindOne(ctx context.Context, id string) (Catalog, error)
	Update(ctx context.Context, catalog Catalog) error
	Delete(ctx context.Context, id string) error
}
