package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/abrishk26/a2sv-project-track/task8/Repositories"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupMongo(t *testing.T) (*mongo.Client, *mongo.Collection, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)

	err = client.Ping(ctx, nil)
	require.NoError(t, err)

	db := client.Database("task_manager_test")
	coll := db.Collection("users")

	err = coll.Drop(ctx)
	require.NoError(t, err)

	cleanup := func() {
		_ = db.Drop(context.Background())
		_ = client.Disconnect(context.Background())
		cancel()
	}

	return client, coll, cleanup
}

func TestUserRepository_CRUD(t *testing.T) {
	_, coll, cleanup := setupMongo(t)
	defer cleanup()

	repo := repositories.NewUserRepository(coll)

	ctx := context.Background()

	user := domain.User{
		ID:           primitive.NewObjectID().Hex(),
		Username:     "testuser",
		IsAdmin:      false,
		PasswordHash: "hashedpassword",
	}

	// Test Add
	err := repo.Add(ctx, user)
	require.NoError(t, err)

	// Test GetByUsername
	fetched, err := repo.GetByUsername(ctx, "testuser")
	require.NoError(t, err)
	assert.Equal(t, user.ID, fetched.ID)
	assert.Equal(t, user.Username, fetched.Username)

	// Test GetByID
	fetchedByID, err := repo.GetByID(ctx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.Username, fetchedByID.Username)

	// Test Update
	user.Username = "updateduser"
	err = repo.Update(ctx, user.ID, user)
	require.NoError(t, err)

	updated, err := repo.GetByID(ctx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, "updateduser", updated.Username)

	// Test GetAll
	allUsers, err := repo.GetAll(ctx)
	require.NoError(t, err)
	assert.Len(t, allUsers, 1)

	// Test Delete
	err = repo.Delete(ctx, user.ID)
	require.NoError(t, err)

	_, err = repo.GetByID(ctx, user.ID)
	assert.ErrorIs(t, err, domain.ErrUserNotFound)
}
