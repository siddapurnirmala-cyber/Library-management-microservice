package schema

import (
	"database/sql"
	"errors"
	"library-system/pkg/db"
	"library-system/pkg/models"

	"github.com/graphql-go/graphql"
)

// RootQuery defines the read operations
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"members": &graphql.Field{
			Type: graphql.NewList(MemberType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rows, err := db.DB.Query("SELECT id, name, email, joined_at FROM members")
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				var members []models.Member
				for rows.Next() {
					var m models.Member
					if err := rows.Scan(&m.ID, &m.Name, &m.Email, &m.JoinedAt); err != nil {
						return nil, err
					}
					members = append(members, m)
				}
				return members, nil
			},
		},
		"books": &graphql.Field{
			Type: graphql.NewList(BookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rows, err := db.DB.Query("SELECT id, title, author, published_year, total_copies, available_copies FROM books")
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				var books []models.Book
				for rows.Next() {
					var b models.Book
					if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear, &b.TotalCopies, &b.AvailableCopies); err != nil {
						return nil, err
					}
					books = append(books, b)
				}
				return books, nil
			},
		},
		"borrows": &graphql.Field{
			Type: graphql.NewList(BorrowType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rows, err := db.DB.Query("SELECT id, member_id, book_id, borrow_date, return_date, status FROM borrow")
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				var borrows []models.Borrow
				for rows.Next() {
					var b models.Borrow
					var returnDate sql.NullTime
					if err := rows.Scan(&b.ID, &b.MemberID, &b.BookID, &b.BorrowDate, &returnDate, &b.Status); err != nil {
						return nil, err
					}
					if returnDate.Valid {
						b.ReturnDate = &returnDate.Time
					}
					borrows = append(borrows, b)
				}
				return borrows, nil
			},
		},
	},
})

