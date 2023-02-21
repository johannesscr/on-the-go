# Setup

Create books table
```postgresql
CREATE TABLE IF NOT EXISTS books
(
    isbn   CHAR(14) PRIMARY KEY,
    title  VARCHAR(255)  NOT NULL,
    author VARCHAR(255)  NOT NULL,
    price  DECIMAL(5, 2) NOT NULL
);
```

Insert data into books table
```postgresql
INSERT INTO books (isbn, title, author, price) 
VALUES ('978-1503261969', 'emma', 'Jayne Austin', 9.44),
       ('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
       ('978-1503379640', 'The Prince', 'Niccolo Machiavelli', 6.99);
```
