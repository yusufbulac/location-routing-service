# Location Routing Service

A backend service built with Golang and Gin framework for managing and routing geographic locations.

## Features

- Add new locations with name, coordinates and custom marker color
- List all saved locations
- View detailed information for a specific location
- Edit existing location data
- Generate a simple point-to-point route based on geographical proximity (bird's-eye view)
- Input validation using go-playground/validator
- Rate limiting (per IP)
- Swagger/OpenAPI documentation (```/swagger/index.html```)
- Unit tests for service layer (testify)
- Dockerized application with MySQL
- Layered architecture (handler, service, repository)
- Optional: Test coverage, CI/CD pipeline, and deployment support

## Tech Stack

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM ORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [Validator.v10](https://github.com/go-playground/validator)
- [Swag CLI](https://github.com/swaggo/swag)
- [Testify](https://github.com/stretchr/testify)
- [gin-contrib/limiter](https://github.com/gin-contrib/limiter)

## Project Structure

```
├── cmd/                    # Main entrypoint (main.go)
├── internal/              # Application logic (modularized)
│   ├── config/            # Configuration and database connection
│   ├── dto/               # Request and response structures
│   ├── handler/           # HTTP layer / API handlers
│   ├── model/             # GORM models
│   ├── repository/        # DB access logic
│   ├── service/           # Business logic
│   └── middleware/        # Custom middleware (rate limiting, etc.)
├── docs/                  # Auto-generated Swagger files
├── Dockerfile             # Container build definition
├── docker-compose.yml     # Multi-container orchestration
├── .env.example           # Environment config template
└── go.mod / go.sum        # Dependencies
```

## Run Locally

```bash
# Clone the repo
git clone https://github.com/yusufbulac/location-routing-service.git
cd location-routing-service

# Create .env from example
cp .env.example .env

# Build and run with Docker
docker-compose up --build
```

Then visit:

http://localhost:8080/swagger/index.html → Swagger UI

## Testing

Unit tests written using testify for service layer:

```bash
go test ./internal/service/...
```

## Notes

- All environment variables should be configured in ```.env``` (see ```.env.example```).

- Swagger documentation is auto-generated using Swag CLI (```swag init -g cmd/main.go```).

- API versioning and graceful shutdown are planned as future enhanceme