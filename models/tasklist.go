package models

import (
	"fmt"
)

type TaskList struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

func (tl TaskList) show() {
	fmt.Printf("id: %d, title: %s, tasks: %v\n", tl.Id, tl.Title, tl.Tasks)
}
