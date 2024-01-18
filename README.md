### Table of Contents

- [Running the Go Project](#running-the-go-project)
- [Accessing API Documentation](#accessing-api-documentation)
- [Performing Static Analysis with golangci-lint](#performing-static-analysis-with-golangci-lint)

## Running the Go Project

To run the Go project, follow these steps:

1. Ensure that you have Go installed on your system.
2. Open a terminal and navigate to the project directory.
3. Run the following command:

    ```bash
    go run cmd/main.go
    ```

This command will execute the `main.go` file located in the `cmd` directory, starting the Go project.

## Accessing API Documentation

Once the Go project is running, you can access the API documentation using the following URL:

[http://localhost:8080/docs/](http://localhost:8080/docs/)

This URL leads to the Swagger documentation, providing details on the available API endpoints, request parameters, and response structures.

## Performing Static Analysis with golangci-lint

To perform static analysis using `golangci-lint`, follow these steps:

1. Ensure that you have `golangci-lint` installed on your system. If not, you can install it with the following command:

    ```bash
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    ```

2. Open a terminal and navigate to the project directory.
3. Run the following command:

    ```bash
    golangci-lint run
    ```
