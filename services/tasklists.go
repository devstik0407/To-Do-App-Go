package services

import (
	"errors"
	"todo/models"
)

var Todos []models.TaskList

func CreateTaskList(title string) error {
	taskList := models.TaskList{
		Id:    newTaskListId(),
		Title: title,
		Tasks: []models.Task{},
	}
	Todos = append(Todos, taskList)
	return nil
}

func DeleteTaskList(listId int64) error {
	_, err := getTaskList(listId)
	if err != nil {
		return err
	}

	index := 0
	for i := range Todos {
		if Todos[i].Id == listId {
			index = i
		}
	}
	Todos = append(Todos[:index], Todos[index+1:]...)
	return nil
}

func getTaskList(listId int64) (*models.TaskList, error) {
	for i, taskList := range Todos {
		if taskList.Id == listId {
			return &Todos[i], nil
		}
	}
	return nil, errors.New("invalid task-list id")
}

func newTaskListId() int64 {
	if len(Todos) == 0 {
		return 1
	}
	lastTaskList := Todos[len(Todos)-1]
	return lastTaskList.Id + 1
}
