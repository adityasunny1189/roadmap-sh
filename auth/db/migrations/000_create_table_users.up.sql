CREATE TABLE IF NOT EXISTS users (
    userId VARCHAR(30) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    user_password VARCHAR(50) NOT NULL 
);
