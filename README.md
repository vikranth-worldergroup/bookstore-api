# Bookstore API - Feature V2
A RESTful Bookstore API built with **Golang**, using the **Echo** web framework and **GORM** for database interactions. 
Includes features like MySQL integration, pagination, CSV export, and clean code architecture for production-ready API development.

---

## Features

- CRUD operations (Books & Authors)
- Echo web framework
- GORM ORM with MySQL
- Pagination
- CSV download of book data
- Validation and error handling

## Tech Stack

- Language: Go (Golang)
- Framework: Echo (Web framework for building REST APIs)
- Storage: MySQL Database 
- Database ORM: Gorm (Ref: https://gorm.io/index.html)
- Validator: Go-Validator (Ref: https://github.com/go-playground/validator)
- Testing Tool: Postman for testing API endpoints

## Book Data Model

Each Book has the following structure:

```go
type Author struct{
    BookID string `json:"-" gorm:"primaryKey"`
    Name string `json:"name" gorm:"primaryKey" validate:"required"`

}

type Book struct{
	ID string `json:"id" gorm:"primaryKey"`
    Title string `json:"title" validate:"required,max=50"`
    Price float64 `json:"price" vaidate:"required,gt=9,lt=501"`
    Content string `json:"content" validate:"required"`  
    Authors []Author `json:"authors" gorm:"foreignKey:BookID"`
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
---
## API Endpoints
GET  /books → Fetch all books from the store.

GET  /books/:id → Retrieve a specific book using its ID.

POST  /books → Add a new book to the collection.

PUT  /books/:id → Update details of an existing book.

DELETE  /books/:id → Remove a book from the store using its ID

GET  /books/download → Download all books data as a CSV file.
