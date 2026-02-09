package schema

import (
	"github.com/graphql-go/graphql"
)

// MemberType defines the GraphQL object for a Member
var MemberType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Member",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"name":      &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"joined_at": &graphql.Field{Type: graphql.String},
	},
})

// BookType defines the GraphQL object for a Book
var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
		"title":            &graphql.Field{Type: graphql.String},
		"author":           &graphql.Field{Type: graphql.String},
		"published_year":   &graphql.Field{Type: graphql.Int},
		"total_copies":     &graphql.Field{Type: graphql.Int},
		"available_copies": &graphql.Field{Type: graphql.Int},
	},
})

// BorrowType defines the GraphQL object for a Borrow record
var BorrowType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Borrow",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"member_id":   &graphql.Field{Type: graphql.Int},
		"book_id":     &graphql.Field{Type: graphql.Int},
		"borrow_date": &graphql.Field{Type: graphql.String},
		"return_date": &graphql.Field{Type: graphql.String},
		"status":      &graphql.Field{Type: graphql.String},
	},
})
