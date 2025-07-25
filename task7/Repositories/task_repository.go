package repositories

import (
	"context"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewTaskRepository(ctx context.Context, db *mongo.Collection) *TaskRepository {
	return &TaskRepository{ctx, db}
}

type TaskRepository struct {
	ctx  context.Context
	coll *mongo.Collection
}

func (tm *TaskRepository) Add(t domain.Task) error {
	_, err := tm.coll.InsertOne(tm.ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) Get(id string) (*domain.Task, error) {
	var res domain.Task

	singleResult := tm.coll.FindOne(tm.ctx, bson.D{bson.E{Key: "_id", Value: id}})
	err := singleResult.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (tm *TaskRepository) Delete(id string) error {
	_, err := tm.Get(id)
	if err != nil {
		return err
	}

	_, err = tm.coll.DeleteOne(tm.ctx, bson.D{bson.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) Update(id string, t domain.Task) error {
	_, err := tm.Get(id)
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

	_, err = tm.coll.UpdateOne(tm.ctx, bson.D{bson.E{Key: "_id", Value: id}}, bson.D{bson.E{Key: "$set", Value: updates}})
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskRepository) GetAll() (*[]domain.Task, error) {
	var tasks []domain.Task

	cursor, err := tm.coll.Find(tm.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(tm.ctx) {
		var task domain.Task
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}
