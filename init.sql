-- init.sql
CREATE TABLE IF NOT EXISTS Customer (
    customerid VARCHAR(10) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(15) NOT NULL
);