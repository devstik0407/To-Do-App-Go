package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/todos"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service TodoService
}

type TodoService interface {
	AddTask(listId int64, desc string) (todos.Task, error)
	GetTask(listId, taskId int64) (todos.Task, error)
	UpdateTask(listId, taskId int64, newDesc string) (todos.Task, error)
	DeleteTask(listId, taskId int64) (todos.Task, error)
	CreateTaskList(title string) (int64, error)
	GetTaskList(listId int64) (todos.TaskList, error)
	DeleteTaskList(listId int64) (int64, error)
	GetTodos() ([]todos.TaskList, error)
	NewTaskListId() (int64, error)
	NewTaskId(listId int64) (int64, error)
}

type TaskResponse struct {
	Status string     `json:"status"`
	Error  string     `json:"error"`
	Task   todos.Task `json:"task"`
}

func (h Handler) AddTask(rw http.ResponseWriter, r *http.Request) {
	resBody := TaskResponse{
		Status: "",
		Error:  "",
		Task:   todos.Task{},
	}

	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	t := todos.Task{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err = d.Decode(&t)
	if err != nil {
		rw.WriteHeader(400)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	task, err := h.Service.AddTask(int64(listId), t.Desc)
	if err != nil {
		rw.WriteHeader(406)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	rw.WriteHeader(201)
	resBody.Status = "successfully added task"
	resBody.Task = task
	json.NewEncoder(rw).Encode(resBody)
}

func (h Handler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	resBody := TaskResponse{
		Status: "",
		Error:  "",
		Task:   todos.Task{},
	}

	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	taskId, err := strconv.Atoi(vars["taskId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	t := todos.Task{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err = d.Decode(&t)
	if err != nil {
		rw.WriteHeader(400)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	task, err := h.Service.UpdateTask(int64(listId), int64(taskId), t.Desc)
	if err != nil {
		rw.WriteHeader(406)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	rw.WriteHeader(201)
	resBody.Status = "successfully updated task"
	resBody.Task = task
	json.NewEncoder(rw).Encode(resBody)
}

func (h Handler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	resBody := TaskResponse{
		Status: "",
		Error:  "",
		Task:   todos.Task{},
	}

	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	taskId, err := strconv.Atoi(vars["taskId"])
	if err != nil {
		rw.WriteHeader(500)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	task, err := h.Service.DeleteTask(int64(listId), int64(taskId))
	if err != nil {
		rw.WriteHeader(406)
		resBody.Error = err.Error()
		resBody.Status = "failed"
		json.NewEncoder(rw).Encode(resBody)
		return
	}

	resBody.Status = "successfully deleted task"
	resBody.Task = task
	json.NewEncoder(rw).Encode(resBody)
}
