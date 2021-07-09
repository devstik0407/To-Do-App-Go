package mongostore

import (
	"context"
	"errors"
	"testing"
	"todo/todos"
)

func TestCreateTaskList(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	newTaskList := todos.TaskList{
		Id:    1,
		Title: "Monday",
		Tasks: []todos.Task{},
	}

	err := md.CreateTaskList(newTaskList)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTaskListSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	actualTaskList, err := md.GetTaskList(1)
	if err != nil {
		t.Error(err)
	}
	expectedTaskList := todos.TaskList{
		Id:    1,
		Title: "Monday",
		Tasks: []todos.Task{},
	}
	if actualTaskList.Id != expectedTaskList.Id || actualTaskList.Title != expectedTaskList.Title {
		t.Errorf("actual: %v\nexpected: %v", actualTaskList, expectedTaskList)
	}
}

func TestGetTaskListFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	_, err := md.GetTaskList(2)
	if err == nil {
		t.Errorf("actual error: none\nexpected error: %v", errors.New("invalid listId"))
	}
	if err.Error() != "invalid listId" {
		t.Errorf("actual error: %v\nexpected error: %v", err, errors.New("invalid listId"))
	}
}
