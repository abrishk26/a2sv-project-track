package usecases

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"
)

type TaskUsecases struct {
	repo domain.ITaskRepository
}

func (ta *TaskUsecases) Add(ctx context.Context, u domain.User) error {
	return ta.Add(ctx, u)
}

func (ta *TaskUsecases) Get(ctx context.Context, id string) (*domain.User, error) {
	return ta.Get(ctx, id)
}

func (ta *TaskUsecases) Delete(ctx context.Context, id string) error {
	return ta.Delete(ctx, id)
}

func (ta *TaskUsecases) Update(ctx context.Context, id string, u domain.User) error {
	return ta.Update(ctx, id, u)
}

func (ta *TaskUsecases) GetAll(ctx context.Context) (*[]domain.User, error) {
	return ta.GetAll(ctx)
}
