package mongostore

import (
	"context"
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
