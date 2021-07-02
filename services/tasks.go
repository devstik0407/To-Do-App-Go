package services

import (
	"errors"
	"todo/models"
)

func AddTask(listId int64, desc string) error {
	taskList, err := getTaskList(listId)
	if err != nil {
		return err
	}

	newTask := models.Task{
		Id:   newTaskId(taskList),
		Desc: desc,
	}
	taskList.Tasks = append(taskList.Tasks, newTask)
	return nil
}

func DeleteTask(listId int64, taskId int64) error {
	_, err := getTask(listId, taskId)
	if err != nil {
		return err
	}

	taskList, _ := getTaskList(listId)
	index := 0
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == taskId {
			index = i
		}
	}
	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
	return nil
}

func UpdateTask(listId, taskId int64, newDesc string) error {
	task, err := getTask(listId, taskId)
	if err != nil {
		return err
	}
	task.Desc = newDesc
	return nil
}

func getTask(listId, taskId int64) (*models.Task, error) {
	taskList, err := getTaskList(listId)
	if err != nil {
		return nil, err
	}
	for i, task := range taskList.Tasks {
		if task.Id == taskId {
			return &taskList.Tasks[i], nil
		}
	}
	return nil, errors.New("invalid task id")
}

func newTaskId(taskList *models.TaskList) int64 {
	if len(taskList.Tasks) == 0 {
		return 1
	}
	lastTask := taskList.Tasks[len(taskList.Tasks)-1]
	return lastTask.Id + 1
}
