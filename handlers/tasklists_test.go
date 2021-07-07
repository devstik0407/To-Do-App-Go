package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/todos"

	"github.com/gorilla/mux"
)

func TestCreateTaskList(t *testing.T) {
	testBody := []byte(`{"title":"Monday"}`)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}

	listId := todos.NewTaskListId()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskList)
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

func TestDeleteTaskList(t *testing.T) {
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte(`{"title":"Tuesday"}`)))
	if err != nil {
		t.Fatal(err)
	}

	listId := todos.NewTaskListId()
	rrC := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/todos", CreateTaskList)
	r.ServeHTTP(rrC, req)

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/todos/%d", listId), nil)
	if err != nil {
		t.Fatal(err)
	}

	rrD := httptest.NewRecorder()
	r.HandleFunc("/todos/{listId}", DeleteTaskList)
	r.ServeHTTP(rrD, req)

	expectedStatusCode := 200
	if status := rrD.Code; status != expectedStatusCode {
		t.Errorf("actual: %v expected: %v", status, expectedStatusCode)
	}

	expectedBody := fmt.Sprintf(`{"status":"successfully deleted task-list","error":"","listId":%d}`, listId) + "\n"
	if rrD.Body.String() != expectedBody {
		t.Errorf("actual: %v expected: %v", rrD.Body.String(), expectedBody)
	}
}
