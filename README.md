# Todo App

This is a simple todo app built using Go. It uses the [Gorilla mux](https://www.gorillatoolkit.org/pkg/mux) router and [Gorm](https://gorm.io/) ORM.

## Endpoints

The app has the following endpoints:

### GET /todos

Returns a JSON array of all todos.

### POST /todos

Creates a new todo. The todo data should be sent in the request body as JSON.

### GET /todos/{id}

Returns the todo with the given ID.

### PUT /todos/{id}

Updates the todo with the given ID. The new todo data should be sent in the request body as JSON.

### DELETE /todos/{id}

Deletes the todo with the given ID.

## How the program works?

1. create a todoModel struct
2. create a todos slice of todoModel
3. create a getTodos function
4. create a createTodo function
5. create a getTodo function
6. create a updateTodo function
7. create a deleteTodo function
8. create a router
9. create a main function
10. create a db connection
11. migrate the db
12. run the server

## Get Started!

- set Environment variable:

``export DATABASE_URL= < url >``

**Database url example:**

``postgres://username:password@localhost:5432/todo?sslmode=disable``

## Start The Server:

-  ``go run todo.go``

or

- ``./bin/todo``

## Example:

- ``curl -X POST -H "Content-Type: application/json" -d '{"title":"Buy milk","completed":false,"due":"0001-01-01T00:00:00Z"}' http://localhost:8000/todos``

- ``curl -X GET -H "Content-Type: application/json" http://localhost:8000/todos``

- ``curl -X GET -H "Content-Type: application/json" http://localhost:8000/todos/1``

- ``curl -X PUT -H "Content-Type: application/json" -d '{"title":"Buy milk","completed":true,"due":"0001-01-01T00:00:00Z"}' http://localhost:8000/todos/1``

- ``curl -X DELETE -H "Content-Type: application/json" http://localhost:8000/todos/1``