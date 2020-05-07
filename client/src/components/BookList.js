import React, { Component } from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import Book from './Book';

const FEED_QUERY = gql`
  query allBooks {
  books {
    id
    name
    Category
  }
}
  `;

class BookList extends Component {

  render() {
    return (
      <Query query={FEED_QUERY}>
        {({ loading, error, data }) => {
          if (loading) return <div>Fetching</div>
          if (error) return <div>Error</div>
    
          const booksToRender = data.books
    
          return (
            <div>
              {booksToRender.map(books => <Book key={books.id} books={books} />)}
            </div>
          )
        }}
      </Query>
    )
  }
}

export default BookList