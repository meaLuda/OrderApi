# Food Ordering API

A robust REST API server implemented in Go that handles food orders, product management, and coupon validation. This implementation follows the OpenAPI 3.1 specification and includes comprehensive error handling, authentication, and validation.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Architecture](#architecture)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contributing](#contributing)

## Features

- **Product Management**
  - List all available products
  - Retrieve detailed product information
  - Category-based product organization

- **Order Processing**
  - Create new orders with multiple items
  - Validate product availability
  - Calculate order totals
  - Apply promotional discounts

- **Coupon System**
  - Sophisticated coupon validation
  - Support for multiple coupon databases
  - Length and presence validation

- **Security**
  - API key authentication
  - Input validation
  - Error handling middleware
  - Safe error responses

## Prerequisites

- Go 1.19 or higher
- Git
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
```bash
git clone git@github.com:meaLuda/OrderApi.git
cd OrderApi
```

2. Install dependencies:
```bash
go mod download
```

3. Set up configuration:
```bash
cp .env.example .env
```

4. Run the server:
```bash
go run main.go
```

## Configuration

The application can be configured using environment variables or a .env file:

```env
PORT=8080                           # Server port (default: 8080)
API_KEY=apitest                     # API key for authentication
COUPON_FILES=file1.txt,file2.txt    # Comma-separated list of coupon files
```

## API Documentation

### Endpoints

#### Products

```
GET /api/product
```
List all available products.

Response:
```json
[
  {
    "id": "1",
    "name": "Chicken Waffle",
    "price": 12.99,
    "category": "Waffle"
  }
]
```

```
GET /api/product/{productId}
```
Get details of a specific product.

Response:
```json
{
  "id": "1",
  "name": "Chicken Waffle",
  "price": 12.99,
  "category": "Waffle"
}
```

#### Orders

```
POST /api/order
```
Place a new order.

Request:
```json
{
  "couponCode": "HAPPYHRS",
  "items": [
    {
      "productId": "1",
      "quantity": 2
    }
  ]
}
```

Response:
```json
{
  "id": "0000-0000-0000-0000",
  "items": [
    {
      "productId": "1",
      "quantity": 2
    }
  ],
  "products": [
    {
      "id": "1",
      "name": "Chicken Waffle",
      "price": 12.99,
      "category": "Waffle"
    }
  ]
}
```

### Authentication

All endpoints require an API key to be provided in the header:
```
api_key: apitest
```

## Architecture

The application follows a clean architecture pattern with the following components:

### Project Structure
```
├── main.go           # Application entry point
├── handlers/         # HTTP request handlers
├── services/         # Business logic
├── models/           # Data structures
├── middleware/       # HTTP middleware
└── tests/           # Test files
```

### Components

- **Handlers**: Handle HTTP requests and responses
- **Services**: Implement business logic and data processing
- **Models**: Define data structures and validation
- **Middleware**: Implement cross-cutting concerns

## Testing

Run the test suite:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Deployment

### Docker

Build the image:
```bash
docker build -t food-ordering-api .
```

Run the container:
```bash
docker run -p 8080:8080 food-ordering-api
```

### Traditional Deployment

1. Build the binary:
```bash
go build -o food-ordering-api
```

2. Run the server:
```bash
./food-ordering-api
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

