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

INSERT INTO books (book_name, author, genre, year, count, price)
VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', 'Classic', 1925, 30, 15.99),
       ('To Kill a Mockingbird', 'Harper Lee', 'Fiction', 1960, 25, 18.99),
       ('1984', 'George Orwell', 'Dystopian', 1949, 40, 19.99),
       ('Pride and Prejudice', 'Jane Austen', 'Romance', 1813, 15, 12.99),
       ('Moby-Dick', 'Herman Melville', 'Adventure', 1851, 20, 22.99),
       ('War and Peace', 'Leo Tolstoy', 'Historical Fiction', 1869, 12, 14.99),
       ('The Catcher in the Rye', 'J.D. Salinger', 'Fiction', 1951, 18, 10.99),
       ('The Hobbit', 'J.R.R. Tolkien', 'Fantasy', 1937, 35, 25.99),
       ('The Odyssey', 'Homer', 'Epic', -700, 50, 9.99), -- Year is an approximation
       ('The Divine Comedy', 'Dante Alighieri', 'Epic', 1320, 20, 30.99);