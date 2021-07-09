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

func TestDeleteTaskListSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	err := md.DeleteTaskList(1)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTaskListFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	err := md.DeleteTaskList(1)
	if err == nil {
		t.Errorf("actual error: none\nexpected error: %v", errors.New("invalid listId"))
	}
	if err.Error() != "invalid listId" {
		t.Errorf("actual error: %v\nexpected error: %v", err, errors.New("invalid listId"))
	}
}

func TestGetTodos(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)

	err := md.CreateTaskList(todos.TaskList{Id: 1, Title: "Monday", Tasks: nil})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = md.DeleteTaskList(1); err != nil {
			t.Fatal(err)
		}
	}()
	err = md.CreateTaskList(todos.TaskList{Id: 2, Title: "Tuesday", Tasks: nil})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = md.DeleteTaskList(2); err != nil {
			t.Fatal(err)
		}
	}()

	actualTodos, err := md.GetTodos()
	if err != nil {
		t.Error(err)
	}
	expectedTodos := []todos.TaskList{
		{Id: 1, Title: "Monday", Tasks: nil},
		{Id: 2, Title: "Tuesday", Tasks: nil},
	}

	for i := 0; i < 1; i++ {
		if actualTodos[i].Id != expectedTodos[i].Id || actualTodos[i].Title != expectedTodos[i].Title || len(actualTodos[i].Tasks) != len(expectedTodos[i].Tasks) {
			t.Errorf("actual: %v\nexpected: %v", actualTodos, expectedTodos)
		}
	}
}

func TestAddTaskSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	err := md.CreateTaskList(todos.TaskList{Id: 1, Title: "Monday", Tasks: []todos.Task{}})
	if err != nil {
		t.Fatal(err)
	}

	err = md.AddTask(1, todos.Task{Id: 1, Desc: "watch movies"})
	if err != nil {
		t.Error(err)
	}
}

func TestAddTaskFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)

	err := md.AddTask(2, todos.Task{Id: 1, Desc: "watch movies"})
	if err == nil {
		t.Errorf("actual error: none\nexpected error: %v", errors.New("invalid listId"))
	}
	if err.Error() != "invalid listId" {
		t.Errorf("actual error: %v\nexpected error: %v", err, errors.New("invalid listId"))
	}
}

func TestGetTaskSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	actualTask, err := md.GetTask(1, 1)
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := todos.Task{
		Id:   1,
		Desc: "watch movies",
	}
	if actualTask != expectedTask {
		t.Errorf("actual: %v\nexpected: %v", actualTask, expectedTask)
	}
}

func TestGetTaskFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	md := Connect(ctx)
	_, err := md.GetTask(1, 2)
	if err == nil {
		t.Errorf("actual error: none\nexpected error: %v", errors.New("invalid taskId"))
	}
	if err.Error() != "invalid taskId" {
		t.Errorf("actual error: %v\nexpected error: %v", err, errors.New("invalid taskId"))
	}
}
