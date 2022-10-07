package todo

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
	"gitlab.com/eliotandelon/TodoGo/models"
)

var taskModel models.TaskModel

func New() {
	taskModel = models.TaskModel{}
}

func Print() error {

	table := simpletable.New()
	tasks, err := taskModel.FindAll()
	if err != nil {
		return err
	}

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "Create At"},
			{Align: simpletable.AlignRight, Text: "Complete At"},
		},
	}

	for _, item := range tasks {
		var compleDate string
		taskName := blue(item.Name)
		taskDone := blue("no")
		if item.Done {
			compleDate = item.CompleteAt.Format(time.RFC822)
			taskName = green(fmt.Sprintf(" %s", item.Name))
			taskDone = green("yes")
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", item.Id)},
			{Text: taskName},
			{Align: simpletable.AlignCenter, Text: taskDone},
			{Text: item.CreateAt.Format(time.RFC822)},
			{Align: simpletable.AlignRight, Text: compleDate},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: "You Tasks are here"},
		},
	}

	table.SetStyle(simpletable.StyleDefault)
	table.Println()

	fmt.Println()
	return nil
}

func Add(task string) error {
	return taskModel.Create(task)
}

func Complete(index int) error {
	return taskModel.MarkComplete(index)
}

func Delete(index int) error {
	return taskModel.Delete(index)
}
