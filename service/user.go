package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hipzz/orm-practice/models"
	"github.com/hipzz/orm-practice/repository"
)

type User interface {
	Save(ctx context.Context, users ...models.User) error
	Get(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, id string) (models.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepository repository.User
}

// Delete implements User
func (us *userService) Delete(ctx context.Context, id string) error {
	return us.userRepository.Delete(ctx, id)
}

// Get implements User
func (us *userService) Get(ctx context.Context) ([]models.User, error) {
	return us.userRepository.Get(ctx)
}

// GetById implements User
func (us *userService) GetById(ctx context.Context, id string) (models.User, error) {
	return us.userRepository.GetById(ctx, id)
}

// Save implements User
func (us *userService) Save(ctx context.Context, users ...models.User) error {
	for i := 0; i < len(users); i++ {
		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		users[i].ID = id.String()
		users[i].LastUpdate = time.Now().Unix()
	}
	return us.userRepository.Save(ctx, users...)
}

var _ User = (*userService)(nil)

func NewUserService(repo repository.User) User {
	return &userService{userRepository: repo}
}
