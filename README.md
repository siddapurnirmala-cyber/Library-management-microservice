# LibFlow: Library Management Microservice

## About
LibFlow is a high-performance **Library Management Microservice** built with **Go (Golang)** and **GraphQL**. It is designed to efficiently manage library operations such as tracking books, registering members, and handling book borrowing/return cycles.

The project features secure **Google OAuth 2.0** authentication, stateless **JWT-based sessions**, and a robust **Role-Based Access Control (RBAC)** system. Secure data persistence is handled by **PostgreSQL** with advanced optimizations like connection pooling and transaction safety.

## Key Features

- **GraphQL API**: Flexible and efficient data querying and mutations for Books, Members, and Borrowing records.
- **Advanced Authentication**:
    - **Google OAuth 2.0 Integration**: Secure login flow with account selection.
    - **JWT (JSON Web Tokens)**: Stateless authentication with 24-hour expiration.
    - **Dual Auth Support**: Works via `HttpOnly` session cookies (for browsers) or `Authorization: Bearer` headers (for APIs).
- **Role-Based Access Control (RBAC)**:
    - Granular permissions enforced at the GraphQL resolver level.
    - Automatic assignment of the `MEMBER` role for new sign-ups.
- **Performance Optimized**: 
    - Database connection pooling.
    - Concurrent request handling in Go.

## Authentication Flow

LibFlow uses a secure OAuth 2.0 flow integrated with Google:

1.  **Initiation**: User hits `/auth/google/login`, which redirects to Google's OAuth consent screen with `prompt=select_account`.
2.  **State Verification**: A secure `oauthstate` cookie is used to prevent CSRF during the callback.
3.  **Token Exchange**: Upon successful login, the backend exchanges the authorization code for a Google access token.
4.  **User Provisioning**: The user's profile is upserted into the database. New users are assigned the `MEMBER` role by default.
5.  **Session Issues**: A JWT is generated containing the `user_id`, `email`, and `role`. This token is returned in the response body and set as an `HttpOnly` cookie.

## Role-Based Authorization

The application enforces permissions based on the user's role stored in the JWT claims:

- **ADMIN**: The highest level of access. Can perform all operations including managing books, members, and user roles.
- **LIBRARIAN**: Focused on operational tasks. Can view all records and manage the issuance (`borrowBook`) and return (`returnBook`) of books.
- **MEMBER**: Standard user level. Has read-only access to the book catalog.

## Prerequisites
- **Go**: Version 1.21 or higher.
- **PostgreSQL**: Running instance.
- **Google Cloud Console**: OAuth2 Credentials.

## Setup Instructions

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/siddapurnirmala-cyber/Library-management-microservice.git
    cd Library-management-microservice
    ```

2.  **Environment Configuration**:
    Create a `.env` file in the root directory:
    ```env
    DB_CONNECTION_STRING=host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=library sslmode=disable
    PORT=8082
    GOOGLE_CLIENT_ID=YOUR_GOOGLE_CLIENT_ID
    GOOGLE_CLIENT_SECRET=YOUR_GOOGLE_CLIENT_SECRET
    JWT_SECRET=YOUR_LONG_RANDOM_SECRET
    ```

3.  **Database Migration**:
    Run the migration script to set up the schema and roles:
    ```bash
    go run scripts/run_migration.go
    ```

4.  **Run the Server**:
    ```bash
    go run cmd/server/main.go
    ```

## API Explorer & Testing

Once the server is running, you can interact with the API:

- **Auth Login**: [http://localhost:8082/auth/google/login](http://localhost:8082/auth/google/login)
- **User Profile**: [http://localhost:8082/auth/me](http://localhost:8082/auth/me)
  - Returns the authenticated user's profile and role.
- **GraphiQL Playground**: [http://localhost:8082/graphql](http://localhost:8082/graphql)
  - Note: Authentication (JWT cookie or header) is required for most operations.

## Role-Permission Matrix

| Role      | Permissions                           |
| --------- | ------------------------------------- |
| Admin     | Add books, delete books, manage users |
| Librarian | Issue / return books                  |
| Member    | View books                            |
