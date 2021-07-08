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

func TestCreateTaskListSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			CreateTaskListFunc: func(title string) (int64, error) {
				return 1, nil
			},
			NewTaskListIdFunc: func() (int64, error) {
				return 1, nil
			},
		},
	}
	testBody := []byte(`{"title":"Monday"}`)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}

	listId, err := mockHandler.Service.NewTaskListId()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateTaskList)
	handler.ServeHTTP(rr, req)
	expectedStatusCode := 201
	if status := rr.Code; status != expectedStatusCode {
		t.Errorf("actual: %v expected: %v", status, expectedStatusCode)
	}

	expectedBody := fmt.Sprintf(`{"status":"successfully created task-list","error":"","listId":%d}`, listId) + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rr.Body.String(), expectedBody)
	}
}

func TestCreateTaskListFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			CreateTaskListFunc: func(title string) (int64, error) {
				return -1, errors.New("database full")
			},
		},
	}
	testBody := []byte(`{"title":"Monday"}`)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateTaskList)
	handler.ServeHTTP(rr, req)

	expectedStatusCode := 500
	if status := rr.Code; status != expectedStatusCode {
		t.Errorf("actual: %v expected: %v", status, expectedStatusCode)
	}

	expectedBody := `{"status":"failed","error":"database full","listId":0}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rr.Body.String(), expectedBody)
	}
}

func TestDeleteTaskListSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			DeleteTaskListFunc: func(listId int64) (int64, error) {
				return 1, nil
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/todos/%d", 1), nil)
	if err != nil {
		t.Fatal(err)
	}

	rrD := httptest.NewRecorder()
	r.HandleFunc("/todos/{listId}", mockHandler.DeleteTaskList)
	r.ServeHTTP(rrD, req)

	expectedStatusCode := 200
	if status := rrD.Code; status != expectedStatusCode {
		t.Errorf("actual: %v expected: %v", status, expectedStatusCode)
	}

	expectedBody := fmt.Sprintf(`{"status":"successfully deleted task-list","error":"","listId":%d}`, 1) + "\n"
	if rrD.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrD.Body.String(), expectedBody)
	}
}

func TestDeleteTaskListFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			DeleteTaskListFunc: func(listId int64) (int64, error) {
				return 0, errors.New("invalid listId")
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/todos/%d", 1), nil)
	if err != nil {
		t.Fatal(err)
	}

	rrD := httptest.NewRecorder()
	r.HandleFunc("/todos/{listId}", mockHandler.DeleteTaskList)
	r.ServeHTTP(rrD, req)

	expectedStatusCode := 406
	if status := rrD.Code; status != expectedStatusCode {
		t.Errorf("actual: %v expected: %v", status, expectedStatusCode)
	}

	expectedBody := `{"status":"failed","error":"invalid listId","listId":0}` + "\n"
	if rrD.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrD.Body.String(), expectedBody)
	}
}

func TestGetTodosSuccess(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			GetTodosFunc: func() ([]todos.TaskList, error) {
				return []todos.TaskList{
					{Id: 1, Title: "Monday", Tasks: []todos.Task{}},
					{Id: 2, Title: "Tuesday", Tasks: []todos.Task{}},
				}, nil
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.HandleFunc("/todos", mockHandler.GetTodos)
	r.ServeHTTP(rr, req)

	expectedStatusCode := 200
	if rr.Code != expectedStatusCode {
		t.Errorf("actual: %v\nexpected: %v", rr.Code, expectedStatusCode)
	}

	expectedBody := `{"status":"successfully fetched all todos","error":"","todos":[{"id":1,"title":"Monday","tasks":[]},{"id":2,"title":"Tuesday","tasks":[]}]}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("actual: %v\nexpected: %v", rr.Body.String(), expectedBody)
	}
}

func TestGetTodosFail(t *testing.T) {
	mockHandler := Handler{
		Service: mockTodoService{
			GetTodosFunc: func() ([]todos.TaskList, error) {
				return nil, errors.New("database error")
			},
		},
	}

	r := mux.NewRouter()

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.HandleFunc("/todos", mockHandler.GetTodos)
	r.ServeHTTP(rr, req)

	expectedStatusCode := 503
	if rr.Code != expectedStatusCode {
		t.Errorf("actual: %v\nexpected: %v", rr.Code, expectedStatusCode)
	}

	expectedBody := `{"status":"failed","error":"database error","todos":null}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("actual: %v\nexpected: %v", rr.Body.String(), expectedBody)
	}
}
