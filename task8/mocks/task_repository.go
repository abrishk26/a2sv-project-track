package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Add(ctx context.Context, t domain.Task) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockTaskRepository) Get(ctx context.Context, id string) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(ctx context.Context, id string, t domain.Task) error {
	args := m.Called(ctx, id, t)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}
