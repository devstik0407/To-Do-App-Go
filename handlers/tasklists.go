package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo/models"

	"github.com/gorilla/mux"
)

func CreateTaskList(rw http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var p models.TaskList
	json.Unmarshal(reqBody, &p)

	err := models.CreateTaskList(p.Title)
	if err != nil {
		fmt.Fprint(rw, err)
	}
	// models.ShowTodos()
}

func DeleteTaskList(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["listId"])
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	err = models.DeleteTaskList(int64(listId))
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	fmt.Fprint(rw, "successfully deleted task-list")
	// models.ShowTodos()
}

func GetTodos(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(models.Todos)
}
