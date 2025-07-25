package usecases

import (
	"github.com/abrishk26/a2sv-project-track/task7/Domain"
)

func NewTaskUsecases(r domain.ITaskRepository) *TaskUsecases {
	return &TaskUsecases{r}
}

type TaskUsecases struct {
	repo domain.ITaskRepository
}

func (ta *TaskUsecases) Add(u domain.Task) error {
	return ta.repo.Add(u)
}

func (ta *TaskUsecases) Get(id string) (*domain.Task, error) {
	return ta.repo.Get(id)
}

func (ta *TaskUsecases) Delete(id string) error {
	return ta.repo.Delete(id)
}

func (ta *TaskUsecases) Update(id string, u domain.Task) error {
	return ta.repo.Update(id, u)
}

func (ta *TaskUsecases) GetAll() (*[]domain.Task, error) {
	return ta.repo.GetAll()
}
