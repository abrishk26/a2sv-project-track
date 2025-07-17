package data

import (
	"context"
	"errors"

	"github.com/abrishk26/a2sv-project-track/task5/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TaskRepository interface {
	Add(t models.Task) (models.Task, error)
	Get(id int) (models.Task, error)
	Delete(id int) (models.Task, error)
	Update(id int, t models.Task) (models.Task, error)
	GetAll() []models.Task
}

func NewTaskManager(db *mongo.Collection) *TaskManager {
	return &TaskManager{db}
}

type TaskManager struct {
	coll *mongo.Collection
}

func (tm *TaskManager) Add(ctx context.Context, t models.Task) (models.Task, error) {
	data := struct {
		Title       string
		Description string
		DueDate     string
		Done        bool
	}{
		t.Title, t.Description, t.DueDate, t.Done,
	}

	insertedTask, err := tm.coll.InsertOne(ctx, data)
	if err != nil {
		return models.Task{}, err
	}

	ID, ok := insertedTask.InsertedID.(bson.ObjectID)
	if !ok {
		return t, errors.New("Internal server errro")
	}

	t.ID = ID.Hex()
	return t, nil
}

func (tm *TaskManager) Get(ctx context.Context, id string) (models.Task, error) {
	var res models.Task
	bsonID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	filter := bson.D{{Key: "_id", Value: bsonID}}

	singleResult := tm.coll.FindOne(ctx, filter)
	err = singleResult.Decode(&res)
	if err != nil {
		return res, err
	}

	res.ID = id

	return res, nil
}

func (tm *TaskManager) Delete(ctx context.Context, id string) (models.Task, error) {
	var res models.Task
	res, err := tm.Get(ctx, id)
	if err != nil {
		return res, err
	}

	bsonID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}

	filter := bson.D{{Key: "_id", Value: bsonID}}
	_, err = tm.coll.DeleteOne(ctx, filter)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (tm *TaskManager) Update(ctx context.Context, id string, t models.Task) (models.Task, error) {
	var res models.Task
	_, err := tm.Get(ctx, id)
	if err != nil {
		return res, err
	}

	bsonID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return res, err
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

	_, err = tm.coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: bsonID}}, bson.D{{Key: "$set", Value: updates}})
	if err != nil {
		return res, err
	}

	res, err = tm.Get(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (tm *TaskManager) GetAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := tm.coll.Find(ctx, bson.D{})
	if err != nil {
		return tasks, err
	}

	for cursor.Next(ctx) {
		var temp struct {
			ID          bson.ObjectID `bson:"_id"`
			Title       string        `bson:"title"`
			Description string        `bson:"description"`
			DueDate     string        `bson:"due_date"`
			Done        bool          `bson:"done"`
		}

		err = cursor.Decode(&temp)
		if err != nil {
			return tasks, err
		}

		var task models.Task
		task.ID = temp.ID.Hex()
		task.Description = temp.Description
		task.Title = temp.Title
		task.DueDate = temp.DueDate
		task.Done = temp.Done

		tasks = append(tasks, task)
	}

	return tasks, nil
}
