import React, { Component } from 'react';

class Book extends Component {
  render() {
    return (
      <div>
        <div>
          Book:  {this.props.books.name} ({this.props.books.Category})
        </div>
      </div>
    )
  }
}

export default Book;