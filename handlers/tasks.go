package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/gorilla/mux"
)

func AddTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var t models.Task
	json.Unmarshal(reqBody, &t)

	err = services.AddTask(int64(listId), t.Desc)
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	// models.ShowTodos()
	fmt.Fprint(rw, "successfully added task")
}

func UpdateTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	taskId, err := strconv.Atoi(vars["taskId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var t models.Task
	json.Unmarshal(reqBody, &t)

	err = services.UpdateTask(int64(listId), int64(taskId), t.Desc)
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	fmt.Fprint(rw, "successfully updated task in the list")
	// models.ShowTodos()
}

func DeleteTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	taskId, err := strconv.Atoi(vars["taskId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	err = services.DeleteTask(int64(listId), int64(taskId))
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	fmt.Fprint(rw, "successfully deleted task from list")
	// models.ShowTodos()
}
