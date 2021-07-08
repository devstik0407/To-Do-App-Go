package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/todos"

	"github.com/gorilla/mux"
)

func TestAddTaskSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			AddTaskFunc: func(listId int64, desc string) (todos.Task, error) {
				return todos.Task{Id: 1, Desc: "code clean"}, nil
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("POST", fmt.Sprintf("/todos/%d", 1), bytes.NewBuffer([]byte(`{"desc":"code clean"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}", mockHandler.AddTask)
	rrA := httptest.NewRecorder()
	r.ServeHTTP(rrA, req)

	if rrA.Code != 201 {
		t.Errorf("actual: %v expected: %v", rrA.Code, 201)
	}

	expectedBody := `{"status":"successfully added task","error":"","task":{"id":1,"desc":"code clean"}}` + "\n"
	if rrA.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrA.Body, expectedBody)
	}
}

func TestAddTaskFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			AddTaskFunc: func(listId int64, desc string) (todos.Task, error) {
				return todos.Task{}, errors.New("max limit of tasks exceeded")
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("POST", fmt.Sprintf("/todos/%d", 1), bytes.NewBuffer([]byte(`{"desc":"code clean"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}", mockHandler.AddTask)
	rrA := httptest.NewRecorder()
	r.ServeHTTP(rrA, req)

	if rrA.Code != 406 {
		t.Errorf("actual: %v expected: %v", rrA.Code, 201)
	}

	expectedBody := `{"status":"failed","error":"max limit of tasks exceeded","task":{"id":0,"desc":""}}` + "\n"
	if rrA.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrA.Body, expectedBody)
	}
}

func TestUpdateTaskSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			UpdateTaskFunc: func(listId, taskId int64, newDesc string) (todos.Task, error) {
				return todos.Task{Id: taskId, Desc: newDesc}, nil
			},
		},
	}

	r := mux.NewRouter()

	taskId := int64(1)
	req, err := http.NewRequest("UPDATE", fmt.Sprintf("/todos/%d/%d", 1, taskId), bytes.NewBuffer([]byte(`{"desc":"sleep"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}/{taskId}", mockHandler.UpdateTask)
	rrU := httptest.NewRecorder()
	r.ServeHTTP(rrU, req)

	if rrU.Code != 201 {
		t.Errorf("actual: %v expected: %v", rrU.Code, 201)
	}

	expectedBody := `{"status":"successfully updated task","error":"","task":{"id":1,"desc":"sleep"}}` + "\n"
	if rrU.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrU.Body, expectedBody)
	}
}

func TestUpdateTaskFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			UpdateTaskFunc: func(listId, taskId int64, newDesc string) (todos.Task, error) {
				return todos.Task{}, errors.New("unable to update in database")
			},
		},
	}

	r := mux.NewRouter()

	taskId := int64(1)
	req, err := http.NewRequest("UPDATE", fmt.Sprintf("/todos/%d/%d", 1, taskId), bytes.NewBuffer([]byte(`{"desc":"sleep"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}/{taskId}", mockHandler.UpdateTask)
	rrU := httptest.NewRecorder()
	r.ServeHTTP(rrU, req)

	if rrU.Code != 406 {
		t.Errorf("actual: %v expected: %v", rrU.Code, 406)
	}

	expectedBody := `{"status":"failed","error":"unable to update in database","task":{"id":0,"desc":""}}` + "\n"
	if rrU.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrU.Body, expectedBody)
	}
}

func TestDeleteTaskSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			DeleteTaskFunc: func(listId, taskId int64) (todos.Task, error) {
				return todos.Task{Id: taskId, Desc: "drink water"}, nil
			},
		},
	}

	r := mux.NewRouter()

	taskId := int64(1)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/todos/%d/%d", 1, taskId), nil)
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}/{taskId}", mockHandler.DeleteTask)
	rrD := httptest.NewRecorder()
	r.ServeHTTP(rrD, req)

	if rrD.Code != 200 {
		t.Errorf("actual: %v expected: %v", rrD.Code, 200)
	}

	expectedBody := `{"status":"successfully deleted task","error":"","task":{"id":1,"desc":"drink water"}}` + "\n"
	if rrD.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrD.Body, expectedBody)
	}
}

func TestDeleteTaskFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			DeleteTaskFunc: func(listId, taskId int64) (todos.Task, error) {
				return todos.Task{}, errors.New("unable to delete from database")
			},
		},
	}

	r := mux.NewRouter()

	taskId := int64(1)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/todos/%d/%d", 1, taskId), nil)
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}/{taskId}", mockHandler.DeleteTask)
	rrD := httptest.NewRecorder()
	r.ServeHTTP(rrD, req)

	if rrD.Code != 406 {
		t.Errorf("actual: %v expected: %v", rrD.Code, 406)
	}

	expectedBody := `{"status":"failed","error":"unable to delete from database","task":{"id":0,"desc":""}}` + "\n"
	if rrD.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrD.Body, expectedBody)
	}
}

type mockTodoService struct {
	AddTaskFunc        func(listId int64, desc string) (todos.Task, error)
	GetTaskFunc        func(listId, taskId int64) (todos.Task, error)
	UpdateTaskFunc     func(listId, taskId int64, newDesc string) (todos.Task, error)
	DeleteTaskFunc     func(listId, taskId int64) (todos.Task, error)
	CreateTaskListFunc func(title string) (int64, error)
	GetTaskListFunc    func(listId int64) (todos.TaskList, error)
	DeleteTaskListFunc func(listId int64) (int64, error)
	GetTodosFunc       func() ([]todos.TaskList, error)
	NewTaskListIdFunc  func() (int64, error)
	NewTaskIdFunc      func(listId int64) (int64, error)
}

func (mts mockTodoService) AddTask(listId int64, desc string) (todos.Task, error) {
	return mts.AddTaskFunc(listId, desc)
}

func (mts mockTodoService) GetTask(listId, taskId int64) (todos.Task, error) {
	return mts.GetTaskFunc(listId, taskId)
}

func (mts mockTodoService) UpdateTask(listId, taskId int64, newDesc string) (todos.Task, error) {
	return mts.UpdateTaskFunc(listId, taskId, newDesc)
}

func (mts mockTodoService) DeleteTask(listId, taskId int64) (todos.Task, error) {
	return mts.DeleteTaskFunc(listId, taskId)
}

func (mts mockTodoService) CreateTaskList(title string) (int64, error) {
	return mts.CreateTaskListFunc(title)
}

func (mts mockTodoService) GetTaskList(listId int64) (todos.TaskList, error) {
	return mts.GetTaskListFunc(listId)
}

func (mts mockTodoService) DeleteTaskList(listId int64) (int64, error) {
	return mts.DeleteTaskListFunc(listId)
}

func (mts mockTodoService) GetTodos() ([]todos.TaskList, error) {
	return mts.GetTodosFunc()
}

func (mts mockTodoService) NewTaskListId() (int64, error) {
	return mts.NewTaskListIdFunc()
}

func (mts mockTodoService) NewTaskId(listId int64) (int64, error) {
	return mts.NewTaskIdFunc(listId)
}
