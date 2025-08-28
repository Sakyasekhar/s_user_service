# User Service - Go Backend

A robust Go backend service built with Gin framework, GORM, and PostgreSQL.

## Features

- RESTful API with CRUD operations
- PostgreSQL database with GORM ORM
- Middleware for logging, CORS, and recovery
- Environment-based configuration
- Clean architecture with separation of concerns
- Health check endpoint
- User management system

## Project Structure

```
user_service/
├── config/          # Configuration management
├── internal/        # Internal application code
│   ├── database/    # Database models and connection
│   ├── handlers/    # HTTP request handlers
│   ├── middleware/  # HTTP middleware
│   └── routes/      # Route definitions
├── main.go          # Application entry point
├── go.mod           # Go module file
├── env.example      # Environment variables example
└── README.md        # This file
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd user_service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables:
```bash
cp env.example .env
# Edit .env with your database credentials
```

4. Set up PostgreSQL database:
```sql
CREATE DATABASE user_service;
CREATE USER postgres WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE user_service TO postgres;
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

2. The server will start on port 8080 (or the port specified in your .env file)

## API Endpoints

### Health Check
- `GET /health` - Service health status

### Users
- `POST /api/v1/users/` - Create a new user
- `GET /api/v1/users/` - Get all users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user by ID
- `DELETE /api/v1/users/:id` - Delete user by ID

### Query Parameters for Users List
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10)

## Example API Usage

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "username",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

### Get Users
```bash
curl http://localhost:8080/api/v1/users/?page=1&limit=5
```

### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | 8080 |
| `ENVIRONMENT` | Environment (development/production) | development |
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 5432 |
| `DB_USER` | Database username | postgres |
| `DB_PASSWORD` | Database password | password |
| `DB_NAME` | Database name | user_service |
| `DB_SSLMODE` | Database SSL mode | disable |

## Development

### Adding New Models
1. Create a new struct in `internal/database/database.go`
2. Add it to the `AutoMigrate` call in `InitDB`
3. Create corresponding handlers in `internal/handlers/`
4. Add routes in `internal/routes/routes.go`

### Adding New Middleware
1. Create new middleware functions in `internal/middleware/middleware.go`
2. Apply them in `main.go` or specific route groups

## Testing

Run tests:
```bash
go test ./...
```

## Production Deployment

1. Set `ENVIRONMENT=production` in your environment
2. Ensure proper database credentials and SSL configuration
3. Use a reverse proxy (nginx) for production
4. Set up proper logging and monitoring
5. Use environment-specific configuration files

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
