# Golang Bookstore API

A simple RESTful API for managing a bookstore, built using the [Echo](https://echo.labstack.com/) web framework in Go.  
This version uses **in-memory storage**, ideal for learning, testing, and prototyping basic CRUD operations.

---

## Features

- Add a new book
- Get all books
- Get a book by ID
- Update a book by ID
- Delete a book by ID

---
## Tech Stack

- Language: Go (Golang)
- Framework: Echo (Web framework for building REST APIs)
- Storage: In-memory map (no database used in this version)
- Testing Tool: Postman for testing API endpoints

## Book Data Model

Each book has the following structure:

```go
type Book struct {
  ID      string  `json:"id"`
  Title   string  `json:"title"`
  Author  string  `json:"author"`
  Price   float64 `json:"price"`
  Content string  `json:"content"`
}
```
---
## How to run?

- Clone the Repository
``` bash
git clone https://github.com/vikranth-worldergroup/golang-bookstore-api.git
cd golang-bookstore-api
```

- Initialize Go Modules
``` bash
go mod tidy
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This installs dependencies like Echo automatically.
- Run all the files
``` bash
go run .
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The server will run on: http://localhost:8080

---
## API Endpoints
GET  /books → Fetch all books from the store.

GET  /books/:id → Retrieve a specific book using its ID.

POST  /books → Add a new book to the collection.

PUT  /books/:id → Update details of an existing book.

DELETE  /books/:id → Remove a book from the store using its ID



