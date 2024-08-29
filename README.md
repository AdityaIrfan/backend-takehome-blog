# Blog Application

Welcome to the **Blog Application**! This Go-based application leverages the Echo framework and a MySQL database to provide a robust platform for creating and managing blog posts and comments, including nested comments.

## Overview

This application allows users to:
- Create new blog posts
- Add comments to posts
- Add nested comments for deeper discussions

## Features

- **Register**
- **Login**
- **Post Creation**: Users can create new posts with titles and content.
- **Commenting**: Users can add comments to posts.
- **Nested Comments**: Support for adding comments on comments, creating a threaded discussion.
- **Post Management**: Easily manage posts with features like updating and deleting.

## Technology Stack

: The primary programming language used for the application.
- **Echo**: A high-performance, extensible, and minimalist web framework for Go.
- **MySQL**: A relational database for storing posts and comments.
- **ULID**: Universally Unique Lexicographically Sortable Identifier for generating unique IDs.

## Quick Getting Started with Air (Hot Reload)
Air is a live-reloading tool for Go applications. It watches for changes in your files and automatically reloads the application. Follow these steps to get the application up and running on your local machine:

### Prerequisites

- [Go](https://golang.org/doc/install) (1.18+)
- [Docker](https://docs.docker.com/get-docker/) (optional, for running MySQL in a container)

### Installation

```bash
cd {project}/app
go install github.com/air-verse/air@latest
air
```

## Getting Started with _docker compose_

Follow these steps to get the application up and running on your local machine:

### Prerequisites

- [Go](https://golang.org/doc/install) (1.18+)
- [Docker](https://docs.docker.com/get-docker/)

### Installation

```bash
docker compose build
docker compose up
```

## Getting Started Manually

Follow these steps to get the application up and running on your local machine:

### Prerequisites

- [Go](https://golang.org/doc/install) (1.18+)
- [MySQL](https://dev.mysql.com/downloads/) (5.7+)
- [Docker](https://docs.docker.com/get-docker/) (optional, for running MySQL in a container)

### Setup Database

1. **Using Docker:**

   ```bash
   docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db -p 3306:3306 -d mysql:latest

2. **Manually Setting Up MySQL:**
    - Install MySQL and create a database named blog_db.

3. **ENV**

    Update your .env file with the correct database connection details:
    ```bash
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=root
    DB_NAME=blog_db
    ```

## Install Dependencies

Ensure you have Go installed, then install the project dependencies:
```bash
go run main.go
```

## Run the Application
To start the application, run:
```bash
go run main.go
```
The application will be available at http://localhost:8080.
## API Endpoints

- ### Register
    - **Path**: `/register`
    - **Method**: `POST`
    - **Request Body**:
        ```json
        {
            "Name": "irfan",
            "Email": "irfan@gmail.com",
            "Password": "password"
        }
    - **Response**
        ```json
        {
            "Success": false,
            "StatusCode": 201,
            "Message": "register success"
        }
- ### Login
    - **Path**: `/login`
    - **Method**: `POST`
    - **Request Body**:
        ```json
        {
            "Email": "irfan@gmail.com",
            "Password": "password"
        }
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "login success",
            "Data": {
                "Token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOjE3MjUwMzg3MDksIklkIjoiMDFKNkZKUEswM0IxUTBGREhDN0c5NFNKWUgifQ.P1C1jD3Rh6Efv1q3Cgr8KmYsHLWIpAFIUVdBs5_DtyUU9iRWpBH9yj0cf9a5apWQRcuedqGosfuPv_dTRehKZ_ZLGbGOqVn4Y_C8rchBFl8cG-vYn85nT6lxBWdkXCHfThBqMcmlpgeYYpml93gdK-GoyZCiViAqi9SuwQlAnaM"
            }
        }
- ### Create Post
    - **Path**: `/posts`
    - **Method**: `POST`
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Request Body**:
        ```json
        {
            "Title": "dari irfan",
            "Content": "kalian tau ga?"
        }
    - **Response**
        ```json
        {
            "Success": false,
            "StatusCode": 201,
            "Message": "success create post",
            "Data": {
                "Id": "01J6FJYGFDJ0YV0QM02WVY7HKG",
                "Title": "dari irfan",
                "Content": "kalian tau ga?",
                "CreatedAt": "2024-08-29 17:29:20",
                "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                "Author": "irfan",
                "TotalComments": 0
            }
        }
- ### Get Detail Post
    - **Path**: `/posts/:id`
    - **Method**: `GET`
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Response**
        ```json
        {
            "Success": false,
            "StatusCode": 201,
            "Message": "success create post",
            "Data": {
                "Id": "01J6FJYGFDJ0YV0QM02WVY7HKG",
                "Title": "dari irfan",
                "Content": "kalian tau ga?",
                "CreatedAt": "2024-08-29 17:29:20",
                "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                "Author": "irfan",
                "TotalComments": 0
            }
        }
- ### Update Post
    - **Path**: `/posts/:id`
    - **Method**: `PUT`
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Request Body**:
        ```json
        {
            "Title": "dari irfan",
            "Content": "kalian tau ga?"
        }
    - **Response**
        ```json
        {
            "Success": false,
            "StatusCode": 201,
            "Message": "success create post",
            "Data": {
                "Id": "01J6FJYGFDJ0YV0QM02WVY7HKG",
                "Title": "dari irfan",
                "Content": "kalian tau ga?",
                "CreatedAt": "2024-08-29 17:29:20",
                "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                "Author": "irfan",
                "TotalComments": 0
            }
        }
- ### Get All Posts
    - **Path**: `/posts`
    - **Method**: `GET`
    - **Query Params**
        ```json
        "Perpage" : (number, default 1),
        "Next"    : {{Next}} (next cursor),
        "Prev"    : {{Prev}} (prev cursor)
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Request Body**:
        ```json
        {
            "Title": "dari irfan",
            "Content": "kalian tau ga?"
        }
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "success get list posts",
            "Data": [
                {
                    "Id": "01J6FJS431TAGE4M3ZKESVYRB0",
                    "Title": "dari irfan",
                    "CreatedAt": "2024-08-29 17:26:24",
                    "TotalComments": 0
                }
            ],
            "Meta": {
                "Next": "",
                "Prev": ""
            }
        }
- ### Get All My Posts
    - **Path**: `my/posts`
    - **Method**: `GET`
    - **Query Params**
        ```json
        "Perpage" : (number, default 1),
        "Next"    : {{Next}} (next cursor),
        "Prev"    : {{Prev}} (prev cursor)
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Request Body**:
        ```json
        {
            "Title": "dari irfan",
            "Content": "kalian tau ga?"
        }
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "success get list posts",
            "Data": [
                {
                    "Id": "01J6FJS431TAGE4M3ZKESVYRB0",
                    "Title": "dari irfan",
                    "CreatedAt": "2024-08-29 17:26:24",
                    "TotalComments": 0
                }
            ],
            "Meta": {
                "Next": "",
                "Prev": ""
            }
        }
- ### Delete Post
    - **Path**: `/posts/:id`
    - **Method**: `DELETE`
    - **Query Params**
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "success delete post"
        }
- ### Create Comment
    - **Path**: `my/posts`
    - **Method**: `GET`
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Request Body**:
        ```json
        {
            "Content": "ada kok banyak",
            "ParentId": "01J6FK8JR055AQBD92WRAGY7EX" // optional
        }
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "success create comment",
            "Data": {
                "Id": "01J6FK9EWMZF2WM60MBWJ0P9JS",
                "Content": "ada kok banyak",
                "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                "AuthorName": "irfan",
                "TotalComment": 0
            }
        }
- ### Get All Comments
    - **Path**: `/posts/:id/comments`
    - **Method**: `GET`
    - **Query Params**
        ```json
        "Perpage" : (number, default 1),
        "Next"    : {{Next}} (next cursor),
        "Prev"    : {{Prev}} (prev cursor)
    - **Request Header**
        ```json
        "Authorization" : "Bearer {{Token}}"
    - **Response**
        ```json
        {
            "Success": true,
            "StatusCode": 200,
            "Message": "success get all comments",
            "Data": [
                {
                    "Id": "01J6FK61Z92Z5Y1PN4P4EW69AQ",
                    "Content": "ada lagi ih",
                    "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                    "AuthorName": "irfan",
                    "TotalComment": 0
                },
                {
                    "Id": "01J6FK8JR055AQBD92WRAGY7EX",
                    "Content": "ini ada jurnalnya ga?",
                    "AuthorId": "01J6FJPK03B1Q0FDHC7G94SJYH",
                    "AuthorName": "irfan",
                    "TotalComment": 2
                }
            ],
            "Meta": {
                "Next": "",
                "Prev": ""
            }
        }

### ...... Happy Blogging!!!# backend-takehome-blog
