**DROX-X**

## How to install

In the project directory, you can run:

1. Copy file `.env.example` to `.env` = `cp .env.example .env`
2. Start by Make file
```
   make up
   make down
   make restart
```
3. Migrate database: `make migrate`

## Project Structure

```
dropx/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── models/
│   │   │   ├── user.go
│   │   │   └── product.go
│   │   └── errors.go
│   │
│   ├── infrastructure/
│   │   └── database/
│   │      └── migration/
│   │
│   ├── repositories/
│   │   ├── init.go
│   │   ├── interfaces/
│   │   └── implementations/
│   │
│   ├── service/
│   │   ├── init.go
│   │   ├── user_service.go
│   │   └── auth_service.go
│   │
│   └── api/
│       ├── handler/
│       │   ├── init.go
│       │   ├── user_handler.go
│       │   └── product_handler.go
│       ├── middleware/
│       │   ├── auth.go
│       │   └── logging.go
│       ├── dto/
│       │   ├── request/
│       │   │   └── user_request.go
│       │   └── response/
│       │       └── user_response.go
│       └── router.go
│
├── pkg/
│   ├── logger/
│   ├── utils/
│   └── config/
│
├── docker/
│   └── Dockerfile
│
├── script/
│   └── migrate.sh
│
├── docker-compose.yml
├── .gitignore
├── .env
├── Makefile                           
├── go.mod
└── go.sum
```