package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/filmons/go-ingnerd-backend/src/config"
	"github.com/filmons/go-ingnerd-backend/src/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.Name == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
        return
    }

	if data.Description == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Description cannot be empty"})
        return
    }

	todo := models.Todo{} 
	// todo := models.TODOS{} 
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo
	// var todos []models.TODOS 
	println(todos)

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// var response []todoResponse
	// for _, todo := range todos {
	// 	response = append(response, todoResponse{
	// 		ID:          todo.ID,
	// 		name:        todo.Name,
	// 		description: todo.Description,
	// 	})
	// }

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})

}

func UpdateTodo(context *gin.Context) {
	var data todoRequest
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
     

	if data.Name == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
        return
    }

	if data.Description == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Description cannot be empty"})
        return
    }

	todo := models.Todo{} 
	// todo := models.TODOS{} 

	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}  
	// todo := models.TODOS{}  
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	
    result := db.First(&todo, idTodo)
    if result.Error != nil {
        // Todo item not found
        context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }


	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}