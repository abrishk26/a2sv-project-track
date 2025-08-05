package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/abrishk26/a2sv-project-track/task8/mocks"
	"github.com/abrishk26/a2sv-project-track/task8/Usecases"
)

func TestTaskUsecases_Add_Success(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenSvc := new(mocks.MockTokenService)

	us := usecases.NewTaskUsecases(mockTaskRepo, mockUserRepo, mockTokenSvc)

	ctx := context.Background()
	token := "token123"
	ctx = usecases.ContextWithToken(ctx, token)

	user := &domain.User{ID: "u1", IsAdmin: false}
	task := domain.Task{Title: "Task1"}

	mockTokenSvc.On("VerifyToken", token).Return(user.ID, nil)
	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil)
	mockTaskRepo.On("Add", mock.Anything, mock.MatchedBy(func(t domain.Task) bool {
		return t.UserID == user.ID && t.Title == "Task1"
	})).Return(nil)

	err := us.Add(ctx, task)
	assert.NoError(t, err)

	mockTokenSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockTaskRepo.AssertExpectations(t)
}

func TestTaskUsecases_Get_AccessControl(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenSvc := new(mocks.MockTokenService)

	us := usecases.NewTaskUsecases(mockTaskRepo, mockUserRepo, mockTokenSvc)

	ctx := context.Background()
	token := "token123"
	ctx = usecases.ContextWithToken(ctx, token)

	task := &domain.Task{ID: "t1", UserID: "u1", Title: "Test Task"}

	// Case 1: Admin user can get any task
	adminUser := &domain.User{ID: "admin", IsAdmin: true}

	mockTokenSvc.On("VerifyToken", token).Return(adminUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, adminUser.ID).Return(adminUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()

	gotTask, err := us.Get(ctx, "t1")
	assert.NoError(t, err)
	assert.Equal(t, task, gotTask)

	// Case 2: Normal user getting own task - allowed
	user := &domain.User{ID: "u1", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(user.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()

	gotTask, err = us.Get(ctx, "t1")
	assert.NoError(t, err)
	assert.Equal(t, task, gotTask)

	// Case 3: Normal user trying to get task of another user - forbidden
	otherUser := &domain.User{ID: "u2", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(otherUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, otherUser.ID).Return(otherUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()

	gotTask, err = us.Get(ctx, "t1")
	assert.ErrorIs(t, err, domain.ErrUnauthorized)
	assert.Nil(t, gotTask)

	mockTokenSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockTaskRepo.AssertExpectations(t)
}

func TestTaskUsecases_Delete_AccessControl(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenSvc := new(mocks.MockTokenService)

	us := usecases.NewTaskUsecases(mockTaskRepo, mockUserRepo, mockTokenSvc)
	ctx := context.Background()
	token := "token123"
	ctx = usecases.ContextWithToken(ctx, token)

	task := &domain.Task{ID: "t1", UserID: "u1"}

	// Admin user can delete any task
	adminUser := &domain.User{ID: "admin", IsAdmin: true}

	mockTokenSvc.On("VerifyToken", token).Return(adminUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, adminUser.ID).Return(adminUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()
	mockTaskRepo.On("Delete", mock.Anything, "t1").Return(nil).Once()

	err := us.Delete(ctx, "t1")
	assert.NoError(t, err)

	// Normal user deleting own task allowed
	user := &domain.User{ID: "u1", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(user.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()
	mockTaskRepo.On("Delete", mock.Anything, "t1").Return(nil).Once()

	err = us.Delete(ctx, "t1")
	assert.NoError(t, err)

	// Normal user deleting another user's task forbidden
	otherUser := &domain.User{ID: "u2", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(otherUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, otherUser.ID).Return(otherUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()

	err = us.Delete(ctx, "t1")
	assert.ErrorIs(t, err, domain.ErrUnauthorized)

	mockTokenSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockTaskRepo.AssertExpectations(t)
}

func TestTaskUsecases_Update_AccessControl(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenSvc := new(mocks.MockTokenService)

	us := usecases.NewTaskUsecases(mockTaskRepo, mockUserRepo, mockTokenSvc)
	ctx := context.Background()
	token := "token123"
	ctx = usecases.ContextWithToken(ctx, token)

	task := &domain.Task{ID: "t1", UserID: "u1"}

	updatedTask := domain.Task{Title: "Updated"}

	// Admin user can update any task
	adminUser := &domain.User{ID: "admin", IsAdmin: true}

	mockTokenSvc.On("VerifyToken", token).Return(adminUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, adminUser.ID).Return(adminUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()
	mockTaskRepo.On("Update", mock.Anything, "t1", updatedTask).Return(nil).Once()

	err := us.Update(ctx, "t1", updatedTask)
	assert.NoError(t, err)

	// Normal user updating own task allowed
	user := &domain.User{ID: "u1", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(user.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()
	mockTaskRepo.On("Update", mock.Anything, "t1", updatedTask).Return(nil).Once()

	err = us.Update(ctx, "t1", updatedTask)
	assert.NoError(t, err)

	// Normal user updating another user's task forbidden
	otherUser := &domain.User{ID: "u2", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(otherUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, otherUser.ID).Return(otherUser, nil).Once()
	mockTaskRepo.On("Get", mock.Anything, "t1").Return(task, nil).Once()

	err = us.Update(ctx, "t1", updatedTask)
	assert.ErrorIs(t, err, domain.ErrUnauthorized)

	mockTokenSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockTaskRepo.AssertExpectations(t)
}

func TestTaskUsecases_GetAll_AdminOnly(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenSvc := new(mocks.MockTokenService)

	us := usecases.NewTaskUsecases(mockTaskRepo, mockUserRepo, mockTokenSvc)
	ctx := context.Background()
	token := "token123"
	ctx = usecases.ContextWithToken(ctx, token)

	adminUser := &domain.User{ID: "admin", IsAdmin: true}
	tasks := []domain.Task{
		{ID: "t1", Title: "Task1"},
		{ID: "t2", Title: "Task2"},
	}

	mockTokenSvc.On("VerifyToken", token).Return(adminUser.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, adminUser.ID).Return(adminUser, nil).Once()
	mockTaskRepo.On("GetAll", mock.Anything).Return(tasks, nil).Once()

	gotTasks, err := us.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, tasks, gotTasks)

	// Non-admin user should get unauthorized error
	user := &domain.User{ID: "u1", IsAdmin: false}

	mockTokenSvc.On("VerifyToken", token).Return(user.ID, nil).Once()
	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil).Once()

	gotTasks, err = us.GetAll(ctx)
	assert.ErrorIs(t, err, domain.ErrUnauthorized)
	assert.Nil(t, gotTasks)

	mockTokenSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockTaskRepo.AssertExpectations(t)
}
