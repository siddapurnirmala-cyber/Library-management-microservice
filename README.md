# Library Management System

A simple backend for a Library Management System using Golang, Gorilla Mux, PostgreSQL, and GraphQL.

## Prerequisites

- [Go](https://golang.org/) (1.18+)
- [PostgreSQL](https://www.postgresql.org/)

## Setup

1.  **Clone the repository** (if applicable) or navigate to the project directory.

2.  **Database Setup**
    - Create a PostgreSQL database named `library`.
    - Run the initialization script to create tables:
      ```bash
      psql -U postgres -d library -f scripts/init.sql
      ```

3.  **Install Dependencies**
    ```bash
    go mod tidy
    ```

4.  **Configuration**
    - The application defaults to: `host=localhost port=5432 user=postgres password=password dbname=library sslmode=disable`
    - You can override this by setting the `DB_CONNECTION_STRING` environment variable.

5.  **Run the Server**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start at `http://localhost:8080`.

## API Usage (GraphQL)

You can use the built-in GraphiQL playground at `http://localhost:8080/graphql` or use Postman.

### Example Queries & Mutations

#### 1. Create a Member
```graphql
mutation {
  createMember(name: "John Doe", email: "john@example.com") {
    id
    name
    email
  }
}
```

#### 2. Create a Book
```graphql
mutation {
  createBook(title: "The Go Programming Language", author: "Donovan & Kernighan", published_year: 2015, total_copies: 5) {
    id
    title
    available_copies
  }
}
```

#### 3. Borrow a Book
```graphql
mutation {
  borrowBook(member_id: 1, book_id: 1) {
    id
    status
    borrow_date
  }
}
```

#### 4. Return a Book
```graphql
mutation {
  returnBook(borrow_id: 1) {
    id
    status
    return_date
  }
}
```

#### 5. Get All Books
```graphql
query {
  books {
    id
    title
    available_copies
  }
}
```
