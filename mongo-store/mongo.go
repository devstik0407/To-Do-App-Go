package mongostore

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo/todos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func Connect(ctx context.Context) MongoDB {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("error while connecting to MongoDB: %v", err))
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return MongoDB{Client: client}
}

func (md MongoDB) CreateTaskList(newTaskList todos.TaskList) error {
	collection := md.Client.Database("todosDB").Collection("todos")
	bsonTaskList, err := bson.Marshal(newTaskList)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, bsonTaskList)
	if err != nil {
		return err
	}
	return nil
}

func (md MongoDB) GetTaskList(listId int64) (todos.TaskList, error) {
	collection := md.Client.Database("todosDB").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res todos.TaskList
	err := collection.FindOne(ctx, bson.D{{"id", listId}}).Decode(&res)
	if err != nil {
		return todos.TaskList{}, errors.New("invalid listId")
	}
	return res, nil
}

func (md MongoDB) DeleteTaskList(listId int64) error {
	collection := md.Client.Database("todosDB").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.D{{"id", listId}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("invalid listId")
	}
	return nil
}

func (md MongoDB) GetTodos() ([]todos.TaskList, error) {
	collection := md.Client.Database("todosDB").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res []todos.TaskList

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (md MongoDB) MaxTaskIdInList(listId int64) (int64, error) {
	taskList, err := md.GetTaskList(listId)
	if err != nil {
		return -1, err
	}

	maxTaskId := int64(0)
	for i := 0; i < len(taskList.Tasks); i++ {
		if maxTaskId < taskList.Tasks[i].Id {
			maxTaskId = taskList.Tasks[i].Id
		}
	}
	return maxTaskId, nil
}

func (md MongoDB) MaxListId() (int64, error) {
	todos, err := md.GetTodos()
	if err != nil {
		return -1, err
	}

	maxListId := int64(0)
	for i := 0; i < len(todos); i++ {
		if maxListId < todos[i].Id {
			maxListId = todos[i].Id
		}
	}
	return maxListId, nil
}

func (md MongoDB) AddTask(listId int64, newTask todos.Task) error {
	taskList, err := md.GetTaskList(listId)
	if err != nil {
		return err
	}

	taskList.Tasks = append(taskList.Tasks, newTask)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := md.Client.Database("todosDB").Collection("todos")
	err = collection.FindOneAndUpdate(ctx, bson.D{{"id", listId}}, bson.D{{"$set", bson.D{{"tasks", taskList.Tasks}}}}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (md MongoDB) GetTask(listId, taskId int64) (todos.Task, error) {
	taskList, err := md.GetTaskList(listId)
	if err != nil {
		return todos.Task{}, err
	}

	for i := 0; i < len(taskList.Tasks); i++ {
		if taskList.Tasks[i].Id == taskId {
			return taskList.Tasks[i], nil
		}
	}
	return todos.Task{}, errors.New("invalid taskId")
}

func (md MongoDB) UpdateTask(listId int64, newTask todos.Task) error {
	taskList, err := md.GetTaskList(listId)
	if err != nil {
		return err
	}

	taskFound := false
	for i := 0; i < len(taskList.Tasks); i++ {
		if taskList.Tasks[i].Id == newTask.Id {
			taskFound = true
			taskList.Tasks[i] = newTask
		}
	}
	if !taskFound {
		return errors.New("invalid taskId")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := md.Client.Database("todosDB").Collection("todos")
	err = collection.FindOneAndUpdate(ctx, bson.D{{"id", listId}}, bson.D{{"$set", bson.D{{"tasks", taskList.Tasks}}}}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (md MongoDB) DeleteTask(listId, taskId int64) (todos.Task, error) {
	taskList, err := md.GetTaskList(listId)
	if err != nil {
		return todos.Task{}, err
	}

	deletedTask := todos.Task{}
	taskFound := false
	index := -1
	for i := 0; i < len(taskList.Tasks); i++ {
		if taskList.Tasks[i].Id == taskId {
			deletedTask = taskList.Tasks[i]
			index = i
			taskFound = true
		}
	}
	if !taskFound {
		return todos.Task{}, errors.New("invalid taskId")
	}

	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := md.Client.Database("todosDB").Collection("todos")
	err = collection.FindOneAndUpdate(ctx, bson.D{{"id", listId}}, bson.D{{"$set", bson.D{{"tasks", taskList.Tasks}}}}).Err()
	if err != nil {
		return todos.Task{}, err
	}
	return deletedTask, nil
}
