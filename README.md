# Location Routing Service

A backend service built with Golang and Gin framework for managing and routing geographic locations.

## Features

- Add new locations with name, coordinates and custom marker color
- List all saved locations
- View detailed information for a specific location
- Edit existing location data
- Generate a simple point-to-point route based on geographical proximity (bird's-eye view)
- Input validation and rate limiting
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
- [gin-contrib/limiter](https://github.com/gin-contrib/limiter)

## Project Structure

```
├── cmd/ 
├── internal/ 
│ ├── handler/
│ ├── service/
│ ├── repository/
│ ├── model/
│ └── config/
├── Dockerfile
├── docker-compose.yml
└── main.go
```

## Run Locally

```bash
# Clone the repo
git clone https://github.com/yusufbulac/location-routing-service.git
cd location-routing-service

cp .env.example .env

# Start services
docker-compose up --build
```
