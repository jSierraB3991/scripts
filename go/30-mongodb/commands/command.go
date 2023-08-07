package commands

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jsierrab3991/scripts/mongodb/model"
	"github.com/jsierrab3991/scripts/mongodb/repository"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gookit/color.v1"
)

type Command struct {
	repo repository.Repository
}

func NewCommand(collection *mongo.Collection, ctx context.Context) *Command {
	return &Command{
		repo: *repository.NewRepository(collection, ctx),
	}
}

func (comm *Command) AddTask(cCtx *cli.Context) error {
	str := cCtx.Args().First()
	if str == "" {
		return errors.New("Cannot add an empty task")
	}

	task := &model.Task{
		ID:        primitive.NewObjectID(),
		CreateAt:  time.Now(),
		Text:      str,
		Completed: false,
	}
	return comm.repo.CreateTask(task)
}

func (comm *Command) GetPending(cCtx *cli.Context) error {
	tasks, err := comm.repo.GetPending()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Nothing to see here.\nRun `add` task to add a task")
		}
		return err
	}
	comm.PrintTasks(tasks)
	return nil
}

func (comm *Command) GetAll(cCtx *cli.Context) error {
	tasks, err := comm.repo.GetAll()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Nothing to see here")
			return nil
		}
		return err
	}
	comm.PrintTasks(tasks)
	return nil
}

func (comm Command) PrintTasks(tasks []*model.Task) {
	for i, v := range tasks {
		if v.Completed {
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else {
			color.Yellow.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}

func (comm *Command) DoneTask(cCtx *cli.Context) error {
	text := cCtx.Args().First()
	return comm.repo.CompleteTask(text)
}
