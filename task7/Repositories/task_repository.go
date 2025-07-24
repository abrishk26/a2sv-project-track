package repositories

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewTaskRepository(db *mongo.Collection) *TaskRepository {
	return &TaskRepository{db}
}

type TaskRepository struct {
	coll *mongo.Collection
}

func (tm *TaskRepository) Add(ctx context.Context, t domain.Task) error {
	_, err := tm.coll.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) Get(ctx context.Context, id string) (*domain.Task, error) {
	var res domain.Task

	singleResult := tm.coll.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (tm *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := tm.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = tm.coll.DeleteOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) Update(ctx context.Context, id string, t domain.Task) error {
	_, err := tm.Get(ctx, id)
	if err != nil {
		return err
	}

	var updates bson.D
	if t.Title != "" {
		updates = append(updates, bson.E{Key: "title", Value: t.Title})
	}

	if t.Description != "" {
		updates = append(updates, bson.E{Key: "description", Value: t.Description})
	}

	if t.DueDate != "" {
		updates = append(updates, bson.E{Key: "due_date", Value: t.DueDate})
	}

	if t.Done {
		updates = append(updates, bson.E{Key: "done", Value: t.Done})
	}

	_, err = tm.coll.UpdateOne(ctx, bson.D{bson.E{Key: "_id", Value: id}}, bson.D{bson.E{Key: "$set", Value: updates}})
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) GetAll(ctx context.Context) (*[]domain.Task, error) {
	var tasks []domain.Task

	cursor, err := tm.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task domain.Task
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
