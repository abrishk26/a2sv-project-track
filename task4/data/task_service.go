package data

import (
	"sync"

	"github.com/abrishk26/a2sv-project-track/task4/models"
)

type TaskRepository interface {
	Add(t models.Task) Task
	Get(id int) (Task, error)
	Delete(id int) (Task, error)
	Update(id int, t models.Task) (Task, error)
	GetAll() []Task
}

type TaskManager struct {
	tasksMu sync.Mutex
	tasks map[int]models[Task]
}

func (tm *TaskManager) Add(t models.Task) models.Task {
	tm.tasksMu.Lock()

	t.ID = len(tm.tasks) + 1
	tm.tasks[t.ID] = t

	tm.tasks.Mu.Unlock()

	return t
}

func (tm *TaskManager) Get(id int) (models.Task, error) {
	tm.tasksMu.Lock()

	if task, ok := tm.tasks[id]; ok {
		return task, nil
	}

	tm.tasksMu.Unlock()
	return models.Task{}, errors.New("Task with the given id does not exist")
}

func (tm *TaskManager) Delete(id int) (models.Task, error) {
	tm.tasksMu.Lock()

	if task, ok := tm.tasks[id]; ok {
		delete(tm.tasks, id)
		return task, nil
	}

	tm.tasksMu.Unlock()
	return models.Task{}, errors.New("Task with the given id does not exist")
}

func (tm *TaskManger) Update(id int, t models.Task) (models.Task, error) {
	tm.tasksMu.Lock()

	if task, ok := tm.tasks[id]; ok {
		if t.Title != "" {
			task.Title = t.Title
		}

		if t.Description != "" {
			task.Description = t.Description
		}

		if t.DueDate != "" {
			task.DueDate = t.DueDate
		}

		if t.Done {
			task.Done = t.Done
		}

		tm.tasks[id] = task
		return task, nil
	}

	tm.tasksMu.Unlock()
	return models.Task{}, errors.New("Task with the given id does not exist")
}

func (tm *TaskManger) GetAll() []models.Task {
	tasks := []models.Task{}

	tm.tasksMu.Lock()

	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}

	tm.tasksMu.Unlock()
	return tasks
}
