package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "gitlab.com/eliotandelon/TodoGo"
)

func tryError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("Empty tasj is not allowed")
	}
	return text, nil
}
func main() {
	addFlag := flag.Bool("add", false, "Add a new task")
	listFlag := flag.Bool("list", false, "List of the all Tasks")
	completeFlag := flag.Int("complete", 0, "Mark a task as complete")
	deleteFlag := flag.Int("delete", 0, "Delte a task")

	flag.Parse()
	todo.New()

	switch {
	case *addFlag:
		taskName, err := getInput(os.Stdin, flag.Args()...)
		tryError(err)
		err = todo.Add(taskName)
		tryError(err)
	case *completeFlag > 0:
		err := todo.Complete(*completeFlag)
		tryError(err)
	case *deleteFlag > 0:
		err := todo.Delete(*deleteFlag)
		tryError(err)
	case *listFlag:
		err := todo.Print()
		tryError(err)
	default:
		fmt.Fprintln(os.Stdout, "Invalid Param")
		os.Exit(1)
	}

}
