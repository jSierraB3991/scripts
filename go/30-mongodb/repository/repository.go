package repository

import (
	"context"
	"errors"

	"github.com/jsierrab3991/scripts/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRepository(collection *mongo.Collection, ctx context.Context) *Repository {
	return &Repository{
		collection: collection,
		ctx:        ctx,
	}
}

func (repo *Repository) CreateTask(task *model.Task) error {
	_, err := repo.collection.InsertOne(repo.ctx, task)
	return err
}

func (repo *Repository) GetAll() ([]*model.Task, error) {
	filter := bson.D{{}}
	return repo.filterTasks(filter)
}

func (repo *Repository) GetPending() ([]*model.Task, error) {
	filter := bson.D{
		primitive.E{Key: "completed", Value: false},
	}
	return repo.filterTasks(filter)
}

func (repo *Repository) filterTasks(filter interface{}) ([]*model.Task, error) {
	var tasks []*model.Task

	cur, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(repo.ctx) {
		var t model.Task
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}
	cur.Close(repo.ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func (repo *Repository) CompleteTask(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &model.Task{}
	return repo.collection.FindOneAndUpdate(repo.ctx, filter, update).Decode(t)
}

func (repo *Repository) DeleteTask(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	res, err := repo.collection.DeleteOne(repo.ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No tasks were deleted")
	}
	return nil
}
