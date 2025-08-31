# User Service API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
This API uses **JWT (JSON Web Token)** authentication. Most endpoints require a valid Bearer token in the Authorization header.

### Authorization Header Format
```
Authorization: Bearer <jwt_token>
```

## Endpoints

### Health Check
Check if the service is running.

**GET** `/health`

**Response:** `200 OK`
```json
{
  "status": "ok",
  "message": "Service is running"
}
```

---

## Authentication Endpoints

### Register User
Create a new user account.

**POST** `/user_service/v1/auth/register`

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
  "message": "User registered successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "user_id": 1,
    "email": "user@example.com",
    "username": "username",
    "first_name": "John",
    "last_name": "Doe",
    "created_at": 1692816000,
    "updated_at": 1692816000
  }
}
```

### Login User
Authenticate an existing user.

**POST** `/user_service/v1/auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:** `200 OK`
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "user_id": 1,
    "email": "user@example.com",
    "username": "username",
    "first_name": "John",
    "last_name": "Doe",
    "created_at": 1692816000,
    "updated_at": 1692816000
  }
}
```

---

## User Management Endpoints
**🔒 All user endpoints require JWT authentication**

### Create User
Create a new user account (admin only).

**POST** `/user_service/v1/users/`
**Headers:** `Authorization: Bearer <token>`

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
  "user_id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "John",
  "last_name": "Doe",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

### Get User by ID
Retrieve a specific user by their ID.

**GET** `/user_service/v1/users/{id}`
**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK`
```json
{
  "user_id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "John",
  "last_name": "Doe",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

### Update User
Update an existing user's information.

**PUT** `/user_service/v1/users/{id}`
**Headers:** `Authorization: Bearer <token>`

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
  "user_id": 1,
  "email": "user@example.com",
  "username": "username",
  "first_name": "Jane",
  "last_name": "Smith",
  "created_at": 1692816000,
  "updated_at": 1692816000
}
```

### Delete User
Delete a user account.

**DELETE** `/user_service/v1/users/{id}`
**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK`
```json
{
  "message": "User deleted successfully"
}
```

---

## Conversation Management Endpoints
**🔒 All conversation endpoints require JWT authentication**

### Create New Conversation
Create a new conversation for the authenticated user.

**POST** `/user_service/v1/conversations/`
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "title": "My New Conversation",
  "model_used": "gpt-4"
}
```

**Response:** `201 Created`
```json
{
  "conversation_id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "My New Conversation",
  "model_used": "gpt-4",
  "created_at": "2024-01-15T10:30:00Z",
  "is_pinned": false
}
```

### Get All Conversations for User
Retrieve all conversations for a specific user.

**GET** `/user_service/v1/users/{id}/conversations`
**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK`
```json
{
  "conversations": [
    {
      "conversation_id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Chat about AI",
      "is_pinned": false,
      "updated_at": "2024-01-15T10:30:00Z"
    },
    {
      "conversation_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
      "title": "Programming Help",
      "is_pinned": true,
      "updated_at": "2024-01-14T15:45:00Z"
    }
  ]
}
```

### Add Message to Conversation
Add a new message to an existing conversation.

**POST** `/user_service/v1/conversations/{conversation_id}/messages`
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "message": "Hello, how can you help me today?",
  "sender": "user"
}
```

**Valid sender values:** `user`, `ai`, `system`

**Response:** `201 Created`
```json
{}
```

### Get Conversation History
Retrieve all messages from a specific conversation.

**GET** `/user_service/v1/conversations/{conversation_id}/history`
**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK`
```json
{
  "messages": [
    {
      "message": "Hello, how can you help me today?",
      "role": "user",
      "timestamp": "2024-01-15T10:30:00Z"
    },
    {
      "message": "I can help you with various tasks. What do you need assistance with?",
      "role": "ai",
      "timestamp": "2024-01-15T10:30:15Z"
    }
  ]
}
```

### Delete Conversation
Delete a conversation and all its messages. Only the conversation owner can delete it.

**DELETE** `/user_service/v1/conversations/{conversation_id}`
**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK`
```json
{
  "message": "Conversation deleted successfully"
}
```

