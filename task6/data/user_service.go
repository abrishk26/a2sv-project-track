package data

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task6/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserManager(db *mongo.Collection) *UserManager {
	return &UserManager{db}
}

type UserManager struct {
	coll *mongo.Collection
}

func (um *UserManager) Add(ctx context.Context, u models.User) error {
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

func (um *UserManager) Get(ctx context.Context, id string) (*models.User, error) {
	var res models.User
	singleResult := um.coll.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (um *UserManager) Delete(ctx context.Context, id string) error {
	_, err := um.coll.DeleteOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}

func (um *UserManager) Update(ctx context.Context, id string, u models.User) error {
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

func (um *UserManager) GetAll(ctx context.Context) (*[]models.User, error) {
	var tasks []models.User

	cursor, err := um.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task models.User
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
