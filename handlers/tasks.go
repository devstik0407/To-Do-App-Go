package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/models"
	"todo/todos"

	"github.com/gorilla/mux"
)

func AddTask(rw http.ResponseWriter, r *http.Request) {
	resBody := struct {
		Status string      `json:"status"`
		Error  string      `json:"error"`
		Task   models.Task `json:"task"`
	}{
		Status: "",
		Error:  "",
		Task:   models.Task{},
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

	t := models.Task{}
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

	task, err := todos.AddTask(int64(listId), t.Desc)
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

func UpdateTask(rw http.ResponseWriter, r *http.Request) {
	resBody := struct {
		Status string      `json:"status"`
		Error  string      `json:"error"`
		Task   models.Task `json:"task"`
	}{
		Status: "",
		Error:  "",
		Task:   models.Task{},
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

	t := models.Task{}
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

	task, err := todos.UpdateTask(int64(listId), int64(taskId), t.Desc)
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

func DeleteTask(rw http.ResponseWriter, r *http.Request) {
	resBody := struct {
		Status string      `json:"status"`
		Error  string      `json:"error"`
		Task   models.Task `json:"task"`
	}{
		Status: "",
		Error:  "",
		Task:   models.Task{},
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

	task, err := todos.DeleteTask(int64(listId), int64(taskId))
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
