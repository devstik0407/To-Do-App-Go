package todos

import (
	"errors"
	"fmt"
)

type Task struct {
	Id   int64  `json:"id"`
	Desc string `json:"desc"`
}

func (t Task) show() {
	fmt.Printf("id: %d, description: %s", t.Id, t.Desc)
}

func AddTask(listId int64, desc string) (Task, error) {
	taskList, err := getTaskList(listId)
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		Id:   newTaskId(taskList),
		Desc: desc,
	}
	taskList.Tasks = append(taskList.Tasks, newTask)
	return newTask, nil
}

func DeleteTask(listId int64, taskId int64) (Task, error) {
	_, err := getTask(listId, taskId)
	if err != nil {
		return Task{}, err
	}

	taskList, _ := getTaskList(listId)
	index := 0
	task := Task{}
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == taskId {
			index = i
			task = taskList.Tasks[i]
		}
	}
	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
	return task, nil
}

func UpdateTask(listId, taskId int64, newDesc string) (Task, error) {
	task, err := getTask(listId, taskId)
	if err != nil {
		return Task{}, err
	}
	task.Desc = newDesc
	return *task, nil
}

func getTask(listId, taskId int64) (*Task, error) {
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

func newTaskId(taskList *TaskList) int64 {
	if len(taskList.Tasks) == 0 {
		return 1
	}
	lastTask := taskList.Tasks[len(taskList.Tasks)-1]
	return lastTask.Id + 1
}
