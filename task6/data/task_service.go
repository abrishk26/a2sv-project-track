package data

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task6/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewTaskManager(db *mongo.Collection) *TaskManager {
	return &TaskManager{db}
}

type TaskManager struct {
	coll *mongo.Collection
}

func (tm *TaskManager) Add(ctx context.Context, t models.Task) error {
	_, err := tm.coll.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskManager) Get(ctx context.Context, id string) (*models.Task, error) {
	var res models.Task

	singleResult := tm.coll.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (tm *TaskManager) Delete(ctx context.Context, id string) error {
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

func (tm *TaskManager) Update(ctx context.Context, id string, t models.Task) error {
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

func (tm *TaskManager) GetAll(ctx context.Context) (*[]models.Task, error) {
	var tasks []models.Task

	cursor, err := tm.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task models.Task
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
