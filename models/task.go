package models

import (
	"fmt"
)

type Task struct {
	Id   int64  `json:"id"`
	Desc string `json:"desc"`
}

func (t Task) show() {
	fmt.Printf("id: %d, description: %s", t.Id, t.Desc)
}
