# Vendor Management System

A backend REST API for managing vendor information, built using **Golang, Gin Framework, GORM, and MySQL**.

The application provides CRUD operations for vendors and follows a modular backend structure separating configuration, controllers, models, repositories, and routes.

---

## 🚀 Features

- Create a new vendor
- Retrieve all vendors
- Retrieve a vendor by ID
- Update vendor information
- Delete a vendor
- MySQL database integration
- GORM ORM for database operations
- RESTful API architecture
- JSON request and response handling
- Modular project structure
- Environment-based database configuration

---

## 🛠️ Tech Stack

| Technology | Purpose |
|------------|---------|
| **Golang** | Backend programming language |
| **Gin** | HTTP web framework |
| **GORM** | ORM / database operations |
| **MySQL** | Relational database |
| **REST API** | Client-server communication |
| **Git & GitHub** | Version control |

---

## 📁 Project Structure

```text
Vendor-Management-System/
│
├── config/
│   └── dB.go
│
├── controllers/
│   └── vendor.go
│
├── models/
│   └── vendor.go
│
├── repository/
│   └── vendor.go
│
├── routes/
│   └── vendor.go
│
├── .gitignore
├── .env                 # Not committed to GitHub
├── go.mod
├── go.sum
├── main.go
└── README.md
```

### Architecture

```text
                    Client
                      │
                      ▼
                REST API / Gin
                      │
                      ▼
                  Routes
                      │
                      ▼
                Controllers
                      │
                      ▼
                 Repository
                      │
                      ▼
                    GORM
                      │
                      ▼
                   MySQL
```

---

## 🔗 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/vendors` | Create a vendor |
| `GET` | `/vendors` | Get all vendors |
| `GET` | `/vendors/:id` | Get vendor by ID |
| `PUT` | `/vendors/:id` | Update vendor |
| `DELETE` | `/vendors/:id` | Delete vendor |

> The exact route prefix depends on the route configuration in `routes/vendor.go`.

---

## 📌 Example API Request

### Create Vendor

```http
POST /vendors
Content-Type: application/json
```

Example request:

```json
{
  "name": "ABC Technologies",
  "email": "abc@example.com",
  "contact": "9876543210"
}
```

Example response:

```json
{
  "message": "Vendor created successfully",
  "vendor": {
    "id": 1,
    "name": "ABC Technologies",
    "email": "abc@example.com",
    "contact": "9876543210"
  }
}
```

---

## 🗄️ Database

The application uses **MySQL** as the relational database and **GORM** for database interaction.

Database configuration is managed using environment variables.

Example:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=vendor_management
```

⚠️ **Never commit the `.env` file to GitHub.**

---

## ⚙️ Prerequisites

Before running the application, install:

- Go
- MySQL
- Git

Verify Go installation:

```bash
go version
```

Verify Git:

```bash
git --version
```

---

## ▶️ How to Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/abhisheekkant14/Vendor-Management-System.git
```

### 2. Navigate to the project

```bash
cd Vendor-Management-System
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Configure environment variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=vendor_management
```

### 5. Start MySQL

Make sure your MySQL server is running.

### 6. Run the application

```bash
go run .
```

The API will start on the configured application port.

---

## 🧪 API Testing

The API can be tested using tools such as:

- Postman
- cURL
- Browser for GET requests

Example:

```bash
curl http://localhost:9091/vendors
```

---

## 🔐 Security

Sensitive configuration such as database credentials is stored in environment variables and excluded from version control using `.gitignore`.

```text
.env
```

---

## 🎯 Learning Objectives

This project demonstrates practical backend development concepts including:

- Golang backend development
- REST API design
- HTTP methods and status codes
- Gin framework
- CRUD operations
- MySQL integration
- GORM ORM
- Layered/modular application structure
- Repository pattern
- Environment configuration
- Git and GitHub workflow

---

## 📈 Future Improvements

Planned improvements include:

- JWT authentication and authorization
- Request validation
- Centralized error handling
- Pagination and filtering
- API logging
- Unit testing
- Docker containerization
- CI/CD pipeline
- Cloud deployment on AWS
- API documentation using Swagger/OpenAPI

---

## 👨‍💻 Author

**Abhishek Kant Mishra**

Golang Backend Developer

- GitHub: https://github.com/abhisheekkant14
- LinkedIn: https://www.linkedin.com/in/abhishek-kant-27aa62315

---

## ⭐ Project Highlights

**Golang • Gin • REST API • GORM • MySQL • CRUD • Git • GitHub**