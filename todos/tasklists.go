package todos

import (
	"errors"
	"fmt"
)

var Todos []TaskList

type TaskList struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

func (tl TaskList) show() {
	fmt.Printf("id: %d, title: %s, tasks: %v\n", tl.Id, tl.Title, tl.Tasks)
}

func CreateTaskList(title string) (int64, error) {
	taskList := TaskList{
		Id:    NewTaskListId(),
		Title: title,
		Tasks: []Task{},
	}
	Todos = append(Todos, taskList)
	return taskList.Id, nil
}

func DeleteTaskList(listId int64) (int64, error) {
	_, err := getTaskList(listId)
	if err != nil {
		return 0, err
	}

	index := 0
	for i := range Todos {
		if Todos[i].Id == listId {
			index = i
		}
	}
	Todos = append(Todos[:index], Todos[index+1:]...)
	return listId, nil
}

func getTaskList(listId int64) (*TaskList, error) {
	for i, taskList := range Todos {
		if taskList.Id == listId {
			return &Todos[i], nil
		}
	}
	return nil, errors.New("invalid task-list id")
}

func NewTaskListId() int64 {
	if len(Todos) == 0 {
		return 1
	}
	lastTaskList := Todos[len(Todos)-1]
	return lastTaskList.Id + 1
}
