# REST-API

## Overview

This application is built using Go and MySQL and utilizes the Gin web framework. It allows you to perform CRUD operations on books.

## Requirements

- Go (Version 1.21 or higher)
- MySQL

## Setup

### Environment Variables

This application uses environment variables to handle configurations. To set these up:

1. Create a /config/.env.local file.
2. Add your MySQL connection string from `.env.template` to `.env.

## Running the application

After setting up the environment variables:

1. Open a terminal and navigate to the root directory of the application.
2. Run the following commands:

```bash
go mod download
go run main.go
```

The application should now be running, and you should be connected to your MySQL database.

### **Using Swagger to Test API**

To test the API using Swagger:
`swag init --parseDependency --parseInternal`
**Navigate to Swagger UI**: Open your web browser and go to **`http://localhost:8080/swagger/index.html`**.
