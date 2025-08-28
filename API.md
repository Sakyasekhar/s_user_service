# User Service API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
Currently, the API does not require authentication. In production, implement JWT or OAuth2.

## Endpoints

### Health Check
Check if the service is running.

**GET** `/health`

**Response:**
```json
{
  "status": "ok",
  "message": "Service is running"
}
```

### Users

#### Create User
Create a new user account.

**POST** `/api/v1/users/`

**Request Body:**
```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "John",
  "last_name": "Doe",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

#### Get All Users
Retrieve a paginated list of users.

**GET** `/api/v1/users/`

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:** `200 OK`
```json
{
  "users": [
    {
      "id": 1,
      "email": "user@example.com",
      "username": "username",
      "first_name": "John",
      "last_name": "Doe",
      "created_at": 1692816000,
      "updated_at": 1692816000
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 10
}
```

#### Get User by ID
Retrieve a specific user by their ID.

**GET** `/api/v1/users/{id}`

**Response:** `200 OK`
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "John",
  "last_name": "Doe",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

**Response:** `404 Not Found`
```json
{
  "error": "User not found"
}
```

#### Update User
Update an existing user's information.

**PUT** `/api/v1/users/{id}`

**Request Body:**
```json
{
  "first_name": "Jane",
  "last_name": "Smith"
}
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "Jane",
  "last_name": "Smith",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

#### Delete User
Delete a user account.

**DELETE** `/api/v1/users/{id}`

**Response:** `200 OK`
```json
{
  "message": "User deleted successfully"
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "error": "Invalid request data"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

## Data Models

### User
```json
{
  "id": "uint",
  "email": "string (unique, required)",
  "username": "string (unique, required)",
  "password": "string (required, hidden in responses)",
  "first_name": "string",
  "last_name": "string",
  "created_at": "int64 (timestamp)",
  "updated_at": "int64 (timestamp)"
}
```

## Rate Limiting
Currently, no rate limiting is implemented. Consider implementing rate limiting for production use.

## CORS
CORS is enabled for all origins in development. Configure appropriately for production.
