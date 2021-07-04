package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/services"

	"github.com/gorilla/mux"
)

func TestAddTask(t *testing.T) {
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte(`{"title":"Monday"}`)))
	if err != nil {
		t.Fatal(err)
	}

	listId := services.NewTaskListId()
	r := mux.NewRouter()
	r.HandleFunc("/todos", CreateTaskList)
	rrC := httptest.NewRecorder()
	r.ServeHTTP(rrC, req)

	req, err = http.NewRequest("POST", fmt.Sprintf("/todos/%d", listId), bytes.NewBuffer([]byte(`{"desc":"code clean"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}", AddTask)
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

func TestUpdateTask(t *testing.T) {
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte(`{"title":"Tuesday"}`)))
	if err != nil {
		t.Fatal(err)
	}

	listId := services.NewTaskListId()
	r := mux.NewRouter()
	r.HandleFunc("/todos", CreateTaskList)
	rrC := httptest.NewRecorder()
	r.ServeHTTP(rrC, req)

	req, err = http.NewRequest("POST", fmt.Sprintf("/todos/%d", listId), bytes.NewBuffer([]byte(`{"desc":"drink water"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}", AddTask)
	rrA := httptest.NewRecorder()
	r.ServeHTTP(rrA, req)

	taskId := int64(1)
	req, err = http.NewRequest("UPDATE", fmt.Sprintf("/todos/%d/%d", listId, taskId), bytes.NewBuffer([]byte(`{"desc":"sleep"}`)))
	if err != nil {
		t.Fatal(err)
	}

	r.HandleFunc("/todos/{listId}/{taskId}", UpdateTask)
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
