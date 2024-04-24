package database

// import (
//     "database/sql"
//     // _ "github.com/driver/for/your/database" // Import your database driver
// 	"gorm.io/driver/mysql"
//     "log"
// )

// var db *sql.DB

// // InitDB initializes the database connection
// func InitDB() (*sql.DB, error) {
//     // Connect to the database
//     connectionString := "your-database-connection-string"
//     db, err := sql.Open("your-database-driver", connectionString)
//     if err != nil {
//         log.Fatal("Error connecting to the database:", err)
//         return nil, err
//     }

//     // Check the database connection
//     err = db.Ping()
//     if err != nil {
//         log.Fatal("Error pinging database:", err)
//         return nil, err
//     }

//     return db, nil
// }

// // CloseDB closes the database connection
// func CloseDB() {
//     if db != nil {
//         db.Close()
//     }
// }
