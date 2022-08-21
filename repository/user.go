package repository

import (
	"context"

	"github.com/hipzz/orm-practice/models"
)

type User interface {
	Save(ctx context.Context, users ...models.User) error
	Get(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, id string) (models.User, error)
	Delete(ctx context.Context, id string) error
}
