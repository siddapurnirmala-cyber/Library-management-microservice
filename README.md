# LibFlow: Library Management Microservice

## About
LibFlow is a high-performance **Library Management Microservice** built with **Go (Golang)** and **GraphQL**. It is designed to efficiently manage library operations such as tracking books, registering members, and handling book borrowing/return cycles.

This project provides a comprehensive solution for modern libraries, combining a robust Go backend with a sleek, reactive frontend. It features secure Google OAuth authentication, real-time data synchronization via GraphQL, and a performant PostgreSQL database. Whether you're tracking thousands of titles or managing hundreds of members, LibFlow's architecture ensures low latency and high reliability through advanced optimizations like connection pooling and transaction safety.

The project now includes a **Premium React Frontend** for a seamless user experience.

## Key Features
- **GraphQL API**: Flexible and efficient data querying for Books, Members, and Borrowing records.
- **Secure Authentication**: 
    - **Google OAuth 2.0 Integration**: Easy and secure login.
    - **JWT (JSON Web Tokens)**: Stateless authentication for API security.
    - **Session Management**: 10-minute secure session tokens.
- **Modern Frontend**: Built with **React + Vite**, featuring glassmorphism and real-time dashboard stats.
- **Performance Optimized**: 
    - Implements **Database Connection Pooling** for low-latency requests.
    - Integrated **pprof** for real-time profiling.
- **Transaction Management**: Ensures data integrity during book operations.

## Tech Stack
- **Backend**: Go (Golang), Gorilla Mux, GraphQL-go
- **Frontend**: React, Vite, Lucide-React (Icons)
- **Database**: PostgreSQL
- **Security**: OAuth2, JWT

## Authentication Overview
All GraphQL APIs are protected by an `AuthMiddleware`. To access the API:
1. Users must log in via the Google OAuth flow at `/auth/google/login`.
2. Upon success, a secure `session_token` (JWT) is set in a Cookie.
3. The frontend automatically handles these tokens to authenticate GraphQL requests.

## Getting Started

### 1. Backend Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/siddapurnirmala-cyber/Library-management-microservice.git
   ```
2. **Configure Environment**: Create a `.env` file in the root directory:
   ```env
   DB_CONNECTION_STRING=your_postgres_url
   PORT=8082
   GOOGLE_CLIENT_ID=your_id
   GOOGLE_CLIENT_SECRET=your_secret
   JWT_SECRET=your_secret
   ```
3. **Run the Server**:
   ```bash
   go run cmd/server/main.go
   ```

### 2. Frontend Setup
1. **Navigate to the frontend folder**:
   ```bash
   cd frontend
   ```
2. **Install dependencies**:
   ```bash
   npm install
   ```
3. **Run the Dev Server**:
   ```bash
   npm run dev
   ```
   Open the URL provided in the terminal (usually `http://localhost:5173` or `5174`).

## API Explorer
- **GraphiQL Playground**: [http://localhost:8082/graphql](http://localhost:8082/graphql) (Requires authentication)
- **Auth Login**: [http://localhost:8082/auth/google/login](http://localhost:8082/auth/google/login)
