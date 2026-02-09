# Optimization & Profiling Guide

This guide covers how to optimize your Go backend to achieve <200ms response times and how to analyze performance using `pprof`.

## 1. Database Optimization

Achieving fast response times often starts with the database.

### Connection Pooling
Opening a new database connection for every request is expensive (handshake, auth). Instead, use a connection pool.
We will configure `pkg/db/postgres.go` to keep connections open:
```go
DB.SetMaxOpenConns(25)
DB.SetMaxIdleConns(25)
DB.SetConnMaxLifetime(5 * time.Minute)
```

### Indexing
Ensure your SQL tables have indexes on columns used in `WHERE` clauses and Foreign Keys.
In `init.sql`, we already have Primary Keys (`id`), which are automatically indexed.
If you search by `book_id` in `borrow` table often, add an index:
```sql
CREATE INDEX idx_borrow_book_id ON borrow(book_id);
```

### GraphQL N+1 Problem
In GraphQL, fetching a list of items (e.g., Borrows) and then fetching a related item for *each* one (e.g., Book details) can cause N+1 queries.
*   **Solution**: Use "DataLoaders" to batch requests (e.g., fetch all Book IDs in one query).
*   **Current State**: Your current implementation performs single SQL queries per resolver. Since our schema is simple, this is fine for now, but be aware of it as you scale.

## 2. Profiling with pprof

`pprof` is a tool for visualization and analysis of profiling data.

### Setup
Import `net/http/pprof` in your `main.go`. This automatically registers debug handlers at `/debug/pprof/`.

### How to Profile
1.  **Start your server**: `go run cmd/server/main.go`
2.  **Generate Load**: Run your Postman collection or use a tool like `wrk` or `hey` to send many requests.
3.  **Capture Profile**:
    While the server is under load, run:
    ```bash
    go tool pprof http://localhost:8081/debug/pprof/profile?seconds=30
    ```
4.  **Analyze**:
    Inside the `pprof` interactive shell, type:
    -   `top`: Shows functions taking the most CPU time.
    -   `list <FunctionName>`: Shows line-by-line breakdown of a function.
    -   `web`: Opens a visualization in your browser (requires Graphviz).

## 3. General Go Tips
-   **JSON Serialization**: `encoding/json` uses reflection and can be slow. For extreme performance, consider `easyjson` or `json-iterator/go`.
-   **Structure**: Keep handlers thin. Move logic to the service layer.
