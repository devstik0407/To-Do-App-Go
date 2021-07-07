package todos

import (
	"errors"
	"todo/models"
)

func AddTask(listId int64, desc string) (models.Task, error) {
	taskList, err := getTaskList(listId)
	if err != nil {
		return models.Task{}, err
	}

	newTask := models.Task{
		Id:   newTaskId(taskList),
		Desc: desc,
	}
	taskList.Tasks = append(taskList.Tasks, newTask)
	return newTask, nil
}

func DeleteTask(listId int64, taskId int64) (models.Task, error) {
	_, err := getTask(listId, taskId)
	if err != nil {
		return models.Task{}, err
	}

	taskList, _ := getTaskList(listId)
	index := 0
	task := models.Task{}
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == taskId {
			index = i
			task = taskList.Tasks[i]
		}
	}
	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
	return task, nil
}

func UpdateTask(listId, taskId int64, newDesc string) (models.Task, error) {
	task, err := getTask(listId, taskId)
	if err != nil {
		return models.Task{}, err
	}
	task.Desc = newDesc
	return *task, nil
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
