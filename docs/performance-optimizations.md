# Performance Optimizations

## Database Layer

The database layer affects the entire system's latency. Optimizations here focus on connection management and query efficiency.

### Connection Pooling
- **Problem**: Opening a new physical connection to the database for every request is expensive (TCP handshake, authentication).
- **Solution**: Maintain a pool of idle connections that can be reused.
- **Configuration**: Implemented in `pkg/db/postgres.go`.
  ```go
  DB.SetMaxOpenConns(25)                 // Limits total open connections
  DB.SetMaxIdleConns(25)                 // Limits idle connections waiting for reuse
  DB.SetConnMaxLifetime(5 * time.Minute) // recycles connections to prevent stale issues
  ```
- **Why it helps**: Reduces latency by ~10-100ms per request (depending on network) by skipping connection setup.

### Query Optimization & Indexing
- **Problem**: Full table scans are slow as data grows.
- **Solution**: Use indexes on columns frequently used in `WHERE`, `JOIN`, or `ORDER BY` clauses.
- **Current State**: 
  - `PRIMARY KEY`s on `id` columns are automatically indexed.
  - `email` in `members` table is `UNIQUE`, creating an implicit index.
- **Recommendation**: As noted in `optimization_guide.md`, add indexes on foreign keys if filtering heavily (e.g., `CREATE INDEX idx_borrow_book_id ON borrow(book_id)`).

### Prepared Statements
- **Problem**: Parsing SQL queries for every execution consumes DB CPU.
- **Solution**: Use parameterized queries (e.g., `Values ($1, $2)`).
- **Current State**: The application uses `database/sql`'s parameterized methods (`QueryRow("... $1 ...", args...)`).
- **Why it helps**: Postgres parses the query plan once and reuses it. Also prevents SQL Injection.

### Reducing N+1 Queries
- **Problem**: Fetching a list of N items and then performing a separate query for each item's related data results in N+1 total queries.
- **Solution**: Use **DataLoaders** to batch requests (e.g., "SELECT * FROM books WHERE id IN (...)").
- **Current State**: 
  - The current GraphQL schema (`pkg/schema/types.go`) uses **flat types** (returning `book_id` instead of a nested `Book` object).
  - This design **avoids** N+1 issues by default but shifts the burden of fetching related data to the client.
  - **Recommendation**: If nesting is added (e.g., `Borrow.book`), implement DataLoaders.

---

## Service Layer (Go Backend)

Optimizations in the application logic ensure efficient use of CPU and Memory.

### CPU Profiling with pprof
- **Problem**: Difficulty identifying code hotspots or memory leaks.
- **Solution**: Enable Go's built-in profiler.
- **Configuration**: Enabled in `cmd/server/main.go`:
  ```go
  import _ "net/http/pprof"
  // ...
  r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
  ```
- **Usage**:
  - `go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30` to analyze CPU usage.
  - `go tool pprof http://localhost:8080/debug/pprof/heap` to analyze memory allocations.

### Goroutine Usage
- **Problem**: Serial processing of requests limits throughput.
- **Solution**: `net/http` automatically spawns a lightweight **goroutine** for every incoming request.
- **Result**: The server can handle thousands of concurrent requests with low memory overhead (~2KB per goroutine).

### Context Timeouts (Recommendation)
- **Problem**: API requests hanging indefinitely if the DB is slow/locked.
- **Current State**: Not currently implemented in `pkg/schema` (uses background context).
- **Recommendation**: Update resolvers to use `ctx, cancel := context.WithTimeout(p.Context, 2*time.Second)` and pass `ctx` to `db.QueryContext`.

---

## API / Network Layer

Optimizations here minimize data transfer and improve client perceived performance.

### GraphQL Field Selection
- **Problem**: Over-fetching data (getting fields the client doesn't need).
- **Solution**: GraphQL clients specify exactly what fields they need.
- **Example**:
  ```graphql
  query {
    members {
      name  # Only fetches name, ignoring id/email/joined_at
    }
  }
  ```
- **Result**: Reduces payload size, saving bandwidth.

### Pagination (Recommendation)
- **Problem**: The `members` and `books` queries currently perform `SELECT *`. Large datasets will crash the server or slow down the response.
- **Recommendation**: Implement `limit` and `offset` arguments in `RootQuery` fields.
    ```sql
    SELECT * FROM books LIMIT $1 OFFSET $2
    ```

### HTTP Compression
- **Problem**: Large JSON responses consume bandwidth.
- **Recommendation**: Add a middleware (like `handlers.CompressHandler`) to enable Gzip compression for responses.
