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
	token, ok := TokenFromContext(ctx)
	if !ok {
		return domain.ErrUnauthorized
	}

	userID, err := tu.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return domain.ErrUnauthorized
		default:
			return err
		}
	}

	user, err := tu.userRepo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return domain.ErrUnauthorized
		default:
			return err
		}
	}

	t.UserID = user.ID

	return tu.taskRepo.Add(ctx, t)
}

func (tu *TaskUsecases) Get(ctx context.Context, id string) (*domain.Task, error) {
	token, ok := TokenFromContext(ctx)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	userID, err := tu.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return nil, domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return nil, domain.ErrUnauthorized
		default:
			return nil, err
		}
	}

	user, err := tu.userRepo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return nil, domain.ErrUnauthorized
		default:
			return nil, err
		}
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
	token, ok := TokenFromContext(ctx)
	if !ok {
		return domain.ErrUnauthorized
	}

	userID, err := tu.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return domain.ErrUnauthorized
		default:
			return err
		}
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

	user, err := tu.userRepo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return domain.ErrUnauthorized
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
	token, ok := TokenFromContext(ctx)
	if !ok {
		return domain.ErrUnauthorized
	}

	userID, err := tu.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return domain.ErrUnauthorized
		default:
			return err
		}
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

	user, err := tu.userRepo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return domain.ErrUnauthorized
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
	token, ok := TokenFromContext(ctx)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	userID, err := tu.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return nil, domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return nil, domain.ErrUnauthorized
		default:
			return nil, err
		}
	}

	user, err := tu.userRepo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return nil, domain.ErrUnauthorized
		default:
			return nil, err
		}
	}

	if !user.IsAdmin {
		return nil, domain.ErrUnauthorized
	}

	return tu.taskRepo.GetAll(ctx)
}
