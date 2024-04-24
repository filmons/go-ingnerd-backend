// package main

// import (
// 	"log"
// 	"net/http"
// 	"github.com/filmons/go-ingnerd-backend/infrastructure/repository"
// 	"github.com/filmons/go-ingnerd-backend/application/services"
// 	"github.com/filmons/go-ingnerd-backend/interfaces/controllers"
// 	"github.com/filmons/go-ingnerd-backend/interfaces/routers"
// 	"github.com/filmons/go-ingnerd-backend/pkg/config"
// 	// "database/sql"
// 	_ "github.com/denisenkom/go-mssqldb" 
// 	"gorm.io/gorm"
// )


// var (
// 	db *gorm.DB = config.ConnectDB()
// )

// func main() {

// 	defer config.DisconnectDB(db)

// 	// Repository and service setup
// 	todoRepo := repository.NewTodoRepository(db)
// 	todoService := services.NewTodoService(todoRepo)

// 	// Controller setup
// 	todoController := controllers.NewTodoController(todoService)

// 	// Router setup
// 	router := routers.SetupRouter(todoController)

// 	log.Println("Starting server on :8080")
// 	http.ListenAndServe(":8080", router)
// }

package main

import (
    "log"
    "net/http"

    "github.com/filmons/go-ingnerd-backend/infrastructure/repository"
    "github.com/filmons/go-ingnerd-backend/application/services"
    "github.com/filmons/go-ingnerd-backend/interfaces/controllers"
    "github.com/filmons/go-ingnerd-backend/interfaces/routers"
    "github.com/filmons/go-ingnerd-backend/pkg/config"
    // "database/sql"
    _ "github.com/denisenkom/go-mssqldb"
    "gorm.io/gorm"
    "github.com/filmons/go-ingnerd-backend/pkg/utils" // Import your logger package
)

var (
    db *gorm.DB = config.ConnectDB()
)

func main() {
    defer config.DisconnectDB(db)

    // Repository and service setup
    todoRepo := repository.NewTodoRepository(db)
    todoService := services.NewTodoService(todoRepo)

    // Controller setup
    todoController := controllers.NewTodoController(todoService)

    // Router setup
    router := routers.SetupRouter(todoController)

    // Start the server
    log.Println("Starting server on :8080")

    // Log an INFO message using the logger utility
    utils.Info("Server started")

    // Start the HTTP server
    if err := http.ListenAndServe(":8080", router); err != nil {
        // Log an ERROR message if server startup fails
        log.Fatalf("Failed to start server: %v", err)
    }
}
