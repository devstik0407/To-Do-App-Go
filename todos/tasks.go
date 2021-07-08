package todos

type Task struct {
	Id   int64  `json:"id"`
	Desc string `json:"desc"`
}

type Service struct {
	DataStore Store
}

type Store interface {
	AddTask(listId int64, newTask Task) error
	GetTask(listId, taskId int64) (Task, error)
	UpdateTask(listId int64, newTask Task) error
	DeleteTask(listId, taskId int64) (Task, error)
	CreateTaskList(newTaskList TaskList) error
	GetTaskList(listId int64) (TaskList, error)
	DeleteTaskList(listId int64) error
	GetTodos() ([]TaskList, error)
	MaxTaskIdInList(listId int64) (int64, error)
	MaxListId() (int64, error)
}

func (s Service) AddTask(listId int64, desc string) (Task, error) {
	newId, err := s.NewTaskId(listId)
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		Id:   newId,
		Desc: desc,
	}
	err = s.DataStore.AddTask(listId, newTask)
	if err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (s Service) DeleteTask(listId int64, taskId int64) (Task, error) {
	task, err := s.DataStore.DeleteTask(listId, taskId)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s Service) UpdateTask(listId, taskId int64, newDesc string) (Task, error) {
	newTask := Task{
		Id:   taskId,
		Desc: newDesc,
	}

	err := s.DataStore.UpdateTask(listId, newTask)
	if err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (s Service) GetTask(listId, taskId int64) (Task, error) {
	task, err := s.DataStore.GetTask(listId, taskId)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s Service) NewTaskId(listId int64) (int64, error) {
	maxId, err := s.DataStore.MaxTaskIdInList(listId)
	if err != nil {
		return 0, err
	}
	return maxId + 1, nil
}
