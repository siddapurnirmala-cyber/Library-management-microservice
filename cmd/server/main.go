package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"library-system/pkg/db"
	"library-system/pkg/schema"

	_ "net/http/pprof" // Register pprof handlers

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connection string - in a real app, use environment variables
	// Default: host=localhost port=5432 user=postgres password=password dbname=library sslmode=disable
	connStr := "host=localhost port=5432 user=postgres password=password dbname=library sslmode=disable"
	if envConn := os.Getenv("DB_CONNECTION_STRING"); envConn != "" {
		connStr = envConn
	}

	db.InitDB(connStr)

	// Create GraphQL Schema handler
	h := handler.New(&handler.Config{
		Schema:   &schema.LibrarySchema,
		Pretty:   true,
		GraphiQL: true,
	})

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Handle("/graphql", h)

	// Enable pprof on Gorilla router by forwarding /debug/pprof/ to DefaultServeMux
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	// Port configuration
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	fmt.Printf("Server is running on port %s\n", port)
	fmt.Printf("GraphiQL playground enabled at http://localhost:%s/graphql\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
