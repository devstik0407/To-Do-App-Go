package todos

import (
	"errors"
	"testing"
)

func TestAddTaskSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return 0, nil
			},
		},
	}

	actualTask, err := mockService.AddTask(1, "take medicine")
	if err != nil {
		t.Error(err)
	}

	expectedTask := Task{
		Id:   1,
		Desc: "take medicine",
	}
	if expectedTask != actualTask {
		t.Errorf("actual: %v \nexpected: %v", actualTask, expectedTask)
	}
}

func TestAddTaskFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return errors.New("invalid listId")
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return -1, errors.New("invalid listId")
			},
		},
	}

	_, err := mockService.AddTask(0, "take medicine")
	if err.Error() != "invalid listId" {
		t.Error(err)
	}
}

func TestDeleteTaskSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			DeleteTaskFunc: func(listId, taskId int64) (Task, error) {
				return Task{Id: taskId, Desc: "take medicine"}, nil
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return 0, nil
			},
		},
	}

	_, err := mockService.AddTask(1, "take medicine")
	if err != nil {
		t.Fatal(err)
	}

	actualTask, err := mockService.DeleteTask(1, 1)
	if err != nil {
		t.Error(err)
	}

	expectedTask := Task{
		Id:   1,
		Desc: "take medicine",
	}
	if expectedTask != actualTask {
		t.Errorf("actual: %v \nexpected: %v", actualTask, expectedTask)
	}
}

func TestDeleteTaskFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			DeleteTaskFunc: func(listId, taskId int64) (Task, error) {
				return Task{}, errors.New("invalid taskId")
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return 0, nil
			},
		},
	}

	_, err := mockService.AddTask(1, "take medicine")
	if err != nil {
		t.Fatal(err)
	}

	_, err = mockService.DeleteTask(1, 2)
	if err.Error() != "invalid taskId" {
		t.Error(err)
	}
}

func TestUpdateTaskSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			UpdateTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return 0, nil
			},
		},
	}
	_, err := mockService.AddTask(1, "eat rice")
	if err != nil {
		t.Fatal(err)
	}

	actualTask, err := mockService.UpdateTask(1, 1, "drink water")
	if err != nil {
		t.Error(err)
	}

	expectedTask := Task{
		Id:   1,
		Desc: "drink water",
	}
	if expectedTask != actualTask {
		t.Errorf("actual: %v \nexpected: %v", actualTask, expectedTask)
	}
}

func TestUpdateTaskFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			AddTaskFunc: func(listId int64, newTask Task) error {
				return nil
			},
			UpdateTaskFunc: func(listId int64, newTask Task) error {
				return errors.New("invalid listId")
			},
			MaxTaskIdInListFunc: func(listId int64) (int64, error) {
				return 0, nil
			},
		},
	}
	_, err := mockService.AddTask(1, "eat rice")
	if err != nil {
		t.Fatal(err)
	}

	_, err = mockService.UpdateTask(1, 0, "drink water")
	if err.Error() != "invalid listId" {
		t.Error(err)
	}
}

type mockStore struct {
	AddTaskFunc         func(listId int64, newTask Task) error
	GetTaskFunc         func(listId, taskId int64) (Task, error)
	UpdateTaskFunc      func(listId int64, newTask Task) error
	DeleteTaskFunc      func(listId, taskId int64) (Task, error)
	CreateTaskListFunc  func(newTaskList TaskList) error
	GetTaskListFunc     func(listId int64) (TaskList, error)
	DeleteTaskListFunc  func(listId int64) error
	GetTodosFunc        func() ([]TaskList, error)
	MaxTaskIdInListFunc func(listId int64) (int64, error)
	MaxListIdFunc       func() (int64, error)
}

func (ms mockStore) AddTask(listId int64, newTask Task) error {
	return ms.AddTaskFunc(listId, newTask)
}

func (ms mockStore) GetTask(listId, taskId int64) (Task, error) {
	return ms.GetTaskFunc(listId, taskId)
}

func (ms mockStore) UpdateTask(listId int64, newTask Task) error {
	return ms.UpdateTaskFunc(listId, newTask)
}

func (ms mockStore) DeleteTask(listId, taskId int64) (Task, error) {
	return ms.DeleteTaskFunc(listId, taskId)
}

func (ms mockStore) CreateTaskList(newTaskList TaskList) error {
	return ms.CreateTaskListFunc(newTaskList)
}

func (ms mockStore) GetTaskList(listId int64) (TaskList, error) {
	return ms.GetTaskListFunc(listId)
}

func (ms mockStore) DeleteTaskList(listId int64) error {
	return ms.DeleteTaskListFunc(listId)
}

func (ms mockStore) GetTodos() ([]TaskList, error) {
	return ms.GetTodosFunc()
}

func (ms mockStore) MaxTaskIdInList(listId int64) (int64, error) {
	return ms.MaxTaskIdInListFunc(listId)
}

func (ms mockStore) MaxListId() (int64, error) {
	return ms.MaxListIdFunc()
}
