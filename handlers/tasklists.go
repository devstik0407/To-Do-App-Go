package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/gorilla/mux"
)

func CreateTaskList(rw http.ResponseWriter, r *http.Request) {
	resBody := struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		ListId int64  `json:"listId"`
	}{
		Status: "",
		Error:  "",
		ListId: 0,
	}

	p := models.TaskList{}
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

	newListId, err := services.CreateTaskList(p.Title)
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

func DeleteTaskList(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	err = services.DeleteTaskList(int64(listId))
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	fmt.Fprint(rw, "successfully deleted task-list")
	// models.ShowTodos()
}

func GetTodos(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(services.Todos)
}
