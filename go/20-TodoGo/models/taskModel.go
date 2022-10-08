package models

import (
	"log"
	"time"

	"gitlab.com/eliotandelon/TodoGo/config"
	"gitlab.com/eliotandelon/TodoGo/entities"
)

type TaskModel struct{}

func (*TaskModel) FindAll() ([]entities.Task, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT id, name, done, create_at, complete_at FROM task")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id, &task.Name, &task.Done, task.CreateAt, task.CompleteAt)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (this *TaskModel) Create(name string) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO task(name, done, create_at) VALUES($1, false, $2)", name, time.Now().UTC().Format(time.RFC822))
	if err != nil {
		return err
	}
	return nil
}

func (*TaskModel) MarkComplete(id int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE task SET done = true, complete_at = $1 WHERE id =$2", time.Now().UTC().Format(time.RFC822), id)
	if err != nil {
		return err
	}
	return nil
}

func (*TaskModel) Delete(id int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM task WHERE id =$1", id)
	if err != nil {
		return err
	}
	return nil
}
