# go-ingnerd-backend

go-ingnerd-backend/
│
├── domain/                          # Core domain logic
│   ├── todo/
│   │   ├── todo.go                  # Domain model for "Todo"
│   │   └── todo_repository.go       # Repository interface for "Todo"
│   └── user/
│       ├── user.go                  # Domain model for "User"
│       └── user_repository.go       # Repository interface for "User"
│
├── application/                     # Application services layer
│   ├── dto/                         # Data Transfer Objects
│   │   ├── todo_dto.go              # DTO for "Todo"
│   │   └── user_dto.go              # DTO for "User"
│   ├── mapper/                      # Mappers for converting domain models to DTOs
│   │   ├── todo_mapper.go
│   │   └── user_mapper.go
│   └── services/
│       ├── todo_service.go          # Application service for "Todo"
│       └── user_service.go          # Application service for "User"
│
├── infrastructure/                  # Infrastructure specific implementations
│   ├── repository/
│   │   ├── todo_repository.go       # Implementation of Todo repository
│   │   └── user_repository.go       # Implementation of User repository
│   ├── database/
│   │   ├── database.go              # Database setup and connection utilities
│   │   └── schema.sql               # SQL schema definitions
│   └── middleware/
│       ├── auth_middleware.go       # Authentication middleware
│       └── logging_middleware.go    # Logging middleware
│
├── interfaces/                      # Interfaces to the outside world
│   ├── controllers/
│   │   ├── todo_controller.go       # Controller for "Todo" operations
│   │   └── user_controller.go       # Controller for "User" operations
│   └── routers/
│       └── router.go                # Setup of HTTP routing
│
├── pkg/                             # Shared packages across the application
│   ├── config/
│   │   └── config.go                # Configuration management
│   └── utils/
│       ├── logger.go                # Logger utility
│       └── validator.go             # Common validation utilities
│
├── .env                             # Environment variables
├── .gitignore                       # Specifies intentionally untracked files to ignore
├── go.mod                           # Go module definitions
├── go.sum                           # Go module checksums
├── main.go                          # Entry point of the application
└── README.md                        # Project overview and setup instructions
