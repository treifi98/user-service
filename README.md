# User CRUD Service

This project is a simple CRUD (Create, Read, Update, Delete) service for managing user data, built with Go and SQLite. It provides a RESTful API for user management operations.

## Project Structure

```
user-crud-service/
├── cmd/
│   └── server/
│       └── main.go           # Entry point of the application
├── internal/
│   ├── model/
│   │   └── user.go           # User model definition
│   ├── database/
│   │   └── database.go       # Database connection and initialization
│   ├── handler/
│   │   └── user_handler.go   # HTTP handlers for user operations
│   └── service/
│       └── user_service.go   # Business logic for user operations
├── go.mod                    # Go module file
└── go.sum                    # Go module checksum file
```

## Prerequisites

- Go 1.16 or later
- SQLite

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/user-crud-service.git
   cd user-crud-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Running the Application

To start the server, run:

```
./server
```

The server will start on `localhost:5000`.

## API Routes

The following API routes are available:

1. Create a new user
   - Method: POST
   - URL: `/users`
   - Body: JSON object with `username`, `password`, and `active` fields
   - Example:
     ```
     POST http://localhost:5000/users
     Content-Type: application/json

     {
       "username": "johndoe",
       "password": "secretpassword",
       "active": true
     }
     ```

2. Get a user by ID
   - Method: GET
   - URL: `/users/{id}`
   - Example: `GET http://localhost:5000/users/1`

3. Update a user
   - Method: PUT
   - URL: `/users/{id}`
   - Body: JSON object with `username`, `password`, and `active` fields
   - Example:
     ```
     PUT http://localhost:5000/users/1
     Content-Type: application/json

     {
       "username": "johndoe",
       "password": "newpassword",
       "active": false
     }
     ```

4. Delete a user
   - Method: DELETE
   - URL: `/users/{id}`
   - Example: `DELETE http://localhost:5000/users/1`

5. Get all users
   - Method: GET
   - URL: `/users`
   - Example: `GET http://localhost:5000/users`

## Testing

You can test the API using tools like cURL, Postman, or any HTTP client. Here are some example cURL commands:

1. Create a user:
   ```
   curl -X POST -H "Content-Type: application/json" -d '{"username":"johndoe","password":"secretpassword","active":true}' http://localhost:5000/users
   ```

2. Get a user:
   ```
   curl http://localhost:5000/users/1
   ```

3. Update a user:
   ```
   curl -X PUT -H "Content-Type: application/json" -d '{"username":"johndoe","password":"newpassword","active":false}' http://localhost:5000/users/1
   ```

4. Delete a user:
   ```
   curl -X DELETE http://localhost:5000/users/1
   ```

5. Get all users:
   ```
   curl http://localhost:5000/users
   ```



