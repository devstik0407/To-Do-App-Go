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
