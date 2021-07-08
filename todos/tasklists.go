package todos

var Todos []TaskList

type TaskList struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

func (s Service) CreateTaskList(title string) (int64, error) {
	newListId, err := s.NewTaskListId()
	if err != nil {
		return 0, err
	}

	taskList := TaskList{
		Id:    newListId,
		Title: title,
		Tasks: []Task{},
	}

	err = s.DataStore.CreateTaskList(taskList)
	if err != nil {
		return 0, err
	}
	return taskList.Id, nil
}

func (s Service) DeleteTaskList(listId int64) (int64, error) {
	err := s.DataStore.DeleteTaskList(listId)
	if err != nil {
		return 0, err
	}
	return listId, nil
}

func (s Service) GetTaskList(listId int64) (TaskList, error) {
	taskList, err := s.DataStore.GetTaskList(listId)
	if err != nil {
		return TaskList{}, err
	}
	return taskList, nil
}

func (s Service) NewTaskListId() (int64, error) {
	maxId, err := s.DataStore.MaxListId()
	if err != nil {
		return 0, err
	}
	return maxId + 1, nil
}
