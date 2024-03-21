DROP TABLE IF EXISTS books;

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       book_name VARCHAR(255),
                       author VARCHAR(255),
                       genre VARCHAR(100),
                       year INT,
                       count INT,
                       price NUMERIC(10, 2)
);

-- Creating indexes
CREATE INDEX idx_books_author ON books(author);
CREATE INDEX idx_books_genre ON books(genre);
CREATE INDEX idx_books_year ON books(year);
CREATE INDEX idx_books_price ON books(price);