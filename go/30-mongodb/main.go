package main

import (
	"context"
	"log"
	"os"

	"github.com/jsierrab3991/scripts/mongodb/commands"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("task")
}

func getSubCommands(commands *commands.Command) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task to the list",
			Action:  commands.AddTask,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all tasks",
			Action:  commands.GetAll,
		},
		{
			Name:    "done",
			Aliases: []string{"d"},
			Usage:   "Complete a task on the list",
			Action:  commands.DoneTask,
		},
		{
			Name:    "pending",
			Aliases: []string{"p"},
			Action:  commands.GetPending,
		},
	}
}

func main() {

	commands := commands.NewCommand(collection, ctx)
	subCommands := getSubCommands(commands)

	app := &cli.App{
		Name:     "tasker",
		Usage:    "A simple CLI program to manage your tasks",
		Commands: subCommands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
