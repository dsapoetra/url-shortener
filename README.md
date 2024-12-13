# URL Shortener Backend

This is a backend project for a URL shortener service, built using **Golang** with modular architecture.

## Features
- Shorten URLs and retrieve original URLs.
- Support for database migrations.
- Configurable via environment variables.
- Lightweight and efficient architecture.

## Project Structure
```
url-shortener-be
├── cmd/                # Main application entry point
├── config/             # Configuration files
├── db/                 # Database initialization and connection
├── handlers/           # HTTP handlers for routes
├── migrations/         # Database migration files
├── models/             # Data models
├── repositories/       # Data access layer
├── routes/             # Route definitions
├── services/           # Business logic layer
├── Makefile            # Commands for building and running
├── .env                # Environment variables file
├── .env_sample         # Sample environment configuration
├── go.mod              # Module definition
├── go.sum              # Dependency checksums
├── README.md           # Project documentation
└── .gitignore          # Ignored files for Git
```

## Prerequisites
- Go 1.20 or later
- PostgreSQL

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd url-shortener-be
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment variables:
   - Copy the `.env_sample` file to `.env`:
     ```bash
     cp .env_sample .env
     ```
   - Update `.env` with your PostgreSQL configuration.

4. Run database migrations:
   ```bash
   make migrate
   ```

5. Start the server:
   ```bash
   go run cmd/main.go
   ```

## Usage
- API Endpoints:
  - **POST /api/urls/**: Shorten a URL.
  - **GET /:shortcode**: Retrieve the original URL from a shortcode.

## Development
- To run tests:
  ```bash
  make test
  ```
- To format code:
  ```bash
  make fmt
  ```

## Contributing
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- [Fiber](https://gofiber.io/) for HTTP handling.
- Community contributors.
