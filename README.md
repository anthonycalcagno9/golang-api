# Golang API

A simple Go API for sending meal data to the front end. Used gorilla mux router and listenAndServe to serve traffic on port 8080.

## Structure

- `models/` - Data models
- `handlers/` - HTTP handlers
- `main.go` - Entry point

## Getting Started

1. Install Go dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

## Models

- `Meal` - Represents a daily meal plan with breakfast, lunch, dinner, rating, and day of the week
