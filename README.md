# Simple CRUD API

A simple CRUD API for Document Management Using Go (Gin framework), GORM, and PostgreSQL.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Running the Application](#running-the-application)
  - [Using Docker Compose (Recommended for DB)](#using-docker-compose-recommended-for-db)
  - [Using Go](#using-go)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

-   Go (check `go.mod` for version, e.g., 1.18 or later) - [Link to Go installation](https://golang.org/doc/install)
-   Docker & Docker Compose (Recommended for running PostgreSQL easily) - [Link to Docker installation](https://docs.docker.com/get-docker/)
-   Git

### Installation

1.  **Clone the repository:**
    ```bash
    git clone <your-repository-url> # Replace with your repo URL
    cd <project-directory>          # Replace with the cloned directory name
    ```

2.  **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

## Running the Application

You can run the application in different ways. Using Docker Compose for the database is recommended.

### Using Docker Compose (Recommended for DB)

The provided `docker-compose.yml` file makes it easy to start the required PostgreSQL database service.

1.  **Start the database service:**
    ```bash
    docker-compose up -d db
    ```
    This command starts a PostgreSQL container named `document_system_db` in the background, listening on port `5432`. It uses the credentials and database name defined in `docker-compose.yml`.

2.  **Run the Go application:**
    After the database is running, you can run the Go application directly (see next section) . The application is currently configured in `config/database.go` to connect to `host=localhost user=postgres password=postgres dbname=documents_db port=5432`.

### Using Go

1.  **Ensure Prerequisites are met:** Make sure you have Go installed and a PostgreSQL database running and accessible. If you used `docker-compose up -d db` from the previous step, the database is ready.
    *(Note: The current database connection string (DSN) is hardcoded in `config/database.go`. It's recommended to modify the code to use environment variables for flexibility. See the [Configuration](#configuration) section.)*

2.  **Run the application:**
    ```bash
    go run main.go
    ```
    The application will start, connect to the database, perform migrations, and listen on port `8080` by default (as seen in `main.go`). You should see output like:
    ```
    Database connected!
    Database Migrated!
    [GIN-debug] Listening and serving HTTP on :8080
    ```

## Configuration

The application currently uses a hardcoded database connection string in `config/database.go`:
`dsn := "host=localhost user=postgres password=postgres dbname=documents_db port=5432 sslmode=disable"`


## API Documentation

The following endpoints are available (base path `/`):

### Documents

#### Create Document

-   **Method:** `POST`
-   **Path:** `/documents`
-   **Description:** Creates a new document.
-   **Request Body:**
    ```json
    {
      "title": "My First Document",
      "content": "This is the content of the document.",
      "author": "John Doe"
    }
    ```
    *(`title` and `content` are required)*
-   **Success Response:**
    -   **Code:** `201 Created`
    -   **Content:** The created document object including `id`, `created_at`, `updated_at`.
        ```json
        {
          "id": 1,
          "title": "My First Document",
          "content": "This is the content of the document.",
          "author": "John Doe",
          "created_at": "2023-10-27T10:00:00Z",
          "updated_at": "2023-10-27T10:00:00Z"
        }
        ```
-   **Error Response:**
    -   **Code:** `400 Bad Request` (e.g., validation error on required fields)
    -   **Content:** `{"error": "Key: 'Document.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`
    -   **Code:** `500 Internal Server Error` (e.g., database error)
    -   **Content:** `{"error": "database error description"}`

#### Get All Documents

-   **Method:** `GET`
-   **Path:** `/documents`
-   **Description:** Retrieves a list of all documents.
-   **Request Body:** None
-   **Success Response:**
    -   **Code:** `200 OK`
    -   **Content:** An array of document objects.
        ```json
        [
          {
            "id": 1,
            "title": "My First Document",
            "content": "This is the content.",
            "author": "John Doe",
            "created_at": "2023-10-27T10:00:00Z",
            "updated_at": "2023-10-27T10:00:00Z"
          },
          {
            "id": 2,
            "title": "Another Document",
            "content": "More content here.",
            "author": "Jane Smith",
            "created_at": "2023-10-27T11:00:00Z",
            "updated_at": "2023-10-27T11:00:00Z"
          }
        ]
        ```
-   **Error Response:**
    -   **Code:** `500 Internal Server Error`
    -   **Content:** `{"error": "database error description"}`

#### Get Document by ID

-   **Method:** `GET`
-   **Path:** `/documents/{id}`
-   **Description:** Retrieves a specific document by its ID.
-   **URL Parameters:**
    -   `id=[uint]` (Required) - The ID of the document to retrieve.
-   **Request Body:** None
-   **Success Response:**
    -   **Code:** `200 OK`
    -   **Content:** The requested document object.
        ```json
        {
          "id": 1,
          "title": "My First Document",
          "content": "This is the content.",
          "author": "John Doe",
          "created_at": "2023-10-27T10:00:00Z",
          "updated_at": "2023-10-27T10:00:00Z"
        }
        ```
-   **Error Response:**
    -   **Code:** `404 Not Found`
    -   **Content:** `{"error": "Document not found"}`
    -   **Code:** `500 Internal Server Error`

#### Update Document

-   **Method:** `PUT`
-   **Path:** `/documents/{id}`
-   **Description:** Updates an existing document by its ID.
-   **URL Parameters:**
    -   `id=[uint]` (Required) - The ID of the document to update.
-   **Request Body:** The fields to update.
    ```json
    {
      "title": "Updated Document Title",
      "content": "Updated content.",
      "author": "John Doe"
    }
    ```
    *(`title` and `content` are required for validation)*
-   **Success Response:**
    -   **Code:** `200 OK`
    -   **Content:** The updated document object.
        ```json
        {
          "id": 1,
          "title": "Updated Document Title",
          "content": "Updated content.",
          "author": "John Doe",
          "created_at": "2023-10-27T10:00:00Z", 
          "updated_at": "2023-10-27T12:00:00Z" 
        }
        ```
-   **Error Response:**
    -   **Code:** `400 Bad Request` (e.g., validation error)
    -   **Content:** `{"error": "Validation error description"}`
    -   **Code:** `404 Not Found`
    -   **Content:** `{"error": "Document not found"}`
    -   **Code:** `500 Internal Server Error`

#### Delete Document

-   **Method:** `DELETE`
-   **Path:** `/documents/{id}`
-   **Description:** Deletes a document by its ID.
-   **URL Parameters:**
    -   `id=[uint]` (Required) - The ID of the document to delete.
-   **Request Body:** None
-   **Success Response:**
    -   **Code:** `200 OK`
    -   **Content:** `{"message": "Document deleted"}`
-   **Error Response:**
    -   **Code:** `404 Not Found`
    -   **Content:** `{"error": "Document not found"}`
    -   **Code:** `500 Internal Server Error`

---



