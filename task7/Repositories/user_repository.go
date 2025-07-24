package repositories

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{db}
}

type UserRepository struct {
	coll *mongo.Collection
}

func (um *UserRepository) Add(ctx context.Context, u domain.User) error {
	userCount, err := um.coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return err
	}

	if userCount == 0 {
		u.Role = "admin"
	} else {
		u.Role = "user"
	}

	_, err = um.coll.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (um *UserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	var res domain.User
	singleResult := um.coll.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (um *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := um.coll.DeleteOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}

func (um *UserRepository) Update(ctx context.Context, id string, u domain.User) error {
	_, err := um.Get(ctx, id)
	if err != nil {
		return err
	}

	var updates bson.D
	if u.Username != "" {
		updates = append(updates, bson.E{Key: "username", Value: u.Username})
	}

	if u.Role != "" {
		updates = append(updates, bson.E{Key: "role", Value: u.Role})
	}

	_, err = um.coll.UpdateOne(ctx, bson.D{bson.E{Key: "_id", Value: id}}, bson.D{bson.E{Key: "$set", Value: updates}})
	if err != nil {
		return err
	}

	return nil
}

func (um *UserRepository) GetAll(ctx context.Context) (*[]domain.User, error) {
	var tasks []domain.User

	cursor, err := um.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task domain.User
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
