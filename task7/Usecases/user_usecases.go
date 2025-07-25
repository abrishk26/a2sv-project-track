package usecases

import (
	"github.com/abrishk26/a2sv-project-track/task7/Domain"
)

func NewUserUsecases(r domain.IUserRepository) *UserUsecases {
	return &UserUsecases{r}
}

type UserUsecases struct {
	repo domain.IUserRepository
}

func (us *UserUsecases) Add(u domain.User) error {
	return us.repo.Add(u)
}

func (us *UserUsecases) Get(id string) (*domain.User, error) {
	return us.repo.Get(id)
}

func (us *UserUsecases) Delete(id string) error {
	return us.repo.Delete(id)
}

func (us *UserUsecases) Update(id string, u domain.User) error {
	return us.repo.Update(id, u)
}

func (us *UserUsecases) GetAll() (*[]domain.User, error) {
	return us.repo.GetAll()
}