**Response:** `403 Forbidden`
```json
{
  "error": "Access denied"
}
```

**Response:** `404 Not Found`
```json
{
  "error": "Conversation not found"
}
```

### Pin/Unpin Conversation
Toggle the pin status of a conversation. Only the conversation owner can pin/unpin it.

**PATCH** `/user_service/v1/conversations/{conversation_id}/pin`
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "is_pinned": true
}
```

**Response:** `200 OK`
```json
{
  "conversation_id": "550e8400-e29b-41d4-a716-446655440000",
  "is_pinned": true,
  "message": "Conversation pinned successfully"
}
```

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request data"
}
```

### 401 Unauthorized
```json
{
  "error": "Authorization header required"
}
```
```json
{
  "error": "Invalid or expired token"
}
```

### 403 Forbidden
```json
{
  "error": "Access denied"
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

---

## Data Models

### User
```json
{
  "user_id": "uint (primary key)",
  "email": "string (unique, required)",
  "username": "string (unique, required)",
  "password": "string (required, hidden in responses)",
  "first_name": "string",
  "last_name": "string",
  "created_at": "int64 (Unix timestamp)",
  "updated_at": "int64 (Unix timestamp)"
}
```

### Conversation
```json
{
  "conversation_id": "UUID (primary key)",
  "user_id": "uint (foreign key)",
  "title": "string (required)",
  "model_used": "string (optional)",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "is_pinned": "boolean"
}
```

### Message
```json
{
  "message_id": "UUID (primary key)",
  "conversation_id": "UUID (foreign key)",
  "parent_message_id": "UUID (optional, foreign key)",
  "sender": "string (enum: 'user', 'ai', 'system')",
  "content": "string (required)",
  "metadata": "JSON object (optional)",
  "timestamp": "timestamp"
}
```

---

## JWT Token Structure

### Claims
```json
{
  "user_id": 1,
  "email": "user@example.com",
  "username": "username",
  "exp": 1692902400,
  "iat": 1692816000
}
```

### Token Expiry
- **Default:** 30 days (720 hours)
- **Format:** Bearer token in Authorization header

---

## Example Usage

### 1. Register and Login
```bash
# Register
curl -X POST http://localhost:8080/user_service/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "username": "john",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'

# Login
curl -X POST http://localhost:8080/user_service/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 2. Use Protected Endpoints
```bash
# Create new conversation
curl -X POST http://localhost:8080/user_service/v1/conversations/ \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My AI Chat",
    "model_used": "gpt-4"
  }'

# Get user conversations
curl -X GET http://localhost:8080/user_service/v1/users/1/conversations \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# Add message to conversation
curl -X POST http://localhost:8080/user_service/v1/conversations/550e8400-e29b-41d4-a716-446655440000/messages \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Hello!",
    "sender": "user"
  }'

# Get conversation history
curl -X GET http://localhost:8080/user_service/v1/conversations/550e8400-e29b-41d4-a716-446655440000/history \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# Delete conversation
curl -X DELETE http://localhost:8080/user_service/v1/conversations/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# Pin conversation
curl -X PATCH http://localhost:8080/user_service/v1/conversations/550e8400-e29b-41d4-a716-446655440000/pin \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{"is_pinned": true}'

# Unpin conversation
curl -X PATCH http://localhost:8080/user_service/v1/conversations/550e8400-e29b-41d4-a716-446655440000/pin \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{"is_pinned": false}'
```

---

## Environment Configuration

Required environment variables in `.env`:
```bash
# Server Configuration
PORT=8080
ENVIRONMENT=development

# Database Configuration
DATABASE_URL=postgresql://user:password@host:port/database

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
```

---

## Security Notes

1. **JWT Tokens:** Store securely on client side (httpOnly cookies recommended)
2. **HTTPS:** Use HTTPS in production
3. **Rate Limiting:** Implement rate limiting for production
4. **CORS:** Configure CORS appropriately for production
5. **Environment Variables:** Never commit secrets to version control