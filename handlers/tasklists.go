package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/todos"

	"github.com/gorilla/mux"
)

type TaskListResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ListId int64  `json:"listId"`
}

func (h Handler) CreateTaskList(rw http.ResponseWriter, r *http.Request) {
	resBody := TaskListResponse{
		Status: "",
		Error:  "",
		ListId: 0,
	}

	p := todos.TaskList{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&p)
	if err != nil {
		rw.WriteHeader(400)
		resBody.Status = "failed"
		resBody.Error = err.Error()
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	newListId, err := h.Service.CreateTaskList(p.Title)
	if err != nil {
		rw.WriteHeader(500)
		resBody.Status = "failed"
		resBody.Error = err.Error()
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	rw.WriteHeader(201)
	resBody.ListId = newListId
	resBody.Status = "successfully created task-list"
	json.NewEncoder(rw).Encode(resBody)
}

func (h Handler) DeleteTaskList(rw http.ResponseWriter, r *http.Request) {
	resBody := TaskListResponse{
		Status: "",
		Error:  "",
		ListId: 0,
	}

	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Status = "failed"
		resBody.Error = err.Error()
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	_, err = h.Service.DeleteTaskList(int64(listId))
	if err != nil {
		rw.WriteHeader(406)
		resBody.Status = "failed"
		resBody.Error = err.Error()
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	rw.WriteHeader(200)
	resBody.Status = "successfully deleted task-list"
	resBody.ListId = int64(listId)
	json.NewEncoder(rw).Encode(resBody)
}

func (h Handler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	resBody := struct {
		Status string           `json:"status"`
		Error  string           `json:"error"`
		Todos  []todos.TaskList `json:"todos"`
	}{}
	defer func() {
		json.NewEncoder(rw).Encode(resBody)
	}()

	todos, err := h.Service.GetTodos()
	if err != nil {
		rw.WriteHeader(503)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		return
	}

	resBody.Status = "successfully fetched all todos"
	resBody.Todos = todos
}
