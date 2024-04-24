package main

import (
	"github.com/filmons/go-ingnerd-backend/src/config"
    "github.com/filmons/go-ingnerd-backend/src/routes"
	"gorm.io/gorm"
)


var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	
	//run all routes
	routes.Routes()
}