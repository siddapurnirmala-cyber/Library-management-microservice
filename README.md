# Library Management Microservice

## About
This is a high-performance **Library Management Microservice** built with **Go (Golang)** and **GraphQL**. It is designed to efficiently manage library operations such as tracking books, registering members, and handling book borrowing/return cycles.

The service leverages **PostgreSQL** for data persistence and implements advanced optimizations like database connection pooling and rigorous performance profiling.

## Key Features
-   **GraphQL API**: Flexible and efficient data querying for Books, Members, and Borrowing records.
-   **Performance Optimized**: 
    -   Implements **Database Connection Pooling** for low-latency requests.
    -   Built-in **N+1 Query Handling** strategies.
    -   Integrated **pprof** for real-time CPU and Memory profiling.
-   **Transaction Management**: Ensures data integrity during book borrowing and returning.
-   **Scalable Architecture**: Service, Repository, and API layers are decoupled for easy maintenance.

## Tech Stack
-   **Language**: Go (Golang)
-   **API Standard**: GraphQL
-   **Database**: PostgreSQL
-   **Routing**: Gorilla Mux
-   **Driver**: lib/pq

## Getting Started
1.  **Clone the repository**:
    ```bash
    git clone https://github.com/siddapurnirmala-cyber/Library-management-microservice.git
    ```
2.  **Setup Database**: Use `scripts/init.sql` to initialize your PostgreSQL instance.
3.  **Run the Server**:
    ```bash
    go run cmd/server/main.go
    ```
4.  **Explore API**: Open `http://localhost:8080/graphql` for the GraphiQL playground.
