package usecases

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"
)

type UserUsecases struct {
	repo domain.IUserRepository
}

func (us *UserUsecases) Add(ctx context.Context, u domain.User) error {
	return us.Add(ctx, u)
}

func (us *UserUsecases) Get(ctx context.Context, id string) (*domain.User, error) {
	return us.Get(ctx, id)
}

func (us *UserUsecases) Delete(ctx context.Context, id string) error {
	return us.Delete(ctx, id)
}

func (us *UserUsecases) Update(ctx context.Context, id string, u domain.User) error {
	return us.Update(ctx, id, u)
}

func (us *UserUsecases) GetAll(ctx context.Context) (*[]domain.User, error) {
	return us.GetAll(ctx)
}
