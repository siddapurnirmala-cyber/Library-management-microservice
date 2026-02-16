const API_URL = 'http://localhost:8082/graphql';

const graphqlRequest = async (query, variables = {}) => {
    const response = await fetch(API_URL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query, variables }),
    });

    const result = await response.json();
    if (result.errors) {
        throw new Error(result.errors[0].message);
    }
    return result.data;
};

export const api = {
    getBooks: () => graphqlRequest(`
    query {
      books {
        id
        title
        author
        published_year
        total_copies
        available_copies
      }
    }
  `),

    getMembers: () => graphqlRequest(`
    query {
      members {
        id
        name
        email
        joined_at
      }
    }
  `),

    createBook: (book) => graphqlRequest(`
    mutation CreateBook($title: String!, $author: String!, $published_year: Int!, $total_copies: Int!) {
      createBook(title: $title, author: $author, published_year: $published_year, total_copies: $total_copies) {
        id
        title
      }
    }
  `, book),

    createMember: (member) => graphqlRequest(`
    mutation CreateMember($name: String!, $email: String!) {
      createMember(name: $name, email: $email) {
        id
        name
      }
    }
  `, member),

    borrowBook: (memberId, bookId) => graphqlRequest(`
    mutation BorrowBook($member_id: Int!, $book_id: Int!) {
      borrowBook(member_id: $member_id, book_id: $book_id) {
        id
        status
      }
    }
  `, { member_id: memberId, book_id: bookId }),

    returnBook: (borrowId) => graphqlRequest(`
    mutation ReturnBook($borrow_id: Int!) {
      returnBook(borrow_id: $borrow_id) {
        id
        status
      }
    }
  `, { borrow_id: borrowId })
};
