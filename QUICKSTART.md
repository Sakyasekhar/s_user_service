# Quick Start Guide

Get your Go backend up and running in minutes!

## Prerequisites
- Go 1.21+ installed
- Docker and Docker Compose (optional, for containerized development)
- PostgreSQL (if running locally)

## Option 1: Local Development (Recommended for beginners)

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Set up Environment
```bash
cp env.example .env
# Edit .env with your database credentials
```

### 3. Set up PostgreSQL Database
```sql
CREATE DATABASE user_service;
CREATE USER postgres WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE user_service TO postgres;
```

### 4. Run the Application
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Option 2: Docker Development (Recommended for team development)

### 1. Start Everything with Docker
```bash
./scripts/dev.sh start
```

This will:
- Start PostgreSQL database
- Build and start your Go application
- Set up all necessary environment variables

### 2. Check Status
```bash
./scripts/dev.sh logs
```

### 3. Stop Everything
```bash
./scripts/dev.sh stop
```

## Test Your API

### Health Check
```bash
curl http://localhost:8080/health
```

### Create a User
```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

### Get All Users
```bash
curl http://localhost:8080/api/v1/users/
```

## Development Commands

### Build
```bash
make build
# or
go build -o bin/user_service main.go
```

### Run Tests
```bash
make test
# or
go test ./...
```

### Run with Hot Reload (requires air)
```bash
make dev
```

### Clean Up
```bash
make clean
```

## Project Structure
```
user_service/
├── config/          # Configuration management
├── internal/        # Internal application code
│   ├── database/    # Database models and connection
│   ├── handlers/    # HTTP request handlers
│   ├── middleware/  # HTTP middleware
│   └── routes/      # Route definitions
├── scripts/         # Development scripts
├── main.go          # Application entry point
├── go.mod           # Go module file
├── Dockerfile       # Container configuration
├── docker-compose.yml # Development environment
└── README.md        # Full documentation
```

## Next Steps

1. **Add Authentication**: Implement JWT or OAuth2
2. **Add Validation**: Use a validation library like `go-playground/validator`
3. **Add Logging**: Implement structured logging with `logrus` or `zap`
4. **Add Testing**: Write comprehensive tests for your handlers
5. **Add Documentation**: Use Swagger/OpenAPI for API documentation
6. **Add Monitoring**: Implement health checks and metrics

## Troubleshooting

### Port Already in Use
If port 8080 is already in use, change the `PORT` in your `.env` file.

### Database Connection Issues
- Ensure PostgreSQL is running
- Check your database credentials in `.env`
- Verify the database exists

### Build Errors
- Ensure Go 1.21+ is installed
- Run `go mod tidy` to fix dependency issues
- Check for syntax errors in your code

## Need Help?

- Check the full [README.md](README.md)
- Review the [API documentation](API.md)
- Run `./scripts/dev.sh help` for available commands
