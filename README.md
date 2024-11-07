# Go_Server

This is a simple Go server for handling CRUD operations. It is based on a tutorial by [arthur404dev](https://github.com/arthur404dev).

## Features
- Basic CRUD endpoints for managing data entries.
- RESTful structure using [Gin](https://github.com/gin-gonic/gin).
- Database integration via [GORM](https://gorm.io/).

## Getting Started

### Prerequisites
- [Go](https://golang.org/) installed on your system.
- [Air](https://github.com/cosmtrek/air) (optional) for live reloading during development.

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/Go_Server.git
    cd Go_Server
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

### Running the Server
To start the server, you can either:
- Use `air` for live reloading during development:
    ```bash
    air
    ```
- Or, if you prefer to run without live reloading, start the server with:
    ```bash
    go run main.go
    ```

## API Endpoints
- `GET /api/v1/openings` - List all entries
- `POST /api/v1/opening` - Create a new entry
- `GET /api/v1/opening?id={id}` - Retrieve a specific entry
- `PUT /api/v1/opening?id={id}` - Update an entry
- `DELETE /api/v1/opening?id={id}` - Delete an entry
