# 🚀 go-rest-api

A production-ready RESTful API built with **Go**, **Gin**, **JWT authentication**, and **MySQL**. Follows clean architecture principles with a clear separation of concerns.

## 🛠️ Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=flat&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=flat&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=flat&logo=jsonwebtokens&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

## ✨ Features

- ✅ User registration and login
- - ✅ JWT-based authentication & authorization
  - - ✅ CRUD operations for resources
    - - ✅ MySQL database with GORM ORM
      - - ✅ Middleware for route protection
        - - ✅ Environment-based configuration
          - - ✅ Clean Architecture (handlers, services, repositories)
            - - ✅ Docker & Docker Compose support
             
              - ## 📁 Project Structure
             
              - ```
                go-rest-api/
                ├── cmd/
                │   └── main.go              # Entry point
                ├── internal/
                │   ├── handlers/            # HTTP handlers
                │   ├── middleware/          # JWT middleware
                │   ├── models/              # Data models
                │   ├── repository/          # Database layer
                │   └── services/            # Business logic
                ├── config/
                │   └── config.go            # App configuration
                ├── docker-compose.yml
                ├── .env.example
                ├── go.mod
                └── README.md
                ```

                ## 🔌 API Endpoints

                ### Auth
                | Method | Endpoint | Description |
                |--------|----------|-------------|
                | POST | `/api/v1/auth/register` | Register a new user |
                | POST | `/api/v1/auth/login` | Login and get JWT token |

                ### Users (Protected)
                | Method | Endpoint | Description |
                |--------|----------|-------------|
                | GET | `/api/v1/users` | Get all users |
                | GET | `/api/v1/users/:id` | Get user by ID |
                | PUT | `/api/v1/users/:id` | Update user |
                | DELETE | `/api/v1/users/:id` | Delete user |

                ## 🚀 Getting Started

                ### Prerequisites
                - Go 1.21+
                - - MySQL 8.0+
                  - - Docker (optional)
                   
                    - ### Run with Docker
                   
                    - ```bash
                      git clone https://github.com/ninjadiego/go-rest-api.git
                      cd go-rest-api
                      cp .env.example .env
                      docker-compose up --build
                      ```

                      ### Run locally

                      ```bash
                      git clone https://github.com/ninjadiego/go-rest-api.git
                      cd go-rest-api
                      cp .env.example .env
                      # Edit .env with your database credentials
                      go mod tidy
                      go run cmd/main.go
                      ```

                      ## 🔧 Environment Variables

                      ```env
                      DB_HOST=localhost
                      DB_PORT=3306
                      DB_USER=root
                      DB_PASSWORD=yourpassword
                      DB_NAME=go_rest_api
                      JWT_SECRET=your_jwt_secret_key
                      PORT=8080
                      ```

                      ## 📄 License

                      This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

                      ---

                      Made with ❤️ by [ninjadiego](https://github.com/ninjadiego)
