package catalog

import (
	"LearningProject/internal/catalog"
	"LearningProject/pkg/client/postgresql"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, catalog *catalog.Catalog) error {
	q := `
		  INSERT INTO public.catalog (name)
		  VALUES ($1)
		  RETURNING id
		  `
	if err := r.client.QueryRow(ctx, q, catalog.Name).Scan(&catalog.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %v, Detail: %v, Where %v, Code: %v, SQLState: %v", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			fmt.Println(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []catalog.Catalog, err error) {
	q := `SELECT to_char(id, '9'), name FROM public.catalog`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	catalogs := make([]catalog.Catalog, 0)

	for rows.Next() {
		var cat catalog.Catalog

		rows.Scan(&cat.ID, &cat.Name)

		catalogs = append(catalogs, cat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return catalogs, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (catalog.Catalog, error) {
	q := `SELECT to_char(id, '9'), name FROM public.catalog where id = $1`

	var cat catalog.Catalog
	err := r.client.QueryRow(ctx, q, id).Scan(&cat.ID, &cat.Name)
	if err != nil {
		return catalog.Catalog{}, err
	}

	return cat, nil
}

func (r *repository) Update(ctx context.Context, catalog catalog.Catalog) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := `
		  DELETE FROM public.catalog WHERE id = $1
		  `
	if err := r.client.QueryRow(ctx, q, id).Scan(); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %v, Detail: %v, Where %v, Code: %v, SQLState: %v", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			fmt.Println(newErr)
			return nil
		}
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client) catalog.Repository {
	return &repository{
		client: client,
	}
}
