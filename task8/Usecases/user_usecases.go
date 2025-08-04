package usecases

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

type ctxKey string

const tokenKey = "token"

func ContextWithToken(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, tokenKey, userID)
}

func TokenFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(tokenKey)
	userID, ok := v.(string)
	return userID, ok
}

func NewUserUsecases(r domain.IUserRepository, ps domain.IPasswordService, ts domain.ITokenService) *UserUsecases {
	return &UserUsecases{r, ps, ts}
}

type UserUsecases struct {
	repo            domain.IUserRepository
	passwordService domain.IPasswordService
	tokenService    domain.ITokenService
}

func (us *UserUsecases) Login(ctx context.Context, email, password string) (string, error) {
	user, err := us.repo.GetByEmail(ctx, email)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return "", domain.ErrInvalidCredential
		default:
			return "", err
		}
	}

	err = us.passwordService.Verify(password, user.PasswordHash)
	if err != nil {
		switch err {
		case domain.ErrPasswordVerificationFailed:
			return "", domain.ErrInvalidCredential
		default:
			return "", err
		}
	}

	token, err := us.tokenService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserUsecases) Register(ctx context.Context, u domain.User) error {
	return us.repo.Add(ctx, u)
}

func (us *UserUsecases) Get(ctx context.Context, id string) (*domain.User, error) {
	token, ok := TokenFromContext(ctx)
	if !ok {
		return nil, domain.ErrUnautorized
	}

	userID, err := us.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return nil, domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return nil, domain.ErrUnautorized
		default:
			return nil, err
		}
	}

	user, err := us.repo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return nil, domain.ErrUnautorized
		default:
			return nil, err
		}
	}

	if user.IsAdmin {
		return us.repo.GetByID(ctx, id)
	}

	if userID != id {
		return nil, domain.ErrUnautorized
	}

	return us.repo.GetByID(ctx, id)
}

func (us *UserUsecases) Delete(ctx context.Context, id string) error {
	token, ok := TokenFromContext(ctx)
	if !ok {
		return domain.ErrUnautorized
	}

	userID, err := us.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return domain.ErrUnautorized
		default:
			return err
		}
	}

	user, err := us.repo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return domain.ErrUnautorized
		default:
			return err
		}
	}

	if user.IsAdmin {
		return us.repo.Delete(ctx, id)
	}

	if userID != id {
		return domain.ErrUnautorized
	}

	return us.repo.Delete(ctx, id)
}

func (us *UserUsecases) Update(ctx context.Context, id string, u domain.User) error {
	token, ok := TokenFromContext(ctx)
	if !ok {
		return domain.ErrUnautorized
	}

	userID, err := us.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return domain.ErrUnautorized
		default:
			return err
		}
	}

	user, err := us.repo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return domain.ErrUnautorized
		default:
			return err
		}
	}

	if user.IsAdmin {
		return us.repo.Update(ctx, id, u)
	}

	if userID != id {
		return domain.ErrUnautorized
	}

	return us.repo.Update(ctx, id, u)
}

func (us *UserUsecases) GetAll(ctx context.Context) ([]domain.User, error) {
	token, ok := TokenFromContext(ctx)
	if !ok {
		return nil, domain.ErrUnautorized
	}

	userID, err := us.tokenService.VerifyToken(token)
	if err != nil {
		switch err {
		case domain.ErrExpiredToken:
			return nil, domain.ErrExpiredToken
		case domain.ErrInvalidToken:
			return nil, domain.ErrUnautorized
		default:
			return nil, err
		}
	}

	user, err := us.repo.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			return nil, domain.ErrUnautorized
		default:
			return nil, err
		}
	}

	if !user.IsAdmin {
		return nil, domain.ErrUnautorized
	}

	return us.repo.GetAll(ctx)
}
