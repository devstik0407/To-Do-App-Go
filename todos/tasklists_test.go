package todos

import (
	"errors"
	"testing"
)

func TestCreateTaskListSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			CreateTaskListFunc: func(newTaskList TaskList) error {
				return nil
			},
			MaxListIdFunc: func() (int64, error) {
				return 0, nil
			},
		},
	}
	actualListId, err := mockService.CreateTaskList("Monday")
	if err != nil {
		t.Error(err)
	}

	expectedListId := int64(1)
	if expectedListId != actualListId {
		t.Errorf("actual: %v \nexpected: %v", actualListId, expectedListId)
	}
}

func TestCreateTaskListFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			CreateTaskListFunc: func(newTaskList TaskList) error {
				return nil
			},
			MaxListIdFunc: func() (int64, error) {
				return 0, errors.New("unable to get max list id from data store")
			},
		},
	}
	_, err := mockService.CreateTaskList("Monday")
	if err == nil {
		t.Errorf("actual error: none \nexpected error: %v", errors.New("unable to get max list id from data store"))
	}
	if err.Error() != "unable to get max list id from data store" {
		t.Errorf("actual error: %v \nexpected error: %v", err, errors.New("unable to get max list id from data store"))
	}
}

func TestDeleteTaskListSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			CreateTaskListFunc: func(newTaskList TaskList) error {
				return nil
			},
			MaxListIdFunc: func() (int64, error) {
				return 0, nil
			},
			DeleteTaskListFunc: func(listId int64) error {
				return nil
			},
		},
	}
	_, err := mockService.CreateTaskList("Monday")
	if err != nil {
		t.Fatal(err)
	}

	actualListId, err := mockService.DeleteTaskList(1)
	if err != nil {
		t.Error(err)
	}
	expectedListId := 1
	if actualListId != int64(expectedListId) {
		t.Errorf("actual: %v \nexpected: %v", actualListId, expectedListId)
	}
}

func TestDeleteTaskListFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			CreateTaskListFunc: func(newTaskList TaskList) error {
				return nil
			},
			MaxListIdFunc: func() (int64, error) {
				return 0, nil
			},
			DeleteTaskListFunc: func(listId int64) error {
				return errors.New("invalid listId")
			},
		},
	}
	_, err := mockService.CreateTaskList("Monday")
	if err != nil {
		t.Fatal(err)
	}

	_, err = mockService.DeleteTaskList(2)
	if err == nil {
		t.Errorf("actual error: none \nexpected error: %v", errors.New("invalid listId"))
	}
	if err.Error() != "invalid listId" {
		t.Errorf("actual error: %v \nexpected error: %v", err, errors.New("invalid listId"))
	}
}

func TestGetTaskListSuccess(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			CreateTaskListFunc: func(newTaskList TaskList) error {
				return nil
			},
			MaxListIdFunc: func() (int64, error) {
				return 0, nil
			},
			GetTaskListFunc: func(listId int64) (TaskList, error) {
				return TaskList{Id: 1, Title: "Monday", Tasks: []Task{}}, nil
			},
		},
	}

	_, err := mockService.CreateTaskList("Monday")
	if err != nil {
		t.Fatal(err)
	}

	actualTaskList, err := mockService.GetTaskList(1)
	if err != nil {
		t.Error(err)
	}

	expectedTaskList := TaskList{
		Id:    1,
		Title: "Monday",
		Tasks: []Task{},
	}
	if (actualTaskList.Id != expectedTaskList.Id) ||
		(actualTaskList.Title != expectedTaskList.Title) ||
		(len(actualTaskList.Tasks) != 0) {
		t.Errorf("actual: %v \nexpected: %v", actualTaskList, expectedTaskList)
	}
}

func TestGetTaskListFail(t *testing.T) {
	mockService := Service{
		DataStore: mockStore{
			GetTaskListFunc: func(listId int64) (TaskList, error) {
				return TaskList{}, errors.New("invalid listId")
			},
		},
	}

	_, err := mockService.GetTaskList(1)
	if err == nil {
		t.Errorf("actual error: none \nexpected error: %v", errors.New("invalid listId"))
	}
	if err.Error() != "invalid listId" {
		t.Errorf("actual error: %v \nexpected error: %v", err, errors.New("invalid listId"))
	}
}
