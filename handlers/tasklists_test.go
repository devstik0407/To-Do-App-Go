package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/services"
)

func TestCreateTaskList(t *testing.T) {
	testBody := []byte(`{"title":"Monday"}`)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}

	listId := services.NewTaskListId()

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
