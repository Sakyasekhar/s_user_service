#!/bin/bash

# Development script for user_service

set -e

case "$1" in
    "start")
        echo "Starting development environment..."
        docker-compose up -d
        echo "Development environment started!"
        echo "API available at: http://localhost:8080"
        echo "Health check: http://localhost:8080/health"
        ;;
    "stop")
        echo "Stopping development environment..."
        docker-compose down
        echo "Development environment stopped!"
        ;;
    "restart")
        echo "Restarting development environment..."
        docker-compose restart
        echo "Development environment restarted!"
        ;;
    "logs")
        docker-compose logs -f
        ;;
    "build")
        echo "Building application..."
        go build -o bin/user_service main.go
        echo "Build complete!"
        ;;
    "test")
        echo "Running tests..."
        go test -v ./...
        ;;
    "clean")
        echo "Cleaning build artifacts..."
        rm -rf bin/
        docker-compose down -v
        echo "Clean complete!"
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|logs|build|test|clean}"
        echo ""
        echo "Commands:"
        echo "  start   - Start development environment with Docker"
        echo "  stop    - Stop development environment"
        echo "  restart - Restart development environment"
        echo "  logs    - Show application logs"
        echo "  build   - Build the application"
        echo "  test    - Run tests"
        echo "  clean   - Clean build artifacts and containers"
        exit 1
        ;;
esac
