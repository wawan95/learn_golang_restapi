# Golang REST API - User Management

A RESTful API project built with **Golang**, **Gin Web Framework**, **GORM**, and **MySQL**. This project demonstrates CRUD (Create, Read, Update, Delete) operations with a structured architecture and API documentation using **OpenAPI 3.0.3**.

## 🚀 Project Overview

This project is a User Management REST API developed to practice backend development using Go (Golang).

It provides endpoints to manage user data, including creating new users, retrieving user information, updating existing users, and deleting users.

The application follows a layered architecture to separate responsibilities and make the code easier to maintain and extend.

## 🛠️ Tech Stack

| Technology | Description |
|---|---|
| Go (Golang) | Backend programming language |
| Gin | HTTP web framework |
| GORM | ORM for database operations |
| MySQL | Relational database |
| OpenAPI 3.0.3 | API specification |
| Swagger UI | API documentation interface |
| GoDotEnv | Environment variable management |

## ✨ Features

- RESTful API development using Gin.
- CRUD operations for user management.
- MySQL database integration.
- GORM for database operations.
- Automatic database migration.
- Request validation.
- Layered architecture.
- Environment variable configuration.
- OpenAPI 3.0.3 API documentation.
- Health check endpoint.
- API testing through Swagger UI and Postman.

## 📁 Project Structure

```text
learn_golang_restapi/
│
├── config/
│   └── database.go
│
├── controllers/
│   └── user_controller.go
│
├── docs/
│   ├── index.html
│   └── openapi.yaml
│
├── models/
│   └── user.go
│
├── repositories/
│   └── user_repository.go
│
├── routes/
│   └── routes.go
│
├── services/
│   └── user_service.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## ⚙️ Getting Started

### 1. Prerequisites

Make sure you have installed:

- Go 1.27.1 or compatible version specified in `go.mod`.
- MySQL Server.
- Git.
- Postman (optional).

### 2. Clone the Repository

```bash
git clone https://github.com/wawan95/learn_golang_restapi.git
```

Navigate to the project directory:

```bash
cd learn_golang_restapi
```

### 3. Install Dependencies

Download the required Go modules:

```bash
go mod tidy
```

### 4. Configure Environment Variables

Create a `.env` file in the root project directory.

```env
# Application Configuration
APP_PORT=8080

# MySQL Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=golang_user_db
```

Replace `your_password` with your MySQL password.

### 5. Create MySQL Database

Open your MySQL client and execute:

```sql
CREATE DATABASE golang_user_db;
```

The application uses GORM AutoMigrate to create or update the users table automatically when the application starts.

### 6. Run the Application

Start the API server:

```bash
go run main.go
```

If everything is configured correctly, the application will connect to MySQL and start the Gin server.

The default application port is:

```text
http://localhost:8080
```

## 📖 API Documentation

This project provides API documentation using OpenAPI 3.0.3.

After starting the application, access the following URLs:

| Documentation | URL |
|---|---|
| Swagger UI | http://localhost:8080/swagger/index.html |
| OpenAPI YAML | http://localhost:8080/openapi.yaml |
| Health Check | http://localhost:8080/health |

## 🔗 API Endpoints

Base URL:

```text
http://localhost:8080/api/v1
```

### User Management

| Method | Endpoint | Description |
|---|---|---|
| GET | `/users` | Get all users |
| GET | `/users/:id` | Get user by ID |
| POST | `/users` | Create a new user |
| PUT | `/users/:id` | Update an existing user |
| DELETE | `/users/:id` | Delete a user |

### 1. Get All Users

```http
GET /api/v1/users
```

Retrieves all registered users.

### 2. Get User By ID

```http
GET /api/v1/users/1
```

Retrieves user information based on the specified ID.

### 3. Create User

```http
POST /api/v1/users
Content-Type: application/json
```

Request Body:

```json
{
    "name": "Sugeng Kurniawan",
    "email": "sugeng@example.com",
    "age": 30
}
```

### 4. Update User

```http
PUT /api/v1/users/1
Content-Type: application/json
```

Request Body:

```json
{
    "name": "Sugeng Kurniawan",
    "email": "sugeng.updated@example.com",
    "age": 31
}
```

### 5. Delete User

```http
DELETE /api/v1/users/1
```

Deletes a user based on the specified ID.

## 🧪 Testing

You can test the API using:

- Swagger UI.
- Postman.
- cURL.

Example using cURL:

```bash
curl http://localhost:8080/api/v1/users
```

Create a new user:

```bash
curl -X POST http://localhost:8080/api/v1/users \
-H "Content-Type: application/json" \
-d '{"name":"Sugeng Kurniawan","email":"sugeng@example.com","age":30}'
```

## 🏗️ Application Architecture

This project uses a layered architecture:

| Layer | Responsibility |
|---|---|
| Routes | Defines API endpoints |
| Controllers | Handles HTTP requests and responses |
| Services | Contains business logic |
| Repositories | Handles database operations |
| Models | Defines database entities and request structures |
| Config | Manages database connections |

This separation of concerns makes the application easier to maintain, test, and scale.

## 🎯 Learning Objectives

Through this project, I aim to improve my backend development skills, particularly in:

- Understanding Go programming fundamentals.
- Building RESTful APIs with Gin.
- Implementing CRUD operations.
- Integrating MySQL with GORM.
- Applying layered architecture.
- Managing application configuration.
- Validating API requests.
- Creating API documentation using OpenAPI.
- Testing REST API endpoints.

## 🔮 Future Improvements

Potential improvements for this project:

- JWT Authentication and Authorization.
- Password Hashing with bcrypt.
- Unit Testing and Integration Testing.
- Pagination and Search Functionality.
- Centralized Error Handling.
- Docker Integration.
- CI/CD Pipeline Integration.

## 👨‍💻 Author

**Sugeng Kurniawan**

 Software Developer | QA Engineer

GitHub: [@wawan95](https://github.com/wawan95)

---

⭐ If you find this project useful, feel free to explore the repository and give it a star!
