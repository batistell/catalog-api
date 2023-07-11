# Catalog API

Catalog API is a backend API service that powers the catalog management system. It allows you to perform CRUD (Create,
Read, Update, Delete) operations on products, manage categories, and handle discounts and attributes.

## Features

- Create, read, update, and delete products
- Manage categories for organizing products
- Apply discounts to products
- Define and assign attributes to products
- API documentation and interactive UI with Swagger/OpenAPI

## Getting Started

These instructions will guide you on setting up the Catalog API locally for development and testing purposes.

### Prerequisites

- Go (v1.16 or higher) installed
- MongoDB database (local or remote) and connection details

### Installation

1. Install the dependencies:

```go mod download```

2. Configure the application:

- Copy the `config.example.yaml` file and rename it to `config.yaml`.
- Update the configuration values in `config.yaml` with your MongoDB connection details and other settings.

3. Run the application:

```go run main.go```

4. The API will start running on `http://localhost:8000`. You can access the API endpoints using a tool like cURL or an
   API client.

### API Documentation

The API is documented using Swagger/OpenAPI specification. You can access the interactive API documentation by
visiting `http://localhost:8000/docs` in your browser.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a
pull request. Make sure to follow the project's code style and guidelines.

## License

This project is licensed under the [MIT License](LICENSE).
