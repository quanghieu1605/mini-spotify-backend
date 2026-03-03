CREATE DATABASE mini_spotify_db;
GO
USE mini_spotify_db;
GO
CREATE TABLE users (
    id INT IDENTITY(1,1) PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT GETDATE()
);