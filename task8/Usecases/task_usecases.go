package usecases

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

func NewTaskUsecases(tr domain.ITaskRepository, ur domain.IUserRepository, ts domain.ITokenService) *TaskUsecases {
	return &TaskUsecases{tr, ur, ts}
}

type TaskUsecases struct {
	taskRepo     domain.ITaskRepository
	userRepo     domain.IUserRepository
	tokenService domain.ITokenService
}

func (tu *TaskUsecases) Add(ctx context.Context, t domain.Task) error {
	userID, err := getUserID(ctx, tu.tokenService)
	if err != nil {
		return err
	}

	user, err := getUser(ctx, userID, tu.userRepo)
	if err != nil {
		return err
	}

	t.UserID = user.ID

	return tu.taskRepo.Add(ctx, t)
}

func (tu *TaskUsecases) Get(ctx context.Context, id string) (*domain.Task, error) {
	userID, err := getUserID(ctx, tu.tokenService)
	if err != nil {
		return nil, err
	}

	user, err := getUser(ctx, userID, tu.userRepo)
	if err != nil {
		return nil, err
	}

	task, err := tu.taskRepo.Get(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			return nil, domain.ErrTaskNotFound
		default:
			return nil, err
		}
	}

	if user.IsAdmin {
		return task, nil
	}

	if task.UserID != user.ID {
		return nil, domain.ErrUnauthorized
	}

	return task, nil
}

func (tu *TaskUsecases) Delete(ctx context.Context, id string) error {
	userID, err := getUserID(ctx, tu.tokenService)
	if err != nil {
		return err
	}

	user, err := getUser(ctx, userID, tu.userRepo)
	if err != nil {
		return err
	}

	task, err := tu.taskRepo.Get(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			return domain.ErrTaskNotFound
		default:
			return err
		}
	}

	if user.IsAdmin {
		return tu.taskRepo.Delete(ctx, id)
	}

	if task.UserID != userID {
		return domain.ErrUnauthorized
	}

	return tu.taskRepo.Delete(ctx, id)
}

func (tu *TaskUsecases) Update(ctx context.Context, id string, t domain.Task) error {
	userID, err := getUserID(ctx, tu.tokenService)
	if err != nil {
		return err
	}

	user, err := getUser(ctx, userID, tu.userRepo)
	if err != nil {
		return err
	}

	task, err := tu.taskRepo.Get(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			return domain.ErrTaskNotFound
		default:
			return err
		}
	}

	if user.IsAdmin {
		return tu.taskRepo.Update(ctx, id, t)
	}
	if task.UserID != userID {
		return domain.ErrUnauthorized
	}

	return tu.taskRepo.Update(ctx, id, t)
}

func (tu *TaskUsecases) GetAll(ctx context.Context) ([]domain.Task, error) {
	userID, err := getUserID(ctx, tu.tokenService)
	if err != nil {
		return nil, err
	}

	user, err := getUser(ctx, userID, tu.userRepo)
	if err != nil {
		return nil, err
	}

	if !user.IsAdmin {
		return nil, domain.ErrUnauthorized
	}

	return tu.taskRepo.GetAll(ctx)
}
