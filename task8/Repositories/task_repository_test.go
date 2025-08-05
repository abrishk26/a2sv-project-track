package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

func setupTestDB(t *testing.T) (*mongo.Client, *mongo.Collection, func()) {
	t.Helper()

	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)

	db := client.Database("task_manager_test")
	coll := db.Collection("tasks")

	// Clean before each test
	err = coll.Drop(ctx)
	require.NoError(t, err)

	cleanup := func() {
		_ = coll.Drop(ctx)
		_ = client.Disconnect(ctx)
	}

	return client, coll, cleanup
}

func TestTaskRepository_CRUD(t *testing.T) {
	_, coll, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(coll)
	ctx := context.Background()

	// Test Add
	task := domain.Task{
		ID:          "task1",
		UserID:      "user1",
		Title:       "Test Task",
		Description: "Testing task repo",
		DueDate:     time.Now().Format(time.RFC3339),
		Done:        false,
	}
	err := repo.Add(ctx, task)
	require.NoError(t, err)

	// Test Get
	got, err := repo.Get(ctx, "task1")
	require.NoError(t, err)
	require.Equal(t, "Test Task", got.Title)

	// Test Update
	updated := domain.Task{
		Title:       "Updated Task",
		Description: "Updated description",
		Done:        true,
	}
	err = repo.Update(ctx, "task1", updated)
	require.NoError(t, err)

	// Verify update
	got, err = repo.Get(ctx, "task1")
	require.NoError(t, err)
	require.Equal(t, "Updated Task", got.Title)
	require.Equal(t, true, got.Done)

	// Test GetAll
	task2 := domain.Task{
		ID:          "task2",
		UserID:      "user2",
		Title:       "Another Task",
		Description: "Another one",
		DueDate:     time.Now().Format(time.RFC3339),
		Done:        false,
	}
	err = repo.Add(ctx, task2)
	require.NoError(t, err)

	tasks, err := repo.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, tasks, 2)

	// Test Delete
	err = repo.Delete(ctx, "task1")
	require.NoError(t, err)

	_, err = repo.Get(ctx, "task1")
	require.ErrorIs(t, err, domain.ErrTaskNotFound)
}
