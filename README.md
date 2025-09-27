# Cowsay Fortune Server

A Go-based web service that serves random fortunes formatted with ASCII art cows. This project combines the classic Unix `fortune` and `cowsay` commands into a modern HTTP API.

## Features

- **Random Fortunes**: Serves random quotes and jokes from a curated collection
- **ASCII Art**: Formats fortunes with the iconic cowsay cow in ASCII art
- **HTTP API**: Simple REST endpoint for fortune retrieval
- **Docker Support**: Containerized deployment ready
- **CI/CD**: Automated testing and Docker image publishing via GitHub Actions

## Quick Start

### Prerequisites

- Go 1.22 or later
- Docker (optional, for containerized deployment)

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Kushal-39/cicd_test.git
   cd cicd_test
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Get a fortune:
   ```bash
   curl http://localhost:8080/fortune
   ```

### Using Docker

1. Build the image:
   ```bash
   docker build -t cowsay-fortune .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 cowsay-fortune
   ```

## API Usage

### GET /fortune

Returns a random fortune formatted with cowsay ASCII art.

**Example Response:**
```
 _________________________________________
< A chicken is just an egg's way of making >
< more eggs.                               >
 -----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Project Structure

```
.
├── main.go              # HTTP server and main application
├── cowsay/              # ASCII art formatting package
│   ├── cowsay.go       # Core cowsay functionality
│   └── cowsay_test.go  # Unit tests
├── fortune/             # Fortune handling package
│   ├── fortune.go      # Fortune loading and selection
│   └── fortune_test.go # Unit tests
├── fortunes.txt         # Fortune database
├── dockerfile           # Docker container definition
├── go.mod              # Go module definition
└── .ci-cd/             # CI/CD configuration
    └── workflows/
        └── go.yml      # GitHub Actions workflow
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o cowsay-fortune ./main.go
```

### Adding New Fortunes

Add fortunes to `fortunes.txt`, separated by `%` characters:

```
Your new fortune here.
%
Another fortune.
%
```

## Deployment

The project includes automated CI/CD that:

- Runs tests on every push and pull request
- Builds and pushes Docker images to Docker Hub
- Publishes images as `boombop/cowsay-fortune:latest`

### Manual Docker Deployment

```bash
# Build and push
docker build -t boombop/cowsay-fortune:latest .
docker push boombop/cowsay-fortune:latest
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is open source. Please check the license file for details.

## Acknowledgments

- Inspired by the classic Unix `fortune` and `cowsay` commands
- Fortune collection from various public sources
