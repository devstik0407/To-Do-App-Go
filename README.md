# To-Do-App-Go
A REST API for managing TODOs.

# Description
This a REST API for listing and maintaining your tasks. You can create a task-list, add, update and delete tasks from that list, delete a task-list and read all the task-lists.

# Requirements
- Go (version `1.16.5`)
- gorilla/mux package (`github.com/gorilla/mux`)

# How to use?
- Change working directory to the root directory of the application
- To run the server, run the following command 
```
go run main.go
```
## Create a task-list 
- Send a POST request to `/todos`
- Request body should contain the title of the task-list to create
```JSON
{
    "title": "Monday"
}
```
- Response body
```JSON
{
    "status": "successfully created task-list",
    "error": "",
    "listId": 1
}
```
## Add a task into a task-list
- Send a POST request to `/todos/{listId}` (`listId` is id of the task-list)
- Request body should contain the description of the task to add
```JSON
{
    "desc": "Read books"
}
```
- Response body
```JSON
{
    "status": "successfully added task",
    "error": "",
    "task": {
        "id": 1,
        "desc": "Read books"
    }
}
```
## Delete a task-list
- Send a DELETE request to `/todos/{listId}` (`listId` is id of the task-list)
- Response body
```JSON
{
    "status": "successfully deleted task-list",
    "error": "",
    "listId": 1
}
```
For more details, please check the `swagger.yaml` file

# Testing Guide
- To test the handlers, cd into the `handlers` directory and run
```
go test
```
- To test todos, cd into the `todos` directory and run
```
go test
```
- To test mongo-store
    <ol>
    <li>Set up a mongoDB server with an empty or no <code>todosDB</code></li>
    <li>cd into the <code>mongo-store</code> directory
    <li>Run the command <code>go test</code></li>
    </ol>
# Contributors
- Swastik Dutta - https://github.com/devstik0407