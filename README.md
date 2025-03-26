# Book Management API

A RESTful API built with Go and Gin framework for managing a library's book inventory. This API provides endpoints for CRUD operations on books and handles book checkout/return functionality.

## Features

- Get all books
- Get a specific book by ID
- Create a new book
- Update an existing book
- Delete a book
- Checkout a book
- Return a book
- Automatic inventory management
- Duplicate book prevention

## Prerequisites

- Go 1.16 or higher
- Git

## Project Structure

```
go-api/
├── main.go
└── README.md
```

## Getting Started

1. Clone the repository:
```bash
git clone <your-repository-url>
cd go-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Running with Docker

The application is available as a Docker image on Docker Hub. You can run it using:

```bash
docker run -it -p 8080:8080 shivang21007/book-api:v1.0
```

This command will:
- Pull the image from Docker Hub
- Run it in interactive mode (-it)
- Map port 8080 from the container to port 8080 on your host machine (-p 8080:8080)

The API will be accessible at `http://localhost:8080`

## API Endpoints

### 1. Get All Books
```bash
GET http://localhost:8080/books
```

### 2. Get Book by ID
```bash
GET http://localhost:8080/books/:id
```

### 3. Create New Book
```bash
POST http://localhost:8080/books
```

Sample request body:
```json
{
    "id": "5",
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "quantity": 3
}
```

### 4. Update Book
```bash
PUT http://localhost:8080/books/:id
```

Sample request body:
```json
{
    "id": "5",
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "quantity": 4
}
```

### 5. Delete Book
```bash
DELETE http://localhost:8080/books/:id
```

### 6. Checkout Book
```bash
PATCH http://localhost:8080/checkout?id=:id
```

### 7. Return Book
```bash
PATCH http://localhost:8080/return?id=:id
```

## Testing with cURL

Here are some example cURL commands to test the API:

1. Get all books:
```bash
curl http://localhost:8080/books
```

2. Get a specific book:
```bash
curl http://localhost:8080/books/1
```

3. Create a new book:
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "5",
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "quantity": 3
  }'
```

4. Update a book:
```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "Updated Title",
    "author": "Updated Author",
    "quantity": 5
  }'
```

5. Delete a book:
```bash
curl -X DELETE http://localhost:8080/books/1
```

6. Checkout a book:
```bash
curl -X PATCH "http://localhost:8080/checkout?id=1"
```

7. Return a book:
```bash
curl -X PATCH "http://localhost:8080/return?id=1"
```

## Error Handling

The API includes proper error handling for various scenarios:
- Invalid request body (400 Bad Request)
- Book not found (404 Not Found)
- Duplicate book ID (409 Conflict)
- Duplicate book title (409 Conflict)
- Book not available for checkout (400 Bad Request)
- Cannot return more books than original quantity (400 Bad Request)

## Contributing

Feel free to submit issues and enhancement requests!