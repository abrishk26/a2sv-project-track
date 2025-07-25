package repositories

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(ctx context.Context, db *mongo.Collection) *UserRepository {
	return &UserRepository{ctx, db}
}

type UserRepository struct {
	ctx  context.Context
	coll *mongo.Collection
}

func (ur *UserRepository) Add(u domain.User) error {
	userCount, err := ur.coll.CountDocuments(ur.ctx, bson.D{})
	if err != nil {
		return err
	}

	if userCount == 0 {
		u.Role = "admin"
	} else {
		u.Role = "user"
	}

	_, err = ur.coll.InsertOne(ur.ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Get(id string) (*domain.User, error) {
	var res domain.User
	singleResult := ur.coll.FindOne(ur.ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ur *UserRepository) Delete(id string) error {
	_, err := ur.coll.DeleteOne(ur.ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(id string, u domain.User) error {
	_, err := ur.Get(id)
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

	_, err = ur.coll.UpdateOne(ur.ctx, bson.D{bson.E{Key: "_id", Value: id}}, bson.D{bson.E{Key: "$set", Value: updates}})
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetAll() (*[]domain.User, error) {
	var tasks []domain.User

	cursor, err := ur.coll.Find(ur.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ur.ctx) {
		var task domain.User
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
