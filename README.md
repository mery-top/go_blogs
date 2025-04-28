# Go Blog App

A simple blog application built with Go, MySQL, and basic HTTP handlers. This app allows you to create and view blog posts through API endpoints.

## Features
- **View all posts** via `/api/posts` endpoint.
- **Create a new post** via `/api/create` endpoint (POST request with title and content).

## Technologies Used
- **Go (Golang)** for backend.
- **MySQL** for database.

## Set Up database
```
CREATE TABLE posts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL
);
```

##Update the DB Credentials
```
dsn := "root:password@tcp(127.0.0.1:3307)/blog_db"
```
