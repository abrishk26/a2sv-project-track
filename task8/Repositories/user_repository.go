package repositories

import (
	"context"
	"errors"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{db}
}

type UserRepository struct {
	coll *mongo.Collection
}

func (ur *UserRepository) Add(ctx context.Context, u domain.User) error {
	_, err := ur.coll.InsertOne(ctx, u)
	if err != nil {
		var writeExe mongo.WriteException
		if errors.As(err, &writeExe) {
			for _, we := range writeExe.WriteErrors {
				if we.Code == 11000 {
					return domain.ErrDuplicateEmail
				}
			}
		}
		return err
	}

	return nil
}

func (ur *UserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	var res domain.User
	singleResult := ur.coll.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &res, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) error {
	res, err := ur.coll.DeleteOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, id string, u domain.User) error {
	var updates bson.D
	if u.Username != "" {
		updates = append(updates, bson.E{Key: "username", Value: u.Username})
	}
	res, err := ur.coll.UpdateOne(ctx, bson.D{bson.E{Key: "_id", Value: id}}, bson.D{bson.E{Key: "$set", Value: updates}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (ur *UserRepository) GetAll(ctx context.Context) (*[]domain.User, error) {
	var tasks []domain.User

	cursor, err := ur.coll.Find(ctx, bson.D{})
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