// RootMutation defines the write operations
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createMember": &graphql.Field{
			Type: MemberType,
			Args: graphql.FieldConfigArgument{
				"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				email := p.Args["email"].(string)
				var m models.Member
				err := db.DB.QueryRow("INSERT INTO members (name, email) VALUES ($1, $2) RETURNING id, name, email, joined_at", name, email).Scan(&m.ID, &m.Name, &m.Email, &m.JoinedAt)
				if err != nil {
					return nil, err
				}
				return m, nil
			},
		},
		"updateMember": &graphql.Field{
			Type: MemberType,
			Args: graphql.FieldConfigArgument{
				"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				name := p.Args["name"].(string)
				email := p.Args["email"].(string)
				var m models.Member
				err := db.DB.QueryRow("UPDATE members SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email, joined_at", name, email, id).Scan(&m.ID, &m.Name, &m.Email, &m.JoinedAt)
				if err != nil {
					return nil, err
				}
				return m, nil
			},
		},
		"deleteMember": &graphql.Field{
			Type: MemberType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				var m models.Member
				err := db.DB.QueryRow("DELETE FROM members WHERE id = $1 RETURNING id, name, email, joined_at", id).Scan(&m.ID, &m.Name, &m.Email, &m.JoinedAt)
				if err != nil {
					return nil, err
				}
				return m, nil
			},
		},
		"createBook": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"title":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"author":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"published_year": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"total_copies":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title := p.Args["title"].(string)
				author := p.Args["author"].(string)
				pubYear := p.Args["published_year"].(int)
				totalCopies := p.Args["total_copies"].(int)

				var b models.Book
				err := db.DB.QueryRow("INSERT INTO books (title, author, published_year, total_copies, available_copies) VALUES ($1, $2, $3, $4, $4) RETURNING id, title, author, published_year, total_copies, available_copies", title, author, pubYear, totalCopies).Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear, &b.TotalCopies, &b.AvailableCopies)
				if err != nil {
					return nil, err
				}
				return b, nil
			},
		},
		"updateBook": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"id":             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"title":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"author":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"published_year": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"total_copies":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				title := p.Args["title"].(string)
				author := p.Args["author"].(string)
				pubYear := p.Args["published_year"].(int)
				totalCopies := p.Args["total_copies"].(int)

				// Calculate difference in copies to update available copies
				var currentTotal int
				err := db.DB.QueryRow("SELECT total_copies FROM books WHERE id = $1", id).Scan(&currentTotal)
				if err != nil {
					return nil, err
				}
				diff := totalCopies - currentTotal

				var b models.Book
				err = db.DB.QueryRow("UPDATE books SET title = $1, author = $2, published_year = $3, total_copies = $4, available_copies = available_copies + $5 WHERE id = $6 RETURNING id, title, author, published_year, total_copies, available_copies", title, author, pubYear, totalCopies, diff, id).Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear, &b.TotalCopies, &b.AvailableCopies)
				if err != nil {
					return nil, err
				}
				return b, nil
			},
		},
		"deleteBook": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				var b models.Book
				err := db.DB.QueryRow("DELETE FROM books WHERE id = $1 RETURNING id, title, author, published_year, total_copies, available_copies", id).Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear, &b.TotalCopies, &b.AvailableCopies)
				if err != nil {
					return nil, err
				}
				return b, nil
			},
		},
		"borrowBook": &graphql.Field{
			Type: BorrowType,
			Args: graphql.FieldConfigArgument{
				"member_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"book_id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				memberID := p.Args["member_id"].(int)
				bookID := p.Args["book_id"].(int)

				tx, err := db.DB.Begin()
				if err != nil {
					return nil, err
				}
				defer tx.Rollback()

				// Check availability
				var available int
				err = tx.QueryRow("SELECT available_copies FROM books WHERE id = $1 FOR UPDATE", bookID).Scan(&available)
				if err != nil {
					return nil, err
				}
				if available <= 0 {
					return nil, errors.New("book not available")
				}

				// Update book availability
				_, err = tx.Exec("UPDATE books SET available_copies = available_copies - 1 WHERE id = $1", bookID)
				if err != nil {
					return nil, err
				}

				// Create borrow record
				var b models.Borrow
				err = tx.QueryRow("INSERT INTO borrow (member_id, book_id, status) VALUES ($1, $2, 'borrowed') RETURNING id, member_id, book_id, borrow_date, status", memberID, bookID).Scan(&b.ID, &b.MemberID, &b.BookID, &b.BorrowDate, &b.Status)
				if err != nil {
					return nil, err
				}

				if err := tx.Commit(); err != nil {
					return nil, err
				}
				return b, nil
			},
		},
		"returnBook": &graphql.Field{
			Type: BorrowType,
			Args: graphql.FieldConfigArgument{
				"borrow_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				borrowID := p.Args["borrow_id"].(int)

				tx, err := db.DB.Begin()
				if err != nil {
					return nil, err
				}
				defer tx.Rollback()

				// Get Borrow Record
				var bookID int
				var status string
				err = tx.QueryRow("SELECT book_id, status FROM borrow WHERE id = $1 FOR UPDATE", borrowID).Scan(&bookID, &status)
				if err != nil {
					return nil, err
				}

				if status == "returned" {
					return nil, errors.New("book already returned")
				}

				// Update Borrow Record
				var b models.Borrow
				var returnDate sql.NullTime
				err = tx.QueryRow("UPDATE borrow SET status = 'returned', return_date = CURRENT_TIMESTAMP WHERE id = $1 RETURNING id, member_id, book_id, borrow_date, return_date, status", borrowID).Scan(&b.ID, &b.MemberID, &b.BookID, &b.BorrowDate, &returnDate, &b.Status)
				if err != nil {
					return nil, err
				}
				if returnDate.Valid {
					b.ReturnDate = &returnDate.Time
				}

				// Update Book Availability
				_, err = tx.Exec("UPDATE books SET available_copies = available_copies + 1 WHERE id = $1", bookID)
				if err != nil {
					return nil, err
				}

				if err := tx.Commit(); err != nil {
					return nil, err
				}
				return b, nil
			},
		},
	},
})

var LibrarySchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
